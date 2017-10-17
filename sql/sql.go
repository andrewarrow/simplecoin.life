package sql

import "database/sql"
import _ "github.com/mattn/go-sqlite3"
import "runtime"
import "time"

import "fmt"

//import "encoding/base64"
import "os"
import "github.com/andrewarrow/simplecoin.life/words"
import "github.com/andrewarrow/simplecoin.life/crypto"

var dbPath string

func SetDbPath(db string) {
	dbPath = db
}

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
	fmt.Println("dbPath", dbPath)
	database, _ := sql.Open("sqlite3", dbPath)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS transactions (id TEXT PRIMARY KEY, owner TEXT, signature TEXT, previous_id TEXT, transfered_at BIGINT, created_at BIGINT)")
	statement.Exec()
	statement.Close()
	return database
}

func TransferCoinFromGenesis(owner string, database *sql.DB) {
	w := words.BigWords()
	statement, _ := database.Prepare("INSERT INTO transactions (id, owner, created_at) VALUES (?, ?, ?)")
	statement.Exec(w, owner, time.Now().Unix())
	defer statement.Close()
}
func TransferCoin(new_owner, previous string, database *sql.DB) {
	w := words.BigWords()
	statement, _ := database.Prepare("INSERT INTO transactions (id, owner, previous_id, created_at) VALUES (?, ?, ?, ?)")
	statement.Exec(w, new_owner, previous, time.Now().Unix())
	statement.Close()
	statement, _ = database.Prepare("UPDATE transactions set transfered_at=? WHERE id=?")
	statement.Exec(time.Now().Unix(), previous)
	statement.Close()
}
func InsertTransactionFromPeer(id, owner, previous string, transfered, created int64, database *sql.DB) {
	statement, _ := database.Prepare("INSERT INTO transactions (id, owner, previous_id, transfered_at, created_at) VALUES (?, ?, ?, ?, ?)")
	statement.Exec(id, owner, previous, transfered, created)
	statement.Close()
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
	rows, _ := database.Query("SELECT count(owner) FROM transactions where owner=? and transfered_at is null", owner)
	var sid uint64
	for rows.Next() {
		rows.Scan(&sid)
		break
	}
	defer rows.Close()
	return sid
}
func FindAvailableCoin(owner string, database *sql.DB) string {
	rows, _ := database.Query("SELECT id FROM transactions where owner=? and transfered_at is null", owner)
	var sid string
	for rows.Next() {
		rows.Scan(&sid)
		break
	}
	defer rows.Close()
	return sid
}
func TransactionsFrom(database *sql.DB) crypto.TransactionList {
	tl := crypto.TransactionList{}
	tl.Items = make([]crypto.Transaction, 0)
	rows, _ := database.Query("SELECT created_at, COALESCE(transfered_at, 0), id, owner, COALESCE(previous_id, ' ') FROM transactions order by created_at desc limit 100")
	for rows.Next() {
		var id, owner, previous string
		var created, transfered int64

		rows.Scan(&created, &transfered, &id, &owner, &previous)

		t := crypto.NewTransaction(id, owner, previous, created, transfered)
		//list = append([]crypto.Transaction{t}, list...)
		tl.Items = append(tl.Items, t)
	}
	defer rows.Close()
	return tl
}