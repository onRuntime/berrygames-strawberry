package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onRuntime/berrygames-strawberry/data"
	"github.com/onRuntime/berrygames-strawberry/model"
	"net/http"
)

func GetServers(d *data.Data, c *gin.Context) {
	c.JSON(http.StatusOK, d.GetAllServers())
}

func GetServer(d *data.Data, c *gin.Context) {
	// - Returns query parameter
	id := c.Param("id")
	c.JSON(http.StatusOK, d.GetServer(id))
}

func PostServer(d *data.Data, c *gin.Context) {
	// - Returns inject request body into server type
	var server *model.Server
	if err := c.BindJSON(&server); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}

	// - Adds server to the database
	d.AddServer(server)
	c.JSON(http.StatusCreated, server)
}

func DeleteServer(d *data.Data, c *gin.Context) {
	_ = c.Param("id")
	// TODO: Remove server by id
	d.RemoveServer(nil)
	c.JSON(http.StatusOK, nil)
}
