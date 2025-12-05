package main 

import (
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	// handle root path

	http.HandleFunc("/", func(rw http.ResponseWriter,  r*http.Request) {
        log.Println("Hello, World!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Failed to read request body", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Received request with body: %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter,  *http.Request) {
        log.Println("Goodbye, World!")
	})

	http.ListenAndServe(":9090", nil)
}
