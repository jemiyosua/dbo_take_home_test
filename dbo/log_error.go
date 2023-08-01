package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func LogError(Page string, BodyJson string, JsonReturn string, ErrorCode string, ErrorCodeReturn string, ErrorMessage string, ErrorMessageReturn string, Source string, c *gin.Context) {

	query := fmt.Sprintf("INSERT INTO dbo_log_error (page, body_json, json_return, query, errorcode, errorcode_return, errormessage, errormessage_return, source) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", Page, BodyJson, JsonReturn, ErrorCode, ErrorCodeReturn, ErrorMessage, ErrorMessageReturn, Source)
	_, err := db.Exec(query)
	if err != nil {
		errorMessage := "Error query, " + err.Error() + " || " + query
		fmt.Println(errorMessage)
		// errorMessageReturn := "Gagal insert data dbo_customer"
		// returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
		// return
	}
}
