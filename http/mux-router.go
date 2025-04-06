package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("GET")
}
func (*muxRouter) POST(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("POST")
}
func (*muxRouter) SERVE(port string) {
	log.Printf("Server is listening on port %s", port)
	http.ListenAndServe(port, muxDispatcher)
}
