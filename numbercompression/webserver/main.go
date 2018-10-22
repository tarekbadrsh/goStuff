package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tarekbadrshalaan/goStuff/numbercompression"
)

func main() {
	err := http.ListenAndServe(":5050", handler())
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/compress", compressHandler)
	r.HandleFunc("/uncompress", uncompressHandler)
	return r
}

func compressHandler(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("v")
	if s == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	number, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		http.Error(w, "not a number: "+s, http.StatusBadRequest)
		return
	}
	text := numbercompression.CompresNumberDefault(number)
	fmt.Fprintln(w, text)
}

func uncompressHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("v")
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}
	num := numbercompression.UncompresNumberDefault(text)
	fmt.Fprintln(w, num)
}
