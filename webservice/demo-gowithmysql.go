package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func creatingTable(db *sql.DB) {
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func query(db *sql.DB) {

	var (
		id         int
		coursename string
		price      float64
		instructor string
	)
	for {
		var inputID int
		fmt.Scan(&inputID)
		query := "SELECT id, coursename, price, instructor FROM onlinecourse WHERE id = ?"
		if err := db.QueryRow(query, inputID).Scan(&id, &coursename, &price, &instructor); err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		} else {
			fmt.Println(id, coursename, price, instructor)
		}
	}

}

func Insert(db *sql.DB) { // Insert a new user
	// username := "johndoe"
	// password := "secret"
	var username string
	var password string
	fmt.Scan(&username)
	fmt.Scan(&password)
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
}

func delete(db *sql.DB) {
	var id int
	fmt.Scan(&id)
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "admin:m6tqUnuYfjWzVFtT@tcp(127.0.0.1:3306)/coursedb")
	if err != nil {
		fmt.Println("failed to connect")
		fmt.Println(err)
	} else {
		fmt.Println("connect successfully")
	}
	// creatingTable(db)
	// Insert(db)
	// delete(db)
}
