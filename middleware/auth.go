package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
var jwtKey = []byte("secret")

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString , err := c.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie{
				/*
				c.JSON(http.StatusUnauthorized,gin.H{
					"error": "unauthorized",
				})
				*/
				c.Abort()
				c.Redirect(http.StatusFound,"/login",)
				return
			}
			c.JSON(http.StatusBadRequest,gin.H{
				"error": "Bad request",
			})
			c.Abort()
			return 
		}
		
		token , err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
			return jwtKey , nil
		})
		
		if err != nil {

			if err ==jwt.ErrSignatureInvalid{
				c.JSON(http.StatusUnauthorized,gin.H{
					"Error": "Unauthorized",
				})
				c.Abort()
				return
			}
			
			c.JSON(http.StatusBadRequest,gin.H{
				"error": "Bad request",
			})
			c.Abort()	
			return

		}

		if !token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}
		claims , ok := token.Claims.(jwt.MapClaims)	
		if !ok {
			c.JSON(http.StatusNotFound,gin.H{
				"error": "failed to parse claims",
			})
		}

		c.Set("claims", claims)

		c.Next()

	}

}
