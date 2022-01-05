package main

import (
	"encoding/json"
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
		return
	}
	n, err1 := strconv.Atoi(r.FormValue("n"))
	k, err2 := strconv.Atoi(r.FormValue("k"))
	if err1 != nil || err2 != nil {
		fmt.Fprintln(w, "Invalid input for n or k")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	scales := scalesalgorithm.GetScales(n, k, k)
	response := make(map[string][][]int)
	response["scales"] = scales
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintln(w, "JSON error", err)
	}
	w.Write(jsonResponse)
	return
}
