package controllers

import (
	"github.com/gin-gonic/gin"
	"go_server/system/dto"
)

func Login(ctx *gin.Context) {
	var logReq *dto.LoginReq

	if err := ctx.ShouldBindJSON(&logReq); err != nil {

	}
}
