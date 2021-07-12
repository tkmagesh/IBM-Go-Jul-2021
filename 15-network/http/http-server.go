package main

import (
	"fmt"
	"net/http"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server listening on 8080")
}
