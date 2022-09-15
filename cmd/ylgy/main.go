// Copyright © 2022 zc2638 <zc2638@qq.com>.

package main

import (
	"math/rand"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/zc2638/ylgy/pkg/command"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cmd := command.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
