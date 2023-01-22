package appctx

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type MetaData struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalPage  int64 `json:"total_page"`
	TotalCount int64 `json:"total_count"`
}

func NewResponse() *Response {
	return &Response{
		Code: http.StatusOK,
	}
}

func (r *Response) WithErrors(err string) *Response {
	r.Code = http.StatusInternalServerError
	r.Errors = append(r.Errors, err)
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}
