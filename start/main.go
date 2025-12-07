package main 

import (
	"microservices/handlers"
	"net/http"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
    hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/hello", hh)

	http.ListenAndServe(":9090", sm)
}
