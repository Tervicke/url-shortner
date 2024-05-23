package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"url-shortner/models"
)

func Addurl(c *gin.Context){
	claims ,  _ := c.Get("claims")
  ShortUrl := c.PostForm("shortUrl")
	LongUrl := c.PostForm("longUrl")
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

	fmt.Println(username)
	fmt.Print(ShortUrl)
	fmt.Print(LongUrl)
	_,err := models.Addurl(ShortUrl , LongUrl , username)
	if err != nil {
		c.HTML(http.StatusFound,"dashboard.html",gin.H{
			"links": models.GetAllLinks(username),
			"Error": true,
			"Message": "Shorturl already in use , Try something else !",
		})
		/*
		c.JSON(http.StatusFound,gin.H{
			"message": "Error",
		})
		*/
	}
	c.Redirect(http.StatusFound,"/dashboard")
}
func HandleDeleteUrl(c *gin.Context){
	ID := c.PostForm("ID") 
	_ , err := models.Deleteurl(ID)
	if err != nil {
		c.HTML(http.StatusFound,"dashboard.html",gin.H{
			"Error": true,
			"Message": "something went wrong please try agian later",
		})
	}
	c.Redirect(http.StatusFound,"/dashboard")
}
