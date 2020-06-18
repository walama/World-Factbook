package main

import (
	"fmt"
	"net/http"
	"os"
	"worldfactbook/scratch_api/handlers"
)

func main() {

	http.HandleFunc("/countries", handlers.CountriesRouter)
	http.HandleFunc("/countries/", handlers.CountriesRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
