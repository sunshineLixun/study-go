package ch9

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", httpHandle)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func httpHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", r.Method)
	// /go
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// 回复
	w.Write([]byte("www.5lmh.com"))
}
