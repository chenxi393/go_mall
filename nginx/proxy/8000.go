package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置路由处理函数
	go func() {
		http.HandleFunc("/", helloHandler8000)

		// 启动服务器并监听 8000 端口
		fmt.Println("Server listening on http://localhost:8000")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()
	select {}
}

func helloHandler8000(w http.ResponseWriter, r *http.Request) {
	// 在浏览器中返回 "Hello, World!"
	fmt.Fprintln(w, "<h1>Hello, 8000!</h1>")
}

