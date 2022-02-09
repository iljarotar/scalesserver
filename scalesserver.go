package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/iljarotar/scalesalgorithm"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
