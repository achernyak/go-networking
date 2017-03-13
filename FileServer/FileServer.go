package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("/home/httpd/html"))

	err := http.ListenAndServe(":8000", fileServer)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
