package main

import "blogApp/internal/http/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":3000")
}
