// Copyright © 2022 zc2638 <zc2638@qq.com>.

package command

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zc2638/ylgy/pkg/core"
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
				if err := core.Send(opt.Token); err != nil {
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
