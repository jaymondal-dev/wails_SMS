// SMS/sms/database.go
package sms

import (
    "log"

    "github.com/jmoiron/sqlx"
    _ "modernc.org/sqlite"
)

var DB *sqlx.DB

func InitDb() {
    var err error
    DB, err = sqlx.Open("sqlite", "SMS.sqlite")
    if err != nil {
        log.Fatal(err)
    }

    // Test the connection
    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // Create the classes table
    _, err = DB.Exec(`CREATE TABLE IF NOT EXISTS classes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        year INTEGER NOT NULL,
        tablename TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`)
    if err != nil {
        log.Fatal(err)
    }
}