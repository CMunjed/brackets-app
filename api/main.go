package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

var testData = []testStruct{
	{Title: "Test Title 1", Message: "This is a test message to see if this works"},
	{Title: "Test Title 2", Message: "If you see this, then the backend has been set up correctly"},
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", testGet).Methods("GET")
	router.HandleFunc("/test/:title", testGetTitle).Methods("GET")
	router.HandleFunc("/test", testPost).Methods("POST")

	router.Run("localhost:8080")
}

func testGet(c *mux.) {
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

func testGetTitle(c *gin.Context) {
	inputTitle := c.Param("title")

	for _, data := range testData {
		if data.Title == inputTitle {
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found! :("})
}

func testPut(c *gin.Context) {
	targetTitle := c.Param("title")
	var editedData testStruct

	for _, data := range testData {
		if data.Title == targetTitle {
			if err := c.BindJSON(&editedData); err != nil {
				return
			}
			c.IndentedJSON(http.StatusOK, editedData)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found! :("})
}