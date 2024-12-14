package controllers

import (
	"github.com/gin-gonic/gin"
	"go_server/global"
	"go_server/models"
	"go_server/utils"
	"net/http"
)

//		Register
//	 @Description: 注册用户并返回Token
//	 @param ctx
func Register(ctx *gin.Context) {
	var user models.User
	/*
		ShouldBindJSON传入数据对应的结构体
		如果json数据与传入的结构体数据相符合(可以有多余字段，但是不能缺少字段)
		成功会返回一个nil,否则会返回一个错误对象
	*/
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPwd, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.DB.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
