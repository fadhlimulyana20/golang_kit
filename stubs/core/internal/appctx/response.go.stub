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
		Code:    http.StatusOK,
		Message: "Success retrieving data",
	}
}

func (r *Response) WithErrors(err string) *Response {
	r.Code = http.StatusInternalServerError
	r.Message = "Failed retrieving data"
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

func (r *Response) WithMeta(page int64, limit int64, totalCount int64) *Response {
	r.Meta = &MetaData{
		Page:       page,
		Limit:      limit,
		TotalPage:  (totalCount / limit) + 1,
		TotalCount: totalCount,
	}
	return r
}
