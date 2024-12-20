package dto

import "go_server/utils"

type LoginReq struct {
	Username  string `json:"username" binding:"required"`  // 用户名，必填
	Password  string `json:"password" binding:"required"`  // 密码，必填
	AutoLogin bool   `json:"autoLogin"`                    // 是否自动登录，可选
	GrantType string `json:"grantType" binding:"required"` // 授权类型，必填
	Type      string `json:"type" binding:"required"`      // 登录类型，必填
}

type LoginRes struct {
	Code        int    `json:"code"`
	AccessToken string `json:"access_token"` // 访问令牌
}

// GetMessages 自定义错误信息
func (loginReq LoginReq) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"Username.required":  "用户名不能为空",
		"Password.required":  "用户密码不能为空",
		"GrantType.required": "授权类型不能为空",
		"Type.required":      "登录类型不能为空",
	}
}
