package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type testStruct struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

var testData = []testStruct{
	{Title: "Test Title 1", Message: "This is a test message to see if this works"},
	{Title: "Test Title 2", Message: "If you see this, then the backend has been set up correctly"},
}

func main() {
	router := gin.Default()

	router.GET("/test", testGet)
	router.POST("/test", testPost)

	router.Run("localhost:8080")
	fmt.Println("Hello! This is working correctly!")
}

func testGet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testData)
}

func testPost(c *gin.Context) {
	var newTest testStruct

	if err := c.BindJSON(&newTest); err != nil {
		return
	}

	testData = append(testData, newTest)
	c.IndentedJSON(http.StatusOK, newTest)
}
