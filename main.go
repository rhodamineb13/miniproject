package main

import (
	"fmt"
	"miniproject/common/config"
	"miniproject/common/shared"
	"net/http"
)

func main() {
	r := shared.Route()

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	conf := config.NewConfig()
	fmt.Println(conf)

	s.ListenAndServe()

}
