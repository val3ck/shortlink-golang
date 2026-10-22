package main

import (
	"fmt"
	rest "golang-api/internal/handler"
	"net/http"
)

func main() {
	our_mux := http.NewServeMux()
	rest.RegisterRoutes(our_mux)

	fmt.Print("Server successfully started.\n")
	http.ListenAndServe(":8080", our_mux)
}
