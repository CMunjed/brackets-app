package main

import (
	"fmt"

	"example.com/api/app"
)

func main() {
	fmt.Println("Hello! This is working correctly!")
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
