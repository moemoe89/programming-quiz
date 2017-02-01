package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CountNumberAppear(min string,max string, num string)(*int,error){

	m := make(map[int]bool)

	minInt,err := strconv.Atoi(min)
	if err != nil {
		return nil,err
	}

	maxInt,err := strconv.Atoi(max)
	if err != nil {
		return nil,err
	}

	_,err = strconv.Atoi(num)
	if err != nil {
		return nil,errors.New("Number appear should be number")
	}

	sNum := strings.Split(num,"")
	lenNum := len(sNum)

	if lenNum < 1 {
		return nil,errors.New("Number appear can't be empty")
	}

	for i := minInt; i <= maxInt; i ++ {

		s := strings.Split(strconv.Itoa(i),"")
		lenS := len(s)
		if lenS < lenNum {
			continue
		}

		for j,k := range s {
			if k == sNum[0] {
				var numberAppear bool
				for x := 1; x <= lenNum-1; x ++ {
					if ((j+x) <= (lenS-1) ) && (s[j+x] == sNum[x]) {
						numberAppear = true
					}
				}
				if numberAppear == true || lenNum == 1 {
					m[i] = true
				}
			}
		}
	}

	var count int
	for _, _ = range m {
		count++
	}

	return &count,nil
}

func Index(c *gin.Context) {

	var htmlResult string

	fromNumber := c.Query("from_number")
	toNumber := c.Query("to_number")
	numberAppear := c.Query("number_appear")

	if fromNumber != "" || toNumber != "" || numberAppear != "" {
		result,err := CountNumberAppear(fromNumber,toNumber,numberAppear)
		if err != nil {
			htmlResult = `<br /><h2>`+err.Error()+`</h2>`
		} else {
			htmlResult = `Number ` + numberAppear + ` appear : ` + strconv.Itoa(*result) + ` times.`
		}
	}

	html := `
	<html>
		<title>Count Number Appear</title>
		<center>
			<h1>Count Number Apper</h1>
			<form action="" method="GET">
				From Number : <input type="text" name="from_number" value="`+c.Query("from_number")+`"><br />
				To Number : <input type="text" name="to_number" value="`+c.Query("to_number")+`"><br />
				Number Appear : <input type="text" name="number_appear" value="`+c.Query("number_appear")+`"><br />
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