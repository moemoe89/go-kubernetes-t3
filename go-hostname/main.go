package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		name = err.Error()
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(name+"\n"))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	http.ListenAndServe(":5000", nil)
}