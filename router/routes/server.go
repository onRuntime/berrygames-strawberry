package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onRuntime/berrygames-strawberry/data"
	"net/http"
)

func GetServers(d *data.Data, c *gin.Context) {
	c.JSON(http.StatusOK, d.GetAllServers())
}

func GetServer(d *data.Data, c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, d.GetServer(id))
}

func PostServer(d *data.Data, c *gin.Context) {}

func DeleteServer(d *data.Data, c *gin.Context) {}

func PutServerPlayer(d *data.Data, c *gin.Context) {}

func DeleteServerPlayer(d *data.Data, c *gin.Context) {}
