package routers

import (
	"net/http"
)

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	Use(f func(h http.Handler) http.Handler)
	Serve(port string)
}
