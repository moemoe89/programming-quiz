package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"net/http"
	"strings"
	"time"
)

func RemoveIndex(s []string, value string) []string {
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == value {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

var bracketPair = map[string]string{
	"]": "[",
	")": "(",
	"}": "{",
	">": "<",
}

func CheckBalance(input string)(string){

	var brackets []string

	s := strings.Split(input,"")
	if len(s)%2 != 0 {
		return "No"
	}

	for _,j := range s {

		if j == "[" || j == "(" || j == "{" || j == "<" {
			brackets = append(brackets,j)
		} else {
			if j == "]" || j == ")" || j == "}" || j == ">" {
				lastBrackets := len(brackets) - 1
				if len(brackets) < 1 || brackets[lastBrackets] != bracketPair[j] {
					return "No"
				} else {
					brackets = RemoveIndex(brackets, bracketPair[j])
				}
			}
		}

	}

	return "Yes"
}

func Index(c *gin.Context) {

	var htmlResult string

	input := c.Query("input")

	if input != "" {
		htmlResult = CheckBalance(input)
	}

	html := `
	<html>
		<title>Parentheses Balance Checker</title>
		<center>
			<h1>Check the brackets string</h1>
			<form action="" method="GET">
				<input type="text" name="input" value="`+c.Query("input")+`">
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