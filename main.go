package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Println(t.Month(), "Request Efetuado")
		w.Write([]byte("<h1>Hello From GoCD  Blefari </h1>"))
	})

	http.ListenAndServe(":8080", nil)
}
