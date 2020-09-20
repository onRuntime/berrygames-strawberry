package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onRuntime/berrygames-strawberry/data"
	"github.com/onRuntime/berrygames-strawberry/router/routes"
	"io"
	"os"
)

type Router struct {
	*gin.Engine
}

func New() *Router {
	return &Router{gin.Default()}
}

func (r *Router) Init(addr string, data *data.Data, logging bool) error {
	// - Creates and '/api' path with credentials.
	api := r.Group("/api", gin.BasicAuth(gin.Accounts{
		os.Getenv("MDLWRE_AUTH"): os.Getenv("MDLWRE_PWD"),
	}))

	// - Creates or not a .log file for the router.
	if logging {
		if f, err := os.Create("router.log"); err != nil {
			return err
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	}

	// - Register routes to router.
	r.registerRoutes(api, data)

	// - Starts the router.
	if err := r.Run(addr); err != nil {
		return err
	}
	return nil
}

func (r *Router) registerRoutes(api *gin.RouterGroup, data *data.Data) {
	api.GET("/servers", handle(data, routes.GetServers))
	api.GET("/servers/:id", handle(data, routes.GetServer))
	api.POST("/servers", handle(data, routes.PostServer))
	api.DELETE("/servers/:id", handle(data, routes.DeleteServer))
	api.PUT("/servers/:id/players", handle(data, routes.PutServerPlayer))
	api.DELETE("/servers/:id/players/:name", handle(data, routes.DeleteServerPlayer))
}

type RequestHandler func(d *data.Data, c *gin.Context)

func handle(d *data.Data, handler RequestHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(d, c)
	}
}
