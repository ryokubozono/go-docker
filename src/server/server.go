package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ryokubozono/go-docker/controller"
	"github.com/ryokubozono/go-docker/docs"
	"github.com/ryokubozono/go-docker/middleware/jwt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/"

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	ctrl := controller.RootController{}
	r.GET("/", ctrl.RootGet)

	r.GET("/ping", ctrl.DbPing)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/create_user", ctrl.CreateUser)

	r.POST("/login", ctrl.Login)

	menu_ctrl := controller.MenuController{}

	menu := r.Group("menu")
	menu.Use(jwt.JWT())
	{
		menu.GET("/top", menu_ctrl.Top)
	}

	return r
}
