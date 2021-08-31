package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/mail":  "https://mail.google.com/",
		"/notes": "https://keep.google.com/",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /mail
  url: https://mail.google.com
- path: /notes
  url: https://keep.google.com
`
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
