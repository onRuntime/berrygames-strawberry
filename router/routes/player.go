package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onRuntime/berrygames-strawberry/data"
	"log"
	"net/http"
)

func GetPlayers(d *data.Data, c *gin.Context) {
	c.JSON(http.StatusOK, d.GetPlayers())
}

func GetPlayer(d *data.Data, c *gin.Context) {
	playerName := c.Param("name")
	action := c.Param("action")
	player, err := d.GetPlayer(playerName)
	log.Print(player)
	log.Print(err)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	}

	if len(action) > 0 {
		switch action {
		case "connect":
			serverType := c.Query("to")
			if ok, queue := d.IsPlayerInQueue(playerName); ok {
				d.RemovePlayerFromQueue(playerName)

				// - Detects if player queue equals to requested queue
				if serverType == queue.ServerType {
					c.JSON(http.StatusOK, nil)
				}
			}

			d.AddPlayerToQueue(serverType, playerName)
			c.JSON(http.StatusOK, nil)
		case "disconnect":
			return
		}
	}

	c.JSON(http.StatusOK, player)
}

func PatchPlayer(d *data.Data, c *gin.Context) {
	playerName := c.Param("name")
	player, err := d.GetPlayer(playerName)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
	}
	if err := c.BindJSON(&player); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusCreated, player)
}
