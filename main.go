package main

import (
	"aries-financial-challenge/controllers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/analyze", analyzeHandler)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	controllers.AnalysisHandler(w, r)
}
