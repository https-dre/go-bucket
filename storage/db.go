package storage

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}

	return db
}

func RunSeed(db *sql.DB) {
	query, err := os.ReadFile("../seed.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		panic(err)
	}
}

/*
	query := `SELECT id, name, age FROM users`
    rows, err := db.Query(query)
    if err != nil {
        panic(err)
    }

	// Insere um novo registro
    insertUser := `INSERT INTO users (name, age) VALUES (?, ?)`
    _, err = db.Exec(insertUser, "Alice", 25)
    if err != nil {
        panic(err)
    }

	// Imprime os dados
    for rows.Next() {
        var id int
        var name string
        var age int
        err = rows.Scan(&id, &name, &age)
        if err != nil {
            panic(err)
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }
*/
