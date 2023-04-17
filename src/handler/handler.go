package handler

import (
	"encoding/json"
	"net/http"
)

type Handler func(r *http.Request) (status int, res interface{})

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res := h(r)
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(res)
}
