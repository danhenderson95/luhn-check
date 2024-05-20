package main

import (
	"danhenderson95/luhncheck/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.CheckHandler)
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		log.Fatal(err)
	}

}
