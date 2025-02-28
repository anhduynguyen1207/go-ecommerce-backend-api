package main

import "github.com/anhduynguyen1207/go-ecommerce-backend-api/internal/initialize"

func main() {
	// r := routers.NewRoute()
	// r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r := initialize.Run()
	r.Run(":8002")
}
