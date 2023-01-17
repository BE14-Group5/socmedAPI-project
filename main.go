package main

import (
	"log"
	"simple-social-media-API/config"
	ud "simple-social-media-API/features/user/data"
	uh "simple-social-media-API/features/user/handler"
	us "simple-social-media-API/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := ud.New(db)
	userSrv := us.New(userData)
	userHdl := uh.New(userSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.Static("/files", "./files")
	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
