package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArtist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"create": "artist",
	})
}

func GetArtist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"get": "artist",
	})
}

func UpdateArtist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"update": "artist",
	})
}

func DeleteArtist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"delete": "artist",
	})
}
