package client

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "runtime"

//import "encoding/base64"
import "os"
import "github.com/andrewarrow/simplecoin.life/words"

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

func SqlInit() *sql.DB {
	database, _ := sql.Open("sqlite3", UserHomeDir()+"/.scl.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS transactions (id TEXT PRIMARY KEY, owner TEXT, signature TEXT, previous_id TEXT)")
	statement.Exec()
	defer statement.Close()
	return database
}

func AddRow(owner string, database *sql.DB) {
	w := words.BigWords()
	statement, _ := database.Prepare("INSERT INTO transactions (id, owner) VALUES (?, ?)")
	statement.Exec(w, owner)
	defer statement.Close()
}
func TransferCoin(id, new_owner string, database *sql.DB) {
	w := words.BigWords()
	statement, _ := database.Prepare("INSERT INTO transactions (id, owner, previous_id) VALUES (?, ?, ?)")
	statement.Exec(w, new_owner, id)
	defer statement.Close()
}

func SelectId(id string, database *sql.DB) string {
	rows, _ := database.Query("SELECT id FROM transactions where id=?", id)
	var sid string
	for rows.Next() {
		rows.Scan(&sid)
		break
	}
	defer rows.Close()
	return sid
}
func CountByOwner(owner string, database *sql.DB) uint64 {
	rows, _ := database.Query("SELECT count(owner) FROM transactions where owner=?", owner)
	var sid uint64
	for rows.Next() {
		rows.Scan(&sid)
		break
	}
	defer rows.Close()
	return sid
}
func FindAvailableCoin(owner string, database *sql.DB) string {
	rows, _ := database.Query("SELECT id FROM transactions where owner=?", owner)
	var sid string
	for rows.Next() {
		rows.Scan(&sid)
		break
	}
	defer rows.Close()
	return sid
}

//StructureMailJobPeaceGrowthThroatActivity
