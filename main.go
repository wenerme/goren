package main

import (
	"os"
)

func main() {
	app := NewGorenApp()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
