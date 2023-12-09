package main

import (
	"miniproject/common/shared"
	"net/http"
)

func main() {
	r := shared.Route()

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	s.ListenAndServe()
}
