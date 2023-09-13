package main

import (
	"log/slog"
	"net/http"

	"github.com/ryohei1216/go-mvc/cmd/router"
	"github.com/ryohei1216/go-mvc/db"
)

func main() {
	db, err := db.New()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	r := router.New(db)

	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
