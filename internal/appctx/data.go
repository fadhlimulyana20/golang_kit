package appctx

import "net/http"

type Data struct {
	Request *http.Request
}
