package controllers

import (
	"tng/shipper-service/dtos"
	"tng/shipper-service/services"
)


type UserController struct {
	BaseController
	userService services.UserService
}

func (u *UserController) Prepare() {
	_ = services.GetServiceContainer().Invoke(func(s services.UserService) {
		u.userService = s
	})
}

// Login
// @Title Login with zalo social
// @Description Login with zalo social
// @Param body body	dtos.LoginRequest	true	"Information request"
// @Success 200 {object} dtos.LoginResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /loginzalo [post]
func (u *UserController) Login(body *dtos.LoginRequest) {
	u.Respond(u.userService.LoginWithZalo(u.Ctx.Request.Context(), body, *u.Ctx.Input))
}

// Login using username and password
// @Title Login with password
// @Description Login with zalo social
// @Param body body	dtos.LoginWithPasswordRequest	true	"Information request"
// @Success 200 {object} dtos.LoginWithPasswordResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /login [post]
func (u *UserController) LoginWithPassword(body *dtos.LoginWithPasswordRequest) {
	u.Respond(u.userService.LoginWithPassword(u.Ctx.Request.Context(), body, *u.Ctx.Input))
}

// Sign up
// @Title Sign up
// @Description Register username and password
// @Param body body	dtos.SignUpRequest	true	"Information request"
// @Success 200 {object} dtos.SignUpResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /signup [post]
func (u *UserController) Signup(body *dtos.SignUpRequest) {
	u.Respond(u.userService.SignUp(u.Ctx.Request.Context(), body))
}

// Verify Phone Number
// @Title Verify Phone Number
// @Description Verify phone number
// @Param body body	dtos.VerifyPhoneNumberRequest	true	"Information request"
// @Success 200 {object} dtos.VerifyPhoneNumberResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /verify [post]
func (u *UserController) VerifyPhoneNumber(body *dtos.VerifyPhoneNumberRequest) {
	u.Respond(u.userService.VerifyPhoneNumber(u.Ctx.Request.Context(), body))
}

// Refresh Token
// @Title Refresh Token
// @Description Refresh token
// @Param body body	dtos.RefreshTokenRequest	true	"Information request"
// @Success 200 {object} dtos.RefreshTokenResponse
// @Failure 400 Bad Request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @router /refresh [post]
func (u *UserController) RefreshToken(body *dtos.RefreshTokenRequest) {
	u.Respond(u.userService.RefreshToken(u.Ctx.Request.Context(), body))
}