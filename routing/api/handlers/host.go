package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"create": "host",
	})
}

func GetHost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"get": "host",
	})
}

func UpdateHost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"update": "host",
	})
}

func DeleteHost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"delete": "host",
	})
}
