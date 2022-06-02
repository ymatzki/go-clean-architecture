package rest

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func Execute() {
	db, err := sql.Open("postgres", `host=localhost port=5432 user=admin password=pass dbname=test sslmode=disable`)
	if err != nil {
		log.Fatalf("unable to use data source name", err)
	}
	defer db.Close()

	// dependency injection
	h := Initialize(db)

	// specify port
	p := 8080

	// execute
	http.HandleFunc("/", h.Get)
	log.Printf("About to listen on %d. Go to http://127.0.0.1:%d/", p, p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", p), nil))
}
