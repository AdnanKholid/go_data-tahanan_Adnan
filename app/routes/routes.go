package routes

import (
	"mini_project/app/middlewares"

	csel "mini_project/controllers/c_sel"
	csipir "mini_project/controllers/c_sipir"
	ctahanan "mini_project/controllers/c_tahanan"
	cuser "mini_project/controllers/c_user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RoutService(server *echo.Echo) {

	selRoute := server.Group("/sel")

	selRoute.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))
	selRoute.Use(middlewares.CheckTokenMiddlewareUser)

	selRoute.GET("/get_all", csel.GetAll)
	selRoute.GET("/get_by_id/:id", csel.GetByID)
	selRoute.POST("/create", csel.Create)
	selRoute.PUT("/update/:id", csel.Update)
	selRoute.DELETE("/delete/:id", csel.Delete)

	tahananRoute := server.Group("/tahanan")

	tahananRoute.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))
	tahananRoute.Use(middlewares.CheckTokenMiddlewareUser)

	tahananRoute.GET("/get_all", ctahanan.GetAll)
	tahananRoute.GET("/get_by_id/:id", ctahanan.GetByID)
	tahananRoute.POST("/create", ctahanan.Create)
	tahananRoute.PUT("/update/:id", ctahanan.Update)
	tahananRoute.DELETE("/delete/:id", ctahanan.Delete)

	sipirRoute := server.Group("/sipir")

	sipirRoute.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))
	sipirRoute.Use(middlewares.CheckTokenMiddlewareUser)

	sipirRoute.GET("/get_all", csipir.GetAll)
	sipirRoute.GET("/get_by_id/:id", csipir.GetByID)
	sipirRoute.POST("/create", csipir.Create)
	sipirRoute.PUT("/update/:id", csipir.Update)
	sipirRoute.DELETE("/delete/:id", csipir.Delete)

	// Auth User
	server.POST("/user/register", cuser.RegisterUser)
	server.POST("/user/login", cuser.LoginUser)

	privateUser := server.Group("")

	privateUser.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))
	privateUser.Use(middlewares.CheckTokenMiddlewareUser)
	privateUser.POST("user/logout", cuser.LogoutUser)

}
