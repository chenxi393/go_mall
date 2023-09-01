package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	go func() {
		http.HandleFunc("/", helloHandler8002)

		// 启动服务器并监听 8000 端口
		fmt.Println("Server listening on http://localhost:8002")
		log.Fatal(http.ListenAndServe(":8002", nil))
	}()
	select {}
}

func helloHandler8002(w http.ResponseWriter, r *http.Request) {
	// 在浏览器中返回 "Hello, World!"
	fmt.Fprintln(w, "<h1>Hello, 8002!</h1>")
}
