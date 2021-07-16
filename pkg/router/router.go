package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
	"nubeio-rubix-lib-rest-go/controller"
	"nubeio-rubix-lib-rest-go/pkg/middleware"
)


func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.Security())
	r.Use(middleware.MyLimit())

	api := controller.Controller{DB: db}

	// Non-protected routes
	networks := r.Group("/api/networks")
	{
		networks.GET("/", api.GetNetworks)
		networks.GET("/:uuid", api.GetNetwork)
		networks.POST("/", api.AddNetwork)
		networks.PATCH("/:uuid", api.UpdateNetwork)
		networks.DELETE("/:uuid", api.DeleteNetwork)
	}

	devices := r.Group("/api/devices")
	{
		devices.GET("/", api.GetDevices)
		devices.GET("/:uuid", api.GetDevice)
		devices.POST("/", api.AddDevice)
		devices.PATCH("/:uuid", api.UpdateDevice)
		devices.DELETE("/:uuid", api.DeleteDevice)
	}

	points := r.Group("/api/points")
	{
		points.GET("/", api.GetPoints)
		points.GET("/:uuid", api.GetPoint)
		points.POST("/", api.AddPoint)
		points.PUT("/:uuid", api.UpdatePoint)
		points.DELETE("/:uuid", api.DeletePoint)
	}

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signup", api.Signup)
		authRouter.POST("/signin", api.Signin)
		authRouter.POST("/refresh", api.RefreshToken)
		authRouter.POST("/check", api.CheckToken)
	}


	// JWT-protected routes
	//postsjwt := r.Group("/postsjwt", middleware.Authorize())
	//{
	//	postsjwt.GET("/", api.GetPoints)
	//	postsjwt.GET("/:id", api.GetPoints)
	//	postsjwt.POST("/", api.GetPoints)
	//	postsjwt.PUT("/:id", api.GetPoints)
	//	postsjwt.DELETE("/:id", api.GetPoints)
	//}

	// Protected routes
	// For authorized access, group protected routes using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	//authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
	//	"username1": "password1",
	//	"username2": "password2",
	//	"username3": "password3",
	//}))

	// /admin/dashboard endpoint is now protected
	//authorized.GET("/dashboard", controllers.Dashboard)
	// /swagger/index.html
	url := ginSwagger.URL("http://localhost:1920/swagger/doc.json")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
