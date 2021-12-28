package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var (
	host     = goDotEnvVariable("host")
	port     = goDotEnvVariable("port")
	user     = goDotEnvVariable("user")
	password = goDotEnvVariable("password")
	dbname   = goDotEnvVariable("dbname")
	driver   = goDotEnvVariable("driver")
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open(driver, psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// QUERY
	rows, err := db.Query(`SELECT * from salaries where salary > 50000 limit 10`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var emp_no int
		var salary string
		var from_date string
		var to_date string

		err = rows.Scan(&salary, &emp_no, &from_date, &to_date)
		CheckError(err)
		layout := "2006-01-02T15:04:05Z0700"
		updatedAt, _ := time.Parse(layout, from_date)
		fmt.Println(salary, emp_no, updatedAt, to_date)
	}
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
