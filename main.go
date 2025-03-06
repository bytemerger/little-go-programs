package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/get", handleGet)
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println("Server started at port 4000")
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we go a get request")
}
