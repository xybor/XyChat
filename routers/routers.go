package routers

import (
	"github.com/gin-gonic/gin"
	api1 "github.com/xybor/xychat/controllers/api/v1"
	ws1 "github.com/xybor/xychat/controllers/ws/v1"
	"github.com/xybor/xychat/helpers/context"
	"github.com/xybor/xychat/middlewares"
)

// Route combines middlewares and controllers to handle given url paths in the
// application.
func Route() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("vue/dist/*.html")

	router.StaticFile("/", "vue/dist/index.html")
	router.Static("/js", "vue/dist/js")
	router.Static("/css", "vue/dist/css")

	router.NoRoute(func(ctx *gin.Context) { ctx.HTML(200, "index.html", gin.H{}) })

	rapi := router.Group("api")
	rapi.Use(
		middlewares.VerifyUserToken(true),
		middlewares.ApplyCORSHeader(),
	)
	{
		rapi1 := rapi.Group("v1")
		{
			rapi1.POST("auth",
				middlewares.MustHaveQueryParam(context.POST, "username", "password"),
				api1.UserAuthenticateHandler,
			)
			rapi1.POST("register",
				middlewares.MustHaveQueryParam(context.POST, "username", "password"),
				api1.UserRegisterHandler,
			)
			rapi1.GET("profile", api1.UserProfileHandler)

			rapi1.GET("users/:id", api1.UserGETHandler)
			rapi1.PUT("users/:id", api1.UserPUTHandler)
			rapi1.PUT("users/:id/role",
				middlewares.MustHaveQueryParam(context.POST, "role"),
				api1.UserChangeRoleHandler,
			)
			rapi1.PUT("users/:id/password",
				middlewares.MustHaveQueryParam(context.POST, "newpassword"),
				api1.UserChangePasswordHandler,
			)
		}
	}

	rws := router.Group("ws")
	rws.Use(
		middlewares.VerifyUserToken(false),
		middlewares.UpgradeToWebSocket,
	)
	{
		rws1 := rws.Group("v1")
		{
			rws1.GET("match", ws1.WSMatchHandler)
			rws1.GET("chat", ws1.WSChatHandler)
		}
	}

	return router
}
