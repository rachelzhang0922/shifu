package main

import (
	"fmt"
	"net/http"
	"os"
)

func processHello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello_world from device via shifu3!")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, header)
		}
	}
}

func getHumidity(w http.ResponseWriter, req *http.Request) {
	dat, _ := os.ReadFile("raw_data")
	fmt.Print(string(dat))
	fmt.Fprintln(w, string(dat))
}

func main() {
	http.HandleFunc("/hello", processHello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/humidity", getHumidity)

	http.ListenAndServe(":11111", nil)
}
