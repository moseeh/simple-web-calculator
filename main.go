package main

import (
	"calc/utilities"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", utilities.FormHandler)
	http.HandleFunc("/calculate", utilities.CalculateHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
