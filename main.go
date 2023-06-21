package main

import (
	"net/http"

	"go.kuoruan.net/adguard-upstream/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/china-list", http.HandlerFunc(handler.ChinaList))

	http.ListenAndServe(":8080", mux)
}
