package main

import (
	"myModule/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/add", handlers.Add)
	http.HandleFunc("/subtract", handlers.Subtract)
	http.HandleFunc("/multiply", handlers.Multiply)
	http.HandleFunc("/divide", handlers.Divide)

	_ = http.ListenAndServe(":3000", nil)

}
