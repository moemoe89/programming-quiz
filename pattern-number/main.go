package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"net/http"
	"strconv"
	"time"
)

func LoopToString(min,max,n int)(string){
	var s string
	for i := min; i <= max ; i++ {
		s += strconv.Itoa(n)
	}
	return s
}

func PatternNumber(s string)(*string,error){

	var o string
	n,err := strconv.Atoi(s)
	if err != nil {
		return nil,err
	}

	for i := n; i >= 1 ; i-- {
		o += LoopToString(1,2,i-1)
		o += LoopToString(1,i,i+1)+"<br/>"
	}
	return &o,nil
}

func Index(c *gin.Context) {

	var htmlResult string

	input := c.Query("input")

	if input != "" {
		result,err := PatternNumber(input)
		if err!= nil {
			htmlResult = `<br /><h2>Something went wrong, please try again.</h2>`
		} else {
			htmlResult = *result
		}
	}

	html := `
	<html>
		<title>Pattern Number</title>
		<center>
			<h1>Input the number of loop</h1>
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