package main

import (
	"fmt"
	"net/http"

	"coding/session/internal"
)

func main() {
	s := internal.New()
	fmt.Println("Server started at", s.Addr())
	err := http.ListenAndServe(s.Addr(), s.Handler())
	if err != nil {
		fmt.Println("[error] starting app:", err)
	}
}
