package handler

import "net/http"

func ChinaList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("china list"))
}
