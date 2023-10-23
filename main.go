package main

import (
	"ornicode-connect/app"
	"ornicode-connect/initializers"
)

func init() {
	initializers.Load()
}

// @title Connect service api
func main() {
	app.Kernel()
}
