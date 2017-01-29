package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"net/http"
	"strconv"
	"strings"
	"time"
)

func Index(c *gin.Context) {

	var err error
	var htmlResult string
	var selectedTrue string
	var selectedFalse string
	var useCaseSensitive bool

	firstInput := c.Query("first_input")
	secondInput := c.Query("second_input")
	caseSensitive := c.Query("case_sensitive")

	if firstInput != "" || secondInput != "" {
		if caseSensitive != "" {
			useCaseSensitive, err = strconv.ParseBool(caseSensitive)
			if err != nil {
				htmlResult = `<br /><h2>Something went wrong, please try again.</h2>`
			}
		}

		if useCaseSensitive == false {
			firstInput = strings.ToLower(firstInput)
			secondInput = strings.ToLower(secondInput)
			selectedFalse = "selected"
		} else {
			selectedTrue = "selected"
		}

		s1 := strings.Split(firstInput,",")
		s2 := strings.Split(secondInput,",")

		result := []string{}

		check := []bool{}
		for i,j := range s1 {
			check = append(check,false)
			for _,k := range s2 {
				if j == k {
					check[i] = true
				}
			}

			if check[i] == false {
				result = append(result,j)
			}
		}

		htmlResult = strings.Join(result,",")
	}

	html := `
	<html>
		<title>Missing Array Value</title>
		<center>
			<h1>Check Missing Value of Array</h1>
			<form action="" method="GET">
				First Array : <input type="text" name="first_input" value="`+c.Query("first_input")+`"><br />
				Second Array : <input type="text" name="second_input" value="`+c.Query("second_input")+`"><br />
				<select name="case_sensitive">
					<option value="true" `+selectedTrue+`>True</option>
					<option value="false" `+selectedFalse+`>False</option>
				</select>
				<input type="submit" value="Check">
			</form>
			`+htmlResult+`
		</center>
	</html>`

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(html))
}

func main() {

	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, PUT, PATCH, HEAD, DELETE",
		RequestHeaders:  "Authorization, Content-Type, Content-Length, Expected, Transfer-Encoding, X-Requested-With",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: true,
		MaxAge:          60 * time.Second,
	}))

	router.GET("/",Index)
	router.Run()
}