package controllers

import (
	"net/http"
	"url-shortner/models"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)
var jwtKey = []byte("secret")

func UserLogin(c *gin.Context){
    username := c.PostForm("username")
		password := c.PostForm("password")
		if models.UserAuthenticate(username,password) {
			//user is authencitated on the backend
			// c.JSON(http.StatusOK, gin.H{"message": "Login successful"}) // no need to send this

			//generate a jwt token and save it as a cookie 	

			expirationTime := time.Now().Add(24 * time.Hour)

			claims := jwt.StandardClaims{
				Subject:username,
				ExpiresAt: expirationTime.Unix(),	
			}		
			
			token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
			tokenString , err := token.SignedString(jwtKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":"could not create a jwt token",
				})
				return 
			}

			c.SetCookie("token",tokenString,int(expirationTime.Unix()),"/","localhost",false,true)
			c.Redirect(http.StatusFound,"/dashboard")

		}else{
			c.HTML(http.StatusOK, "login.html", gin.H{
				"LoginFailed": true,
			})	
	}			
}

func HandleSignup(c *gin.Context){
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || username == "" || password == "" {
			c.JSON(http.StatusOK, gin.H{"message": "Singup failed"})
			return 
	}
	if models.UserExist(username) {
			c.JSON(http.StatusOK, gin.H{"message": "user already exists"})
			return 
	} 
	if models.UserExistEmail(email){
			c.JSON(http.StatusOK, gin.H{"message": "user already exists"})
			return
	}
	fmt.Println()
}

