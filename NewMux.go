package main

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
		// 静的解析のえらーを回避するため明示的に戻り値を捨てている
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	return mux
}
