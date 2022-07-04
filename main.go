package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func buildApp() *gin.Engine {
	r := gin.Default()
	r.GET("/custom-fizzbuzz/:int1/:int2/:limit/:str1/:str2", getResponse)
	return r
}

func main() {
	buildApp().Run(":8080")
}

func getResponse(c *gin.Context) {
	intParam, err := getIntParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "the first 3 parameters should be integers.")
		return
	}
	strParam := getStrParam(c)

	rst := getCustomFizzbuzz(intParam, strParam)
	c.String(http.StatusOK, rst)
}

func getIntParam(c *gin.Context) ([]int, error) {
	param := []string{
		c.Param("int1"),
		c.Param("int2"),
		c.Param("limit"),
	}

	intParam := make([]int, 3)
	for i := 0; i < len(param); i++ {
		v, err := strconv.Atoi(param[i])
		if err != nil {
			return nil, err
		}
		intParam[i] = v
	}

	return intParam, nil
}

func getStrParam(c *gin.Context) []string {
	return []string{
		c.Param("str1"),
		c.Param("str2"),
	}
}

func getCustomFizzbuzz(nums []int, keyword []string) string {
	var s string

	for i := 1; i <= nums[2]; i++ {

		if i%nums[0] == 0 {
			s += fmt.Sprintf("%v", keyword[0])
		}
		if i%nums[1] == 0 {
			s += fmt.Sprintf("%v", keyword[1])
		}
		if i%nums[0] != 0 && i%nums[1] != 0 {
			s += fmt.Sprintf("%v", i)
		}
		s += ","
	}

	return s
}
