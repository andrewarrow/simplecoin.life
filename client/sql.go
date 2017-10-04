package client

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "runtime"

//import "encoding/base64"
import "fmt"
import "os"

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func SqlInit() {
	database, _ := sql.Open("sqlite3", UserHomeDir()+"/.scl.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS transactions (id UNSIGNED BIG INT PRIMARY KEY, owner TEXT, signature TEXT, previous_id UNSIGNED BIG INT)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO transactions (id, owner, signature) VALUES (?, ?, ?)")
	statement.Exec(2, "test", "sig")
	rows, _ := database.Query("SELECT id, owner FROM transactions")
	var id int
	var owner string
	for rows.Next() {
		rows.Scan(&id, &owner)
		fmt.Println("hi", id, owner)
	}
}
