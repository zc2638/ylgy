// Copyright © 2022 zc2638 <zc2638@qq.com>.

package command

import (
	"errors"
	"fmt"
	"sync"

	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cobra"
	"github.com/zc2638/ylgy/pkg/core"
)

type Option struct {
	UID   string
	Token string
	Times int
}

func NewRootCommand() *cobra.Command {
	opt := new(Option)

	cmd := &cobra.Command{
		Use:          "ylgy",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(opt.Token) == 0 && len(opt.UID) == 0 {
				return errors.New("token 或 uid 必须任选其一填写")
			}

			times := 1
			if opt.Times > 0 {
				times = opt.Times
			}

			var wg sync.WaitGroup
			p, err := ants.NewPoolWithFunc(1000, func(i interface{}) {
				defer wg.Done()

				var (
					mode string
					err  error
				)
				if len(opt.UID) > 0 {
					mode = "uid"
					err = core.SendByUID(opt.UID)
				} else {
					mode = "token"
					err = core.Send(opt.Token)
				}

				if err != nil {
					fmt.Printf("[MODE=%s][%d] 失败! 请求超时错误! \n", mode, i)
				} else {
					fmt.Printf("[MODE=%s][%d] 通关！\n", mode, i)
				}
			})
			if err != nil {
				return err
			}
			defer p.Release()

			for i := 1; i <= times; i++ {
				wg.Add(1)
				_ = p.Invoke(i)
			}
			wg.Wait()
			return nil
		},
	}

	cmd.Flags().StringVar(&opt.UID, "uid", opt.UID, "设置uid")
	cmd.Flags().StringVar(&opt.Token, "token", opt.Token, "设置token")
	cmd.Flags().IntVar(&opt.Times, "times", opt.Times, "设置次数")
	return cmd
}
