package main

import (
	"github.com/gin-gonic/gin"
	"url-shortner/middleware"
	"url-shortner/controllers"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.NoRoute(controllers.NoRoute)
	router.GET("/login",controllers.HandleLogin)
	router.POST("/login", controllers.UserLogin)
	router.GET("/",controllers.HandleHomePage)
	router.GET("/:username/:shortcut",controllers.HandleUrl)
	router.POST("/signup",controllers.HandleSignup)
	router.GET("/dashboard", middleware.Authenticate() , controllers.HandleDashboard)
	router.POST("/addurl",middleware.Authenticate(),controllers.Addurl)
	router.POST("/deleteurl",controllers.HandleDeleteUrl)
	router.Run(":8080")
}



/*
*/
