package main

import (
	"log"
	"simple-social-media-API/config"
	cd "simple-social-media-API/features/comment/data"
	ch "simple-social-media-API/features/comment/handler"
	cs "simple-social-media-API/features/comment/services"
	pd "simple-social-media-API/features/post/data"
	phdl "simple-social-media-API/features/post/handler"
	psrv "simple-social-media-API/features/post/services"
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

	postData := pd.Isolation(db)
	postSrv := psrv.Isolation(postData)
	postHdl := phdl.Isolation(postSrv)

	commentData := cd.New(db)
	commentSrv := cs.New(commentData)
	commentHdl := ch.New(commentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	// e.Static("/files", "./files")
	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Deactive(), middleware.JWT([]byte(config.JWT_KEY)))

	//posting
	e.POST("/posts", postHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/posts/:id", postHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/posts/:id", postHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/posts", postHdl.MyPosts(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/posts/:id", postHdl.GetPostById(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/allposts", postHdl.AllPosts())

	e.GET("/comments", commentHdl.GetComments())
	e.POST("/comments", commentHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
