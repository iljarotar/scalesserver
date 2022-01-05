package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/iljarotar/scalesalgorithm"
)

func main() {
	http.HandleFunc("/", requestHandler)

	fmt.Println("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error: %v", err)
	}
	if n, err1 := strconv.Atoi(r.FormValue("n")); err1 != nil {
		fmt.Fprintln(w, "Invalid input for n")
	} else if k, err2 := strconv.Atoi(r.FormValue("k")); err2 != nil {
		fmt.Fprintln(w, "Invalid input for n")
	} else {
		fmt.Fprintln(w, scalesalgorithm.GetScales(n, k, k))
	}
}
