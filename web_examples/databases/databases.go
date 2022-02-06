package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

const (
	username = "root"
	password = "root"
	hostname = "localhost:3306"
	dbname   = "mydb"
)

func dsn(dbName string) string {
	// Take note of the parseTime = true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
}

func openDB() {
	// Configure the database connection (always check errors)
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error opening the database: %s\n", err)
	}
	// defer db.Close()

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging the database: %s\n", err)
	}
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	_, err = db.Exec(query)
	if err != nil {
		log.Printf("Error executing the query : %s", err)
	}
}

func insertUser() {
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Printf("Error inserting data into table: %s\n", err)
	}

	userID, _ := result.LastInsertId()

	log.Printf("Inserted user %v successfully into USERS table\n", userID)
}

func querySingleUser() {
	var (
		q_id        int
		q_username  string
		q_password  string
		q_createdAt time.Time
	)

	// Query the database and scan the values into out variables. Don't forget to check for errors.
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err = db.QueryRow(query, 2).Scan(&q_id, &q_username, &q_password, &q_createdAt)
	if err != nil {
		log.Printf("Error querying: %s", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "id\tusername\tpassword\tcreated_at\t")
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n\n", q_id, q_username, q_password, q_createdAt)
	w.Flush()

}

func queryAllUsers() {
	type user struct {
		id        string
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
	if err != nil {
		log.Printf("Error querying: %s", err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
		if err != nil {
			log.Printf("Error populating user struct: %s", err)
		}
		users = append(users, u)
	}

	for _, singleUser := range users {
		log.Println(singleUser)
	}
}

func deleteUser() {
	_, err = db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
	if err != nil {
		log.Printf("Error deleting user: %s", err)
	}
}

func main() {
	openDB()
	defer db.Close()

	createTable()
	insertUser()
	querySingleUser()
	queryAllUsers()

	deleteUser()

	fmt.Println("--End of Main--")
}
