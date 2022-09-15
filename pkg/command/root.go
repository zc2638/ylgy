// Copyright © 2022 zc2638 <zc2638@qq.com>.

package command

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

type Option struct {
	Token string
	Times int
}

func NewRootCommand() *cobra.Command {
	opt := new(Option)

	cmd := &cobra.Command{
		Use:          "ylgy",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(opt.Token) == 0 {
				return errors.New("token 必填")
			}

			times := 1
			if opt.Times > 0 {
				times = opt.Times
			}
			for i := 0; i < times; i++ {
				if err := Send(opt.Token); err != nil {
					return err
				}
				fmt.Printf("[%d] 通关！\n", i+1)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&opt.Token, "token", opt.Token, "设置token")
	cmd.Flags().IntVar(&opt.Times, "times", opt.Times, "设置次数")
	return cmd
}

var client = resty.New().
	SetBaseURL("https://cat-match.easygame2021.com").
	SetHeader("Host", "cat-match.easygame2021.com").
	SetHeader("Content-Type", "application/json").
	SetHeader("Accept-Encoding", "gzip,compress,br,deflate").
	SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 15_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.28(0x18001c25) NetType/WIFI Language/zh_CN")

func Send(token string) error {
	// 生成通关时间(s)
	consumeTime := rand.Int63n(1000)

	resp, err := client.R().
		SetHeader("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/15/page-frame.html").
		SetQueryParam("rank_score", "1").
		SetQueryParam("rank_state", "1").
		SetQueryParam("rank_time", strconv.FormatInt(consumeTime, 10)).
		SetQueryParam("rank_role", "1").
		SetQueryParam("skin", "1").
		SetHeader("t", token).
		Get("/sheep/v1/game/game_over")

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
