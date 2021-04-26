package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

type Environment struct {
	DB *pg.DB
}

func (env Environment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Got request")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := env.DB.Ping(ctx); err != nil {
		log.Printf("FAIL: %v", err)
		fmt.Fprint(w, "FAIL")
		fmt.Fprintf(w, "DATABASE_URL=%q", os.Getenv("DATABASE_URL"))
	} else {
		fmt.Fprint(w, "OK")
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	listenAddress := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Addr:     os.Getenv("POSTGRES_ADDR"),
	})
	defer db.Close()

	env := &Environment{
		DB: db,
	}

	log.Printf("start http server at %q", listenAddress)
	if err := http.ListenAndServe(listenAddress, env); err != nil {
		log.Printf("failed to start http server at %q: %v", listenAddress, err)
		return
	}
}
