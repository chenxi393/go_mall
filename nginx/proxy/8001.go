package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	go func() {
		http.HandleFunc("/", helloHandler8001)

		// 启动服务器并监听 8000 端口
		fmt.Println("Server listening on http://localhost:8001")
		log.Fatal(http.ListenAndServe(":8001", nil))
	}()

	select {}
}

func helloHandler8001(w http.ResponseWriter, r *http.Request) {
	// 在浏览器中返回 "Hello, World!"
	fmt.Fprintln(w, "<h1>Hello, 8001!</h1>")
}
