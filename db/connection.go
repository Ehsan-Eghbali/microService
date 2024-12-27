package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB(dbURL string) *sql.DB {
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("Could not connect to the database: %s", err.Error())
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Database connection is not alive: %s", err.Error())
    }

    log.Println("Connected to the database")
    return db
}
