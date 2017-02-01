package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SplitMiddleCharacters(input string)(string,string){

	s := strings.Split(input,"")

	var firstString string
	var secondString string

	middleLenSplit := len(s)/2
	for i,j := range s {
		k := i + 1
		if k <= middleLenSplit {
			firstString = firstString + j
		} else {
			secondString = secondString + j
		}
	}

	return firstString,secondString
}


func MagicNumber(s string)(*string,error){

	n, err := strconv.Atoi(s)
	if err != nil {
		return nil,err
	}

	if n%2 != 0 {
		return nil,errors.New("Not even number")
	}

	nineString := fmt.Sprintf("%0"+s+"d", 9)
	nineString = strings.Replace(nineString, "0", "9", -1)

	minInteger := 0
	maxInteger,err := strconv.Atoi(nineString)
	if err != nil {
		return nil,err
	}

	var result []string
	for i := minInteger; i <= maxInteger; i++ {
		formatted := fmt.Sprintf("%0"+s+"d", i)

		firstNumber,secondNumber := SplitMiddleCharacters(formatted)

		firstNumberInt,err := strconv.Atoi(firstNumber)
		if err != nil {
			return nil,err
		}

		secondNumberInt,err := strconv.Atoi(secondNumber)
		if err != nil {
			return nil,err
		}

		concatNumber := firstNumber + secondNumber
		concatNumberInt,err := strconv.Atoi(concatNumber)
		if err != nil {
			return nil,err
		}

		additionNumber := firstNumberInt + secondNumberInt
		additionNumberSquare := additionNumber * additionNumber

		if additionNumberSquare == concatNumberInt {
			result = append(result,formatted)
		}

	}

	var showResult string
	lastElement := len(result) - 1
	for x,y:= range result {
		if x == 0 {
			showResult = showResult + y
		} else if x ==  lastElement {
			showResult = showResult + " and " + y
		} else {
			showResult = showResult + ", " + y
		}
	}

	return &showResult,nil
}

func Index(c *gin.Context) {

	var htmlResult string

	input := c.Query("input")

	if input != "" {
		result,err := MagicNumber(input)
		if err!= nil {
			htmlResult = `<br /><h2>Something went wrong, please try again.</h2>`
		} else {
			htmlResult = *result
		}
	}

	html := `
	<html>
		<title>Magic Number</title>
		<center>
			<h1>Input the number</h1>
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