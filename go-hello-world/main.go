//
//  Practicing Kubernetes
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Pong</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	port := "8782"
	fmt.Println("Server running on :"+port)
	http.ListenAndServe(":"+port, nil)
}