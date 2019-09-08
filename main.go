package main

import "github.com/maxvoronov/otus-go-10-logger/app"

func main() {
	application, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	application.Start()
}
