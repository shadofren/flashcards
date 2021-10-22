package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type Row struct {
	Id       int
	Question string
	Answer   string
	Correct  int
	Total    int
}

// Create database if doens't exist
func create(db string) {
	log.Println("Creating", db)
	file, err := os.Create(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println(db, "created")

	DB, _ := sql.Open("sqlite3", db)
	defer DB.Close()
	createTable(DB)
}

func Connect(db string) *sql.DB {
	if _, err := os.Stat(db); os.IsNotExist(err) {
		create(db)
	}
	conn, _ := sql.Open("sqlite3", db)
	return conn
}

func Remove(db string) {
	os.Remove(db)
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE data (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"question" TEXT,
		"answer" TEXT,
		"correct" integer,
		"total" integer
	  );`

	log.Println("Create data table...")
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("data table created")
}

// We are passing db reference connection from main to our method with other parameters
func Insert(db *sql.DB, question, answer string) {
	log.Printf("Inserting data record question=%s, answer=%s", question, answer)
	query := `INSERT INTO data(question, answer, correct, total) VALUES (?, ?, 1, 1)`
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(question, answer)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Get(db *sql.DB, question string) *Row {
	row := db.QueryRow(`SELECT * FROM data WHERE question = ?`, question)
	var r Row
	if err := row.Scan(&r.Id, &r.Question, &r.Answer, &r.Correct, &r.Total); err == nil {
		return &r
	}
	return nil
}

func Update(db *sql.DB, row *Row) {
	log.Println("Updating data record ", row)
	query := `UPDATE data set question=?, answer=?, correct=?, total=? where id = ?`
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(row.Question, row.Answer, row.Correct, row.Total, row.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func ListTopRows(db *sql.DB, n int) []*Row {
	rows, err := db.Query(`SELECT d.*, d.correct/d.total as familarity FROM data d ORDER BY familarity LIMIT ?`, n)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	results := make([]*Row, 0)
	for rows.Next() { // Iterate and fetch the records from result cursor
		var r Row
		var familarity float32
		rows.Scan(&r.Id, &r.Question, &r.Answer, &r.Correct, &r.Total, &familarity)
		results = append(results, &r)
	}
	return results
}
