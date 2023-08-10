package main

import (
	"cc_deny/internal"
	"fmt"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"os"
	"path/filepath"
)

func main() {

	path, _ := os.Getwd()
	confPath := filepath.Join(filepath.Dir(path), "conf/cc_deny.json")
	fmt.Println(confPath)

	wrapper.SetCtx(
		// 插件名称
		"cc-deny-plugin",

		// 为解析插件配置，设置自定义函数
		wrapper.ParseConfigBy(internal.ParseConfig),

		// 为处理请求头，设置自定义函数
		wrapper.ProcessRequestHeadersBy(internal.OnHttpRequestHeaders),
	)

	limiter := internal.NewTokenLimiter()

	go func(t *internal.TokenLimiter) {
		if t.Allow() {
			fmt.Println("success")
			// 处理 access 限制
		}
	}(limiter)

	var c = make(chan int, 1)
	<-c
}
