package controllers

import (
	"fmt"
	"net/http"
	"url-shortner/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context){
	c.HTML(http.StatusNotFound , "404.html",nil)
}

func HandleLogin(c *gin.Context){
	c.HTML(http.StatusFound , "login.html",nil)
}

func HandleHomePage(c *gin.Context){
	c.Redirect(http.StatusFound,"/dashboard")
}
func HandleUrl(c *gin.Context){
	username := c.Param("username")
	if !models.UserExist(username)  {
		NoRoute(c)		
	}
	shortcut := c.Param("shortcut")
	link , err := models.GetLink(shortcut,username)
	if err != nil {
		c.HTML(http.StatusNotFound,"404.html",nil)
	}
	c.Redirect(http.StatusFound , link)
	
}

func HandleDashboard(c *gin.Context){
	claims ,  _ := c.Get("claims")

	fmt.Println("Claims:", claims) // Debug logging

	userclaims,ok:= claims.(jwt.MapClaims)		

	if !ok	{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to parse Claims",
		})
		return				
	}

	username , ok := userclaims["sub"].(string)

	if !ok {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error": "Failed to parse the username",
		})
		return
	}
	
	links := models.GetAllLinks(username)
	c.HTML(http.StatusOK,"dashboard.html",gin.H{
		"links": links,
		"username": username,
	})

}
