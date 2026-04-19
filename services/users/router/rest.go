package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Ping(ctx *gin.Context)
	ReadUsers(ctx *gin.Context)
	ReadUserById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type routerHandler struct {
	usecase any
}

func RegisterRoute(engine *gin.Engine, router *routerHandler) *gin.Engine {
	r := gin.Default()
	r.GET("ping", router.Ping)

	group := r.Group("users")
	group.GET("", router.ReadUsers)
	group.GET("/:id", router.ReadUserById)
	group.POST("", router.CreateUser)
	group.PUT("/:id", router.UpdateUser)
	group.DELETE("/:id", router.DeleteUser)

	return r
}

func NewRouter() *routerHandler {
	return &routerHandler{
		usecase: nil,
	}
}

func (r *routerHandler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func (r *routerHandler) ReadUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to users service",
	})
}

func (r *routerHandler) ReadUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}
