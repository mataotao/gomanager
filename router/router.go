package router

import (
	"net/http"
	"apiserver/handler/admin/manager/permission"
	"apiserver/handler/admin/manager/role"
	"apiserver/handler/admin/user"
	"apiserver/handler/sd"
	"apiserver/router/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// The user handlers, requiring authentication
	u := g.Group("/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}
	/////////////////////////////////////////////////////后台 start///////////////////////////////////////////////////////////////////////////

	admin := g.Group("/admin/")
	admin.Use(middleware.AuthMiddleware())
	{
		//manager module
		////////////////////权限
		//新增权限
		admin.POST("manager/permission", permission.Create)
		//删除权限
		admin.DELETE("manager/permission/:id", permission.Delete)
		//修改权限
		admin.PUT("manager/permission/:id", permission.Update)
		//获取一条
		admin.GET("manager/permission/:id", permission.Get)
		//获取权限列表
		admin.GET("manager/permission", permission.List)
		////////////////////角色
		//新增角色
		admin.POST("manager/role", role.Create)
		//删除角色
		admin.DELETE("manager/role/:id", role.Delete)
	}

	/////////////////////////////////////////////////////后台 start///////////////////////////////////////////////////////////////////////////
	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
