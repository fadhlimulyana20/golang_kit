package middleware

import "net/http"

func Pagination(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if !q.Has("page") {
			q.Add("page", "1")
		}
		if !q.Has("limit") {
			q.Add("limit", "25")
		}

		r.URL.RawQuery = q.Encode()
		handler.ServeHTTP(w, r)
	})
}
