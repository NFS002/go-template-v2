package api

import (
	"fmt"
	"net/http"
	"nfs002/template/v2/auth/middleware"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	token := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	claims := token.CustomClaims.(*middleware.Claims)
	fmt.Printf("Claims: %+v", claims)
	c.JSON(http.StatusOK, gin.H{
		"create": "user",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"get": "user",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"update": "user",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"delete": "user",
	})
}
