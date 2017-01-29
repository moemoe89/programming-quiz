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

	input := c.Query("input")
	caseSensitive := c.Query("case_sensitive")

	if input != "" {
		if caseSensitive != "" {
			useCaseSensitive, err = strconv.ParseBool(caseSensitive)
			if err != nil {
				htmlResult = `<br /><h2>Something went wrong, please try again.</h2>`
			}
		}

		if useCaseSensitive == false {
			input = strings.ToLower(input)
			selectedFalse = "selected"
		} else {
			selectedTrue = "selected"
		}

		s := strings.Split(input,"")

		encountered := map[string]bool{}
		result := []string{}

		for v := range s {
			if encountered[s[v]] != true {
				encountered[s[v]] = true
				result = append(result, s[v])
			}
		}

		htmlResult = strings.Join(result,"")
	}

	html := `
	<html>
		<title>Duplicate Character Checker</title>
		<center>
			<h1>Remove Duplicate Character</h1>
			<form action="" method="GET">
				<input type="text" name="input" value="`+c.Query("input")+`">
				<select name="case_sensitive">
					<option value="true" `+selectedTrue+`>True</option>
					<option value="false" `+selectedFalse+`>False</option>
				</select>
				<input type="submit" value="Remove">
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