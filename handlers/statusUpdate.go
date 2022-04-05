package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Weather struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func Update(c *gin.Context) {
	var weather Weather
	weather.Wind = rand.Intn(100)
	weather.Water = rand.Intn(100)
	var result gin.H

	result = gin.H{
		"status": weather,
	}

	c.JSON(http.StatusOK, result)
}

func PageView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"result": "foo",
	})
}
