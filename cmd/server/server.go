package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/ryohei1216/go-mvc/cmd/router"
	"github.com/ryohei1216/go-mvc/db"
)

func main() {
	db, err := db.New()
	if err != nil {
		slog.Error(err.Error())
		panic("failed to connect database")
	}

	r := router.New(db)

	log.Fatal(http.ListenAndServe(":8080", r))
}
