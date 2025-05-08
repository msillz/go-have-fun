package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/msillz/go-have-fun/pkg/morestrings"
	"rsc.io/quote"
)

func main() {

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": quote.Go()})
	})

	http.ListenAndServe(":8080", nil)
}
