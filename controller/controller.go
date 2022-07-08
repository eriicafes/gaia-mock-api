package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Routes(r *gin.RouterGroup)
}

type binder struct {
	router *gin.Engine
}

func NewBinder(router *gin.Engine) *binder {
	return &binder{router}
}

func (b *binder) Bind(path string, controller Controller, middlewares ...gin.HandlerFunc) {
	if controller == nil {
		return
	}

	group := b.router.Group(path, middlewares...)
	controller.Routes(group)
}
