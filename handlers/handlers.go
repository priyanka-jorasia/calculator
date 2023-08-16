package handlers

import (
	"fmt"
	"myModule/functions"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to Calculator")
	if err != nil {
		fmt.Println(err)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	result := functions.AddValues(100, 200)
	_, _ = fmt.Fprint(w, "Result of 100 + 200 is : ", result)
}

func Subtract(w http.ResponseWriter, r *http.Request) {
	result := functions.SubtractValues(100, 200)
	_, _ = fmt.Fprint(w, "Result of 100-200 is : ", result)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	result1 := functions.MultiplyValues(100, 200)
	_, _ = fmt.Fprint(w, "Result of 100 x 200 is : ", result1)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result := functions.DivideValues(100, 200)
	_, _ = fmt.Fprint(w, "Result of 100/200 is : ", result)
}
