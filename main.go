package main

import "github.com/rumbel/belajar/internal/app"

func main() {
	app := app.NewApp()
	// defer app.Close()
	app.Run()
}
