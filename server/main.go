package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v", r.URL.RawQuery)
	arg := r.FormValue("arg")
	fmt.Fprintf(w, "Form Value arg - %v\n", arg)
	fmt.Fprintln(w, "Hello World!")
}
