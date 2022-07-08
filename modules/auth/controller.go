package auth

import (
	"net/http"

	"github.com/eriicafes/go-api-starter/models"
	"github.com/eriicafes/go-api-starter/response"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authService AuthService
}

func NewAuthController(authService AuthService) *authController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Profile(ctx *gin.Context) {
	res := response.New(ctx)

	user := ctx.MustGet("auth").(*models.User)

	res.SetData(user).JSON()
}

func (c *authController) SignIn(ctx *gin.Context) {
	res := response.New(ctx)

	var signInDto struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&signInDto); err != nil {
		res.SetStatusMessage(http.StatusUnprocessableEntity, "validation error").SetError(err.Error()).ErrJSON()
		return
	}

	userData := models.User{
		Name: signInDto.Name,
	}

	user := c.authService.SignIn(userData)

	res.SetData(user).JSON()
}

func (c *authController) SignOut(ctx *gin.Context) {
	res := response.New(ctx)

	user := ctx.MustGet("auth").(*models.User)

	c.authService.SignOut(user.AccountID)

	res.SetStatus(http.StatusNoContent).JSON()
}
