package handler

import (
	"net/http"
	"strings"

	"go.kuoruan.net/adguard-upstream/service"
)

type ChinaList struct {
	Service *service.ChinaList
}

func NewChinaList(service *service.ChinaList) *ChinaList {
	return &ChinaList{
		Service: service,
	}
}

func (h *ChinaList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.Write([]byte(strings.Join(h.Service.AcceleratedDomains, "\n")))
}
