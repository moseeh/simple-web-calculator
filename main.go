package main

import (
	"calc/utilities"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))

	http.HandleFunc("/", utilities.FormHandler)
	http.HandleFunc("/calculate", utilities.CalculateHandler)

	fmt.Println("Server started at :8000")
	http.ListenAndServe(":8000", nil)
}
