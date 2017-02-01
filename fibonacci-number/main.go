package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"net/http"
	"strconv"
	"time"
)

func FibonacciNumber(s string)(*string,error){

	var fibonacci []int
	var o string
	n,err := strconv.Atoi(s)
	if err != nil {
		return nil,err
	}

	for i := 0; i <= n - 1 ; i++ {

		if i == 0 || i == 1 {
			o += strconv.Itoa(i) + ` `
			fibonacci = append(fibonacci,i)
		} else {
			o += strconv.Itoa(fibonacci[i-1]+fibonacci[i-2]) + ` `
			fibonacci = append(fibonacci,fibonacci[i-1]+fibonacci[i-2])
		}

	}
	return &o,nil
}

func Index(c *gin.Context) {

	var htmlResult string

	input := c.Query("input")

	if input != "" {
		result,err := FibonacciNumber(input)
		if err!= nil {
			htmlResult = `<br /><h2>Something went wrong, please try again.</h2>`
		} else {
			htmlResult = *result
		}
	}

	html := `
	<html>
		<title>Fibonacci Number</title>
		<center>
			<h1>Input the max number</h1>
			<form action="" method="GET">
				<input type="text" name="input" value="`+c.Query("input")+`">
				<input type="submit" value="Show">
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