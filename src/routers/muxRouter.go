package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxRouter = mux.NewRouter().StrictSlash(false)
)

type muxRouterStruct struct{}

func NewMuxRouter() Router {
	return &muxRouterStruct{}
}

func (mr *muxRouterStruct) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(uri, f).Methods("GET")
}

func (mr *muxRouterStruct) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(uri, f).Methods("POST")
}

func (mr *muxRouterStruct) Use(f func(h http.Handler) http.Handler) {
	muxRouter.Use(f)
}

func (mr *muxRouterStruct) Serve(port string) {
	fmt.Printf("Server running on port %v", port)
	http.ListenAndServe(":"+port, muxRouter)
}
