package main

import (
	"provi/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run()
}
