package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
    "github.com/ryokubozono/go-docker/controller"
    "github.com/ryokubozono/go-docker/docs"
    "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	r.GET("/logout", ctrl.Logout)

	menu_ctrl := controller.MenuController{}

	menu := r.Group("menu")
	menu.Use(sessionCheck)
	{
		menu.GET("/top", menu_ctrl.Top)
	}

	return r
}

func sessionCheck(c *gin.Context) {
	session := sessions.Default(c)
	loginInfo := session.Get("Username")
	if loginInfo == nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
	} else {
		c.Next()
	}
}