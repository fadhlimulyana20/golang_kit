package handler

import (
	"net/http"
	"template/internal/appctx"
	"template/internal/usecase/hello"
	"time"
)

type helloHandler struct {
	handler      handler
	helloUsecase hello.HelloUsecase
}

func NewHelloHandler() *helloHandler {
	return &helloHandler{
		helloUsecase: hello.NewHelloUsecase(),
	}
}

func (h *helloHandler) Hello(w http.ResponseWriter, r *http.Request) {
	d := appctx.Data{
		Request: r,
	}
	resp := h.helloUsecase.Hello(d)
	h.handler.Response(w, resp, time.Now())
}
