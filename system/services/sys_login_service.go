package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_server/common"
	"go_server/global"
	"go_server/system/dto"
	"go_server/system/models"
	"go_server/utils"
)

type GinJWTMiddleware struct {
}

func Login(c *gin.Context) {
	//获取用户输入的用户名和密码绑定解析参数
	var login dto.LoginReq
	if err := c.ShouldBindJSON(&login); err != nil {
		fmt.Println("JSON有误", err.Error())
		common.Error(c, common.ErrorCode, utils.GetErrorMsg(login, err))
		return
	}
	//查询数据库数据用户是否存在
	var user models.SysUser
	err := global.DB.Where("username = ?", login.Username).First(&user).Error
	if err != nil {
		fmt.Println("数据库报错", err.Error())
		common.Error(c, common.ErrorCode, "登录用户不存在")
		return
	}
	//校验密码是否正确
	_, err = utils.CompareHashAndPassword(user.Password, login.Password)
	if err != nil {
		common.Error(c, common.ErrorCode, "密码不正确")
		return
	}
	common.LoginSuccess(c, dto.LoginRes{
		Code:        200,
		AccessToken: "xxxxxxxx111",
	})
	fmt.Println("成功", "Hello World")
}
