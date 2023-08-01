package main

import (
	"database/sql"
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	db *sql.DB
)

func main() {
	router := gin.Default()

	db = DB()

	apiVersion := "/api/v1/"

	router.POST(apiVersion+"Login", Login)
	router.POST(apiVersion+"GetDataCustomer", GetDataCustomer)
	// router.POST(apiVersion+"LaporanC", LaporanC)

	router.Run(":1000")
}

func auth(c *gin.Context) (string, string) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			errorMessage := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			fmt.Println(errorMessage)
			return "1", nil
		}

		SecretKey := os.Getenv("SecretKey")
		return []byte(SecretKey), nil
	})

	if token != nil && err == nil {
		return "0", "token verified"
	} else {
		return "1", "not authorized, " + err.Error()
	}
}
