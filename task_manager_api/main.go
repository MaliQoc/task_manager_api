package main

import (
	"restapi_base/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}