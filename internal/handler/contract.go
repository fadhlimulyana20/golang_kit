package handler

import (
	"encoding/json"
	"net/http"
	"template/internal/appctx"
	"time"
)

// type response func(w http.ResponseWriter, resp appctx.Response, startTime time.Time)

type handler struct{}

func (h *handler) Response(w http.ResponseWriter, resp appctx.Response, startTime time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	d, _ := json.Marshal(resp)
	w.Write(d)
}
