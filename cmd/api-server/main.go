package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/joho/godotenv"
	"github.com/msillz/go-have-fun/pkg/morestrings"
	"rsc.io/quote"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("CLICKHOUSE_HOST")
	dbPassword := os.Getenv("CLICKHOUSE_PASSWORD")

	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr:     []string{dbHost}, // 9440 is a secure native TCP port
		Protocol: clickhouse.Native,
		TLS:      &tls.Config{}, // enable secure TLS
		Auth: clickhouse.Auth{
			Username: "default",
			Password: dbPassword,
		},
	})

	row := conn.QueryRow("SELECT 1")
	var col uint8
	if err := row.Scan(&col); err != nil {
		fmt.Printf("An error while reading the data: %s", err)
	} else {
		fmt.Printf("Result: %d", col)
	}

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": morestrings.Hello(quote.Go())})
	})

	http.ListenAndServe(":8080", nil)
}
