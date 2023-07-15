package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Student_info struct {
	id       int
	name     string
	favorite string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "student"
)

func QueryFromDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database Open Failed")
		log.Fatal(err)
		return
	}

	var student Student_info

	query, err := database.Query("select * from student_info")
	if err != nil {
		fmt.Println("Query Generating Error")
		return
	}

	for query.Next() {
		query.Scan(&student.id, &student.name, &student.favorite)

		fmt.Println(student)
	}

	query.Close()

	database.Close()
}

func InsertIntoDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database Open Failed")
		log.Fatal(err)
		return
	}

	tx, err := database.Begin()
	if err != nil {
		fmt.Println("Transaction Begin Failed")
		return
	}

	stmt, err := tx.Prepare("insert into student_info(sname, sfavorite) values ('Corax', 'Riders Republic')")
	if err != nil {
		fmt.Println("Statement Preparation Failed: ", err)
		return
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Statement Execution Failed: ", err)
		return
	}

	tx.Commit()

}

func main() {
	InsertIntoDatabase()

	QueryFromDatabase()
}
