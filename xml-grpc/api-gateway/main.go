package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    fmt.Println("In go")
    // create a new Gorilla/mux router
    r := mux.NewRouter()

    // define a handler for the root route "/"
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("in handlefunc")
        w.Write([]byte("Hello, World!"))
    })

    // start the server
    log.Fatal(http.ListenAndServe(":8010", r))
}
