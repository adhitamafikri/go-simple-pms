package rest

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
	ReadMe(ctx *gin.Context)

	ReadRoles(ctx *gin.Context)
	ReadRoleById(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)

	AssignRole(ctx *gin.Context)
}

type routerHandler struct {
	usecase any
}

func RegisterRoute(engine *gin.Engine, router *routerHandler) *gin.Engine {
	engine.GET("ping", router.Ping)

	gUsers := engine.Group("users")
	gUsers.GET("", router.ReadUsers)
	gUsers.GET("/:id", router.ReadUserById)
	gUsers.POST("", router.CreateUser)
	gUsers.PUT("/:id", router.UpdateUser)
	gUsers.DELETE("/:id", router.DeleteUser)
	gUsers.GET("/me", router.ReadMe)

	gRoles := engine.Group("roles")
	gRoles.GET("", router.ReadRoles)
	gRoles.GET("/:id", router.ReadRoleById)
	gRoles.POST("", router.CreateRole)
	gRoles.PUT("/:id", router.UpdateRole)
	gRoles.DELETE("/:id", router.DeleteRole)

	gUsers.PUT("/:id/roles", router.AssignRole)

	return engine
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

func (r *routerHandler) ReadMe(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) ReadRoles(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) ReadRoleById(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) CreateRole(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) UpdateRole(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) DeleteRole(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}

func (r *routerHandler) AssignRole(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not Implemented Yet",
	})
}
