package main

import (
	"fmt"
<<<<<<< Updated upstream
	"net/http"

	"github.com/gin-gonic/gin"
=======
	//"gorm.io/gorm"
	//"gorm.io/driver/sqlite"
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"net/http"
	"example.com/api/app"
	"example.com/api/config"
>>>>>>> Stashed changes
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
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
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
