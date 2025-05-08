package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/msillz/go-have-fun/pkg/morestrings"
	"rsc.io/quote"
)

func main() {

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		Protocol: clickhouse.Native, // or clickhouse.HTTP
	})

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": morestrings.Hello(quote.Go())})
	})

	http.ListenAndServe(":8080", nil)
}
