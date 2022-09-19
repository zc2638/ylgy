// Copyright © 2022 zc2638 <zc2638@qq.com>.

package core

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

var client = resty.New().
	SetBaseURL("https://cat-match.easygame2021.com").
	SetHeader("Host", "cat-match.easygame2021.com").
	SetHeader("Content-Type", "application/json").
	SetHeader("Accept-Encoding", "gzip,compress,br,deflate").
	SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.28(0x18001c25) NetType/WIFI Language/zh_CN")

func send(url, rankRole, token string) error {
	// 生成通关时间(s)
	consumeTime := rand.Int63n(1000)
	resp, err := client.R().
		SetHeader("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/15/page-frame.html").
		SetQueryParam("rank_score", "1").
		SetQueryParam("rank_state", "1").
		SetQueryParam("rank_time", strconv.FormatInt(consumeTime, 10)).
		SetQueryParam("rank_role", rankRole).
		SetQueryParam("skin", "1").
		SetHeader("t", token).
		Get(url)

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
	return nil
}

func Send(token string) error {
	if err := send("/sheep/v1/game/game_over", "1", token); err != nil {
		return fmt.Errorf("完成闯关失败: %v", err)
	}
	if err := send("/sheep/v1/game/topic_game_over", "2", token); err != nil {
		return fmt.Errorf("完成话题失败: %v", err)
	}
	return nil
}
