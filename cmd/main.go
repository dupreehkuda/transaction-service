package main

import "github.com/dupreehkuda/transaction-service/internal/api"

func main() {
	srv := api.NewByConfig()
	srv.Run()
}
