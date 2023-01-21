package appctx

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
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
	return &Response{}
}
