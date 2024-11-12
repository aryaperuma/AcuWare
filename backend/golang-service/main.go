package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    _ "github.com/lib/pq"
)

const (
    host     = "postgres"
    port     = 5432
    user     = "user"
    password = "password"
    dbname   = "inventory"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Inventory Endpoint: Go Service Connected!")
    })

    log.Println("Golang service running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
