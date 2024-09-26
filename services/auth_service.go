package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "Login Success",
	})
}
