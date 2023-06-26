package main

import (
	"log"
	"net/http"

	"go.kuoruan.net/adguard-upstream/handler"
	"go.kuoruan.net/adguard-upstream/service"
)

func main() {
	mux := http.NewServeMux()

	chinaListService := service.NewChinaList()

	if err := chinaListService.Update(); err != nil {
		log.Fatalln(err)
	}

	chinaListHandler := handler.NewChinaList(chinaListService)

	mux.Handle("/china-list", chinaListHandler)

	http.ListenAndServe(":8080", mux)
}
