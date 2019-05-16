package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	log.Print("ready: listening port: ", os.Getenv("PORT"))

	err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
	if err != nil {
		log.Printf("Fail to Listen port")
	}
	defer log.Printf("Application finished")
}
