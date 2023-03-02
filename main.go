package main

import (
	"fmt"

	"example.com/api/app"
)

func main() {
	fmt.Println("Hello! If you see this, then I am working!")
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
