package router

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ginRouter struct {
	engine *gin.Engine
}

func NewGinRouter() Router {
	return &ginRouter{
		engine: gin.Default(),
	}
}

func (g *ginRouter) GET(url string, f func(http.ResponseWriter, *http.Request)) {
	g.engine.GET(url, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		req := c.Request.WithContext(
			context.WithValue(c.Request.Context(), "id", id),
		)
		f(c.Writer, req)
	})
}

func (g *ginRouter) POST(url string, f func(http.ResponseWriter, *http.Request)) {
	g.engine.POST(url, func(c *gin.Context) {
		f(c.Writer, c.Request)
	})
}

func (g *ginRouter) SERVE(port string) {
	g.engine.Run(port)
}
