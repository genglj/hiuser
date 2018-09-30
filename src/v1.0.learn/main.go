package main

import (
	"net/http"
)

// 程序入口函数
func main() {
	http.Handle("/api/users", "")
}
