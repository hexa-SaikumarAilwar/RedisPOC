package router

import "net/http"

type Router interface {
	GET(url string, f func(resp http.ResponseWriter, req *http.Request))
	POST(url string, f func(resp http.ResponseWriter, req *http.Request))
	SERVE(port string)
}
