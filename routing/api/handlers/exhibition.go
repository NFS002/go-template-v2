package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateExhibition(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"create": "exhibition",
	})
}

func GetExhibition(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"get": "exhibition",
	})
}

func UpdateExhibition(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"update": "exhibition",
	})
}

func DeleteExhibition(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"delete": "exhibition",
	})
}
