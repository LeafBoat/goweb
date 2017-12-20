package main

import (
	"log"
	"net/http"
)

func main() {
	FormPage()
	HanderUpload()
	var err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
