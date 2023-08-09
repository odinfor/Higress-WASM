package main

import (
	"cc_deny/internal"
	"cc_deny/internal/types"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, _ := os.Getwd()
	confPath := filepath.Join(filepath.Dir(path), "conf/cc_deny.yaml")
	fmt.Println(confPath)

	types.InitConf(confPath)

	fmt.Println(types.NewConfDo().LimiterConf())

	limiter := internal.NewTokenLimiter()

	go func(t *internal.TokenLimiter) {
		if t.Allow() {
			fmt.Println("success")
			// 处理 access 限制
		}
	}(limiter)

	var c chan int
	<-c
}
