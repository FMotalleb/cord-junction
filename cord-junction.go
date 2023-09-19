package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request headers:\n")
		fmt.Printf("%s", r.Host)
		for header, value := range r.Header {
			fmt.Printf("%s: %s\n", header, value)
		}
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}
}
