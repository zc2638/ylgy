// Copyright © 2022 zc2638 <zc2638@qq.com>.

package core

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"

	"google.golang.org/protobuf/proto"

	"github.com/zc2638/ylgy/pb"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type Engine struct {
	client *resty.Client
	token  string
}

func New(token string) *Engine {
	client := resty.New().
		SetBaseURL("https://cat-match.easygame2021.com").
		SetHeader("Host", "cat-match.easygame2021.com").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip,compress,br,deflate").
		SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.28(0x18001c25) NetType/WIFI Language/zh_CN")
	return &Engine{token: token, client: client}
}

func dealError(resp *resty.Response, err error) error {
	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("[%d] 请求错误: %s", resp.StatusCode(), resp.String())
	}

	jsonResult := gjson.Parse(resp.String())
	if jsonResult.Get("err_code").Int() != 0 {
		return fmt.Errorf("请求错误: %s", resp.String())
	}
	return err
}

func (e *Engine) getMapHash() (string, error) {
	resp, err := e.client.R().
		SetHeader("t", e.token).
		SetQueryParam("matchType", "3").
		Get("/sheep/v1/game/map_info_ex")
	if err := dealError(resp, err); err != nil {
		return "", err
	}
	jsonResult := gjson.Parse(resp.String())
	mapHash := jsonResult.Get("data.map_md5.1").String()
	if mapHash == "" {
		return "", fmt.Errorf("获取map信息失败: %v", jsonResult.Get("data").String())
	}
	return mapHash, nil
}

func (e *Engine) getMapInfo(name string) ([]Item, error) {
	resp, err := resty.New().R().Get(fmt.Sprintf("https://cat-match-static.easygame2021.com/maps/%s.txt", name))
	if err := dealError(resp, err); err != nil {
		return nil, err
	}

	jsonResult := gjson.Parse(resp.String())
	levelData := jsonResult.Get("levelData").Map()
	// 统计块数
	var total int
	for _, v := range levelData {
		total += len(v.Array())
	}

	items := make(map[int]Item)
	for i := 0; i < total; i++ {
		items[i] = Item{ChessIndex: i, TimeTag: rand.Intn(1000)}
	}

	// 打乱通关顺序
	result := make([]Item, 0, total)
	for _, v := range items {
		result = append(result, v)
	}
	return result, nil
}

func (e *Engine) Send() error {
	mapHash, err := e.getMapHash()
	if err != nil {
		return err
	}
	steps, err := e.getMapInfo(mapHash)
	if err != nil {
		return fmt.Errorf("获取第二关steps失败: %v", err)
	}
	pbSteps := make([]*pb.Step, 0, len(steps))
	for _, v := range steps {
		pbSteps = append(pbSteps, &pb.Step{
			ChessIndex: uint32(v.ChessIndex),
			TimeTag:    uint32(v.TimeTag),
		})
	}

	data := &pb.Data{
		GameType:     3,
		StepInfoList: pbSteps,
	}
	signData, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	sign := base64.StdEncoding.EncodeToString(signData)

	var duration int64
	for _, v := range data.StepInfoList {
		duration += int64(v.TimeTag)
	}
	resp, err := e.client.R().
		SetBody(GameOverData{
			RankTime:      duration,
			RankRole:      1,
			RankScore:     1,
			RankState:     1,
			Skin:          1,
			MatchPlayInfo: sign,
		}).
		SetHeader("t", e.token).
		Post("/sheep/v1/game/game_over_ex")
	if err := dealError(resp, err); err != nil {
		return err
	}
	return nil
}

type GameOverData struct {
	RankTime      int64  `json:"rank_time"`
	RankRole      int    `json:"rank_role"`
	RankScore     int    `json:"rank_score"`
	RankState     int    `json:"rank_state"`
	Skin          int    `json:"skin"`
	MatchPlayInfo string `json:"MatchPlayInfo"`
}

type Data struct {
	GameType     int    `json:"gameType"`
	StepInfoList []Item `json:"stepInfoList"`
}

type Item struct {
	ChessIndex int `json:"chessIndex"`
	TimeTag    int `json:"timeTag"`
}
