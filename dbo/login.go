package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type JLoginRequest struct {
	Username string
	Password string
}

func Login(c *gin.Context) {

	var (
		reqBody   JLoginRequest
		bodyBytes []byte
		Token     string
	)

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)

	// ------ Body Json Validation ------
	if string(bodyString) == "" {
		errorMessage := "Error, Body is empty"
		returnDataJsonLogin("1", errorMessage, Token, c)
		return
	}

	is_Json := isJSON(bodyString)
	if is_Json == false {
		errorMessage := "Error, Body - invalid json data"
		returnDataJsonLogin("1", errorMessage, Token, c)
		return
	}
	// ------ end of Body Json Validation ------

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		errorMessageReturn := "Error, Bind Json Data"
		errorMessage := err.Error()
		returnDataJsonLogin("1", errorMessageReturn+" | "+errorMessage, Token, c)
		return
	} else {
		Username := reqBody.Username
		Password := reqBody.Password
		errorMessage := ""

		// ------ Param Validation ------
		if Username == "" {
			errorMessage = errorMessage + "\n- " + "Username cannot be null"
		}

		if Password == "" {
			errorMessage = errorMessage + "\n- " + "Password cannot be null"
		}

		if errorMessage != "" {
			returnDataJsonLogin("1", errorMessage, Token, c)
			return
		}
		// ------ end of Param Validation ------

		Cnt := 0
		query := "SELECT COUNT(1) AS cnt FROM dbo_login WHERE username = '" + Username + "' AND password = '" + Password + "' "
		if err := db.QueryRow(query).Scan(&Cnt); err != nil {
			errorMessage := "Error query, " + err.Error()
			returnDataJsonLogin("1", errorMessage, Token, c)
			return
		}

		if Cnt > 0 {
			SecretKey := os.Getenv("SecretKey")
			sign := jwt.New(jwt.GetSigningMethod("HS256"))
			Token, err = sign.SignedString([]byte(SecretKey))
			if err != nil {
				errorMessage := "Error, " + err.Error()
				returnDataJsonLogin("1", errorMessage, Token, c)
				return
			}

			returnDataJsonLogin("0", "", Token, c)
		} else {
			errorMessage = "Username or Password invalid, please try again!"
			returnDataJsonLogin("1", errorMessage, Token, c)
		}
	}
}

func returnDataJsonLogin(ErrorCode string, ErrorMessage string, Token string, c *gin.Context) {

	if strings.Contains(ErrorMessage, "Error running") {
		ErrorMessage = "Error Execute data"
	}

	if ErrorCode == "504" {
		c.String(http.StatusUnauthorized, "")
		return
	} else {
		c.PureJSON(http.StatusOK, gin.H{
			"ErrorCode":    ErrorCode,
			"ErrorMessage": ErrorMessage,
			"Token":        Token,
		})
		return
	}
}
