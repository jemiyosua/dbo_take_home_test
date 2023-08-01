package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type JGetDataCustomerRequest struct {
	Method       string
	Id           string
	Nama         string
	Password     string
	Email        string
	NomorHP      string
	JenisKelamin string
	TanggalLahir string
	KotaDomisili string
	StatusAkun   int
	Page         int
	RowPage      int
	OrderBy      string
	Order        string
}

type JGetDataCustomerResponse struct {
	Id           string
	Nama         string
	Email        string
	NomorHP      string
	JenisKelamin string
	TanggalLahir string
	KotaDomisili string
	TanggalInput string
	StatusAkun   string
}

func GetDataCustomer(c *gin.Context) {

	var (
		reqBody          JGetDataCustomerRequest
		resBody          JGetDataCustomerResponse
		ListDataCustomer []JGetDataCustomerResponse
		bodyBytes        []byte
		totalPage        float64
		totalRecords     float64
	)

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)

	PageGO := "DATA CUSTOMER"

	// ------ Body Json Validation ------
	if string(bodyString) == "" {
		errorMessage := "Error, Body is empty"
		errorMessageReturn := "Error, Body is empty"
		LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
		returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
		return
	}

	is_Json := isJSON(bodyString)
	if is_Json == false {
		errorMessage := "Error, Body - invalid json data"
		errorMessageReturn := "Error, Body - invalid json data"
		LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
		returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
		return
	}
	// ------ end of Body Json Validation ------

	ErrorCodeAuth, ErrorMessageAuth := auth(c)

	if ErrorCodeAuth == "0" {
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			errorMessageReturn := "Error, Bind Json Data"
			errorMessage := err.Error()
			LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
			returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
			return
		} else {
			Method := reqBody.Method
			Id := reqBody.Id
			Nama := reqBody.Nama
			Email := reqBody.Email
			Password := reqBody.Password
			NomorHP := reqBody.NomorHP
			JenisKelamin := reqBody.JenisKelamin
			TanggalLahir := reqBody.TanggalLahir
			KotaDomisili := reqBody.KotaDomisili
			StatusAkun := reqBody.StatusAkun
			Page := reqBody.Page
			RowPage := reqBody.RowPage
			OrderBy := reqBody.OrderBy
			Order := reqBody.Order

			if Method == "" {
				errorMessageReturn := "Method cannot null"
				errorMessage := "Method cannot null"
				LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
				returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
				return
			} else {
				if Method == "INSERT" {

					CountUser := 0
					query := "SELECT COUNT(1) AS cnt FROM dbo_customer WHERE email = '" + Email + "'"
					if err := db.QueryRow(query).Scan(&CountUser); err != nil {
						errorMessage := "Error query, " + err.Error() + " || " + query
						errorMessageReturn := "Gagal get data dbo_customer"
						LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
						returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
						return
					}

					if CountUser > 0 {
						errorMessage := ""
						errorMessageReturn := "Data Customer sudah ada"
						LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
						returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
						return
					} else {
						query := fmt.Sprintf("INSERT INTO dbo_customer(nama, password, email, nomor_hp, jenis_kelamin, tanggal_lahir, kota_domisili, status_akun) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)", Nama, Password, Email, NomorHP, JenisKelamin, TanggalLahir, KotaDomisili, StatusAkun)
						_, err := db.Exec(query)
						if err != nil {
							errorMessage := "Error query, " + err.Error() + " || " + query
							errorMessageReturn := "Gagal insert data dbo_customer"
							LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
							returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
							return
						}

						returnGetDataCustomer(ListDataCustomer, "0", "0", "", "Berhasil insert data customer", totalPage, totalRecords, c)
						return
					}

				} else if Method == "UPDATE" {

					queryUpdate := ""
					if Nama != "" {
						queryUpdate = queryUpdate + fmt.Sprintf(", nama = '%s' ", Nama)
					}
					if Email != "" {
						queryUpdate = queryUpdate + fmt.Sprintf(", email = '%s' ", Email)
					}
					if Password != "" {
						queryUpdate = queryUpdate + fmt.Sprintf(", password = '%s' ", Password)
					}
					if NomorHP != "" {
						queryUpdate = queryUpdate + fmt.Sprintf(", nomor_hp = '%s' ", NomorHP)
					}
					if JenisKelamin != "" {
						queryUpdate = queryUpdate + fmt.Sprintf(", jenis_kelamin = '%s' ", JenisKelamin)
					}
					if TanggalLahir != "" {
						queryUpdate = queryUpdate + fmt.Sprintf(", tanggal_lahir = '%s' ", TanggalLahir)
					}
					if KotaDomisili != "" {
						queryUpdate = queryUpdate + fmt.Sprintf(", kota_domisili = '%s' ", KotaDomisili)
					}

					query := fmt.Sprintf("UPDATE dbo_customer SET tgl_update = NOW() %s WHERE id = '%s'", queryUpdate, Id)
					_, err := db.Exec(query)
					if err != nil {
						errorMessage := "Error query, " + err.Error()
						errorMessageReturn := "Gagal update data dbo_customer"
						LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
						returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
						return
					}

					returnGetDataCustomer(ListDataCustomer, "0", "0", "", "Berhasil update data customer", totalPage, totalRecords, c)
					return

				} else if Method == "DELETE" {

				} else if Method == "SELECT" {

					PageNow := (Page - 1) * RowPage

					queryWhere := ""
					if Nama != "" {
						if queryWhere != "" {
							queryWhere += " AND "
						}

						queryWhere += " nama LIKE '%" + Nama + "%' "
					}

					if queryWhere != "" {
						queryWhere = " WHERE " + queryWhere
					}

					queryOrder := ""
					if OrderBy != "" {
						queryOrder = fmt.Sprintf(" ORDER BY %s %s", OrderBy, Order)
					}

					totalRecords = 0
					totalPage = 0
					query := fmt.Sprintf("SELECT COUNT(1) AS cnt FROM dbo_customer %s", queryWhere)
					if err := db.QueryRow(query).Scan(&totalRecords); err != nil {
						errorMessage := "Error query, " + err.Error()
						errorMessageReturn := "Gagal get data dbo_customer"
						LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
						returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
						return
					}
					totalPage = math.Ceil(float64(totalRecords) / float64(RowPage))

					query1 := fmt.Sprintf(`SELECT id, IFNULL(nama, ''), IFNULL(email, ''), IFNULL(nomor_hp, ''), IFNULL(jenis_kelamin, ''), IFNULL(tanggal_lahir, ''), IFNULL(kota_domisili, ''), IFNULL(tgl_input, ''), IFNULL(status_akun, '') FROM dbo_customer %s %s LIMIT %d,%d;`, queryWhere, queryOrder, PageNow, RowPage)
					rows, err := db.Query(query1)
					defer rows.Close()
					if err != nil {
						errorMessage := fmt.Sprintf("Error running %q: %+v", query1, err)
						errorMessageReturn := "Gagal get data dbo_customer"
						LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
						returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
						return
					}
					JenisKelamin := ""
					for rows.Next() {
						err = rows.Scan(
							&resBody.Id,
							&resBody.Nama,
							&resBody.Email,
							&resBody.NomorHP,
							&JenisKelamin,
							&resBody.TanggalLahir,
							&resBody.KotaDomisili,
							&resBody.TanggalInput,
							&resBody.StatusAkun,
						)

						if JenisKelamin == "L" {
							resBody.JenisKelamin = "Pria"
						} else if JenisKelamin == "P" {
							resBody.JenisKelamin = "Pria"
						} else {
							resBody.JenisKelamin = ""
						}

						ListDataCustomer = append(ListDataCustomer, resBody)

						if err != nil {
							errorMessage := fmt.Sprintf("Error running %q: %+v", query1, err)
							errorMessageReturn := "Gagal get data dbo_customer"
							LogError(PageGO, bodyString, "", "1", "1", errorMessage, errorMessageReturn, "GO", c)
							returnGetDataCustomer(ListDataCustomer, "1", "1", errorMessage, errorMessageReturn, totalPage, totalRecords, c)
							return
						}
					}

					returnGetDataCustomer(ListDataCustomer, "0", "0", "", "", totalPage, totalRecords, c)
					return
				}
			}
		}
	} else {
		returnGetDataCustomer(ListDataCustomer, "1", "1", ErrorMessageAuth, ErrorMessageAuth, totalPage, totalRecords, c)
		return
	}

}

func returnGetDataCustomer(ListDataCustomer []JGetDataCustomerResponse, ErrorCode string, ErrorCodeReturn string, ErrorMessage string, ErrorMessageReturn string, totalPage float64, totalRecords float64, c *gin.Context) {

	if strings.Contains(ErrorMessage, "Error running") {
		ErrorMessage = "Error Execute data"
	}

	if ErrorCode == "504" {
		c.String(http.StatusUnauthorized, "")
		return
	} else {
		c.PureJSON(http.StatusOK, gin.H{
			"ErrorCode":    ErrorCodeReturn,
			"ErrorMessage": ErrorMessageReturn,
			"Result":       ListDataCustomer,
			"TotalPage":    totalPage,
			"TotalRecords": totalRecords,
		})
		return
	}
}
