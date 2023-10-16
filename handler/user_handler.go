package handler

import (
	"Valter/service"
	"Valter/utility"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) *UserHandler {
	return &UserHandler{us}
}

func (uh *UserHandler) Register(c *gin.Context) {
	username := c.PostForm("uname")
	email := c.PostForm("email")
	number := c.PostForm("number")
	address := c.PostForm("address")
	password := c.PostForm("password")

	user, err := uh.userService.AddNewUser(c, username, email, number, address, password)
	if err != nil {
		return
	}
	utility.HttpSuccessResponse(c, "New user created", user)
}

func (uh *UserHandler) Login(c *gin.Context) {
	uname := c.PostForm("uname")
	pass := c.PostForm("pass")

	tokenStr, err := uh.userService.LoginUser(c, uname, pass)
	if err != nil {
		return
	}
	utility.HttpSuccessResponse(c, "Login success", map[string]string{
		"token": tokenStr,
	})
}

func (uh *UserHandler) ForgotPass(c *gin.Context) {
	email := c.PostForm("email")
	data, err := uh.userService.GetUserbyEmail(email)
	if err != nil {
		utility.HttpDataNotFound(c, "Data not found", err)
		return
	}
	code := utility.GenerateCode()
	err = uh.userService.SetToken(data.ID, code)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to save token", err)
		return
	}
	utility.SendMails(email, code)
	utility.HttpSuccessResponse(c, "Code sent to email", nil)
}

func (uh *UserHandler) VerfiyCode(c *gin.Context) {
	email := c.PostForm("email")
	newPass := c.PostForm("pass")
	code := c.PostForm("code")

	data, err := uh.userService.GetUserbyEmail(email)
	if err != nil {
		utility.HttpDataNotFound(c, "Data not found", err)
		return
	}
	if code != data.Token {
		utility.HttpBadRequest(c, "Your code is wrong")
		return
	}
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to hashing", err)
		return
	}
	err = uh.userService.ForgotPass(c, data.ID, newPass)
	if err != nil {
		utility.HttpInternalErrorResponse(c, "Failed to update password", err)
		return
	}
	utility.HttpSuccessResponse(c, "Success to update password", nil)
}
