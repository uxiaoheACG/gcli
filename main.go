package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	// 参数定义
	var method string
	var url string
	var headers []string
	var body string
	var version bool

	// 参数绑定
	pflag.StringVarP(&method, "method", "X", "GET", "请求方法（GET/POST）")
	pflag.StringVarP(&url, "url", "u", "", "请求地址")
	pflag.StringSliceVarP(&headers, "header", "H", []string{}, "请求头（格式: Key: Value）")
	pflag.StringVarP(&body, "body", "d", "", "请求体（仅用于 POST）")
	pflag.BoolVarP(&version, "version", "v", false, "显示版本信息")

	pflag.Parse()

	// 版本信息
	if version {
		fmt.Println("HttpCLI v0.1.0 by 华强 徐")
		os.Exit(0)
	}

	// 参数校验
	if url == "" {
		fmt.Println("错误：必须提供 --url 参数")
		os.Exit(1)
	}

	// 构造请求
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBufferString(body))
	if err != nil {
		fmt.Println("构造请求失败:", err)
		os.Exit(1)
	}

	// 设置请求头
	for _, h := range headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			fmt.Println("错误的 header 格式，应为 Key: Value")
			os.Exit(1)
		}
		req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 输出结果
	fmt.Println("响应状态码:", resp.Status)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("响应体:")
	fmt.Println(string(respBody))
}
