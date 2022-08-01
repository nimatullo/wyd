package database

import (
	"database/sql"

	"os"

	"log"

	_ "github.com/lib/pq"

	"wyd/activity"
)

var (
	DB *sql.DB
)

func InitDatabase() {
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Println("No database connection string found. Using default.")
		connStr = "postgres://puffer:puffer@localhost:5432/wyd"
	}

	// connStr += "?sslmode=disable"

	log.Println(DB)

	DB, _ = sql.Open("postgres", connStr)

	CreateTable()
}

func CreateSQLiteFile() {
	log.Println("Creating wyd.db...")
	file, err := os.Create("wyd.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println("wyd.db created.")
}

func CreateTable() {
	log.Println("Creating table...")
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS activity (name TEXT, website TEXT, since TEXT, ready BOOLEAN)")
	if err != nil {
		log.Fatal(err)
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func UpdateCurrentActivityInDb(activityUpdate activity.Activity) bool {
	stmt, err := DB.Prepare("UPDATE activity SET name=$1, website=$2, since=$3, ready=$4 WHERE name=$5")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()

	res, err := stmt.Exec(activityUpdate.Name, activityUpdate.Website, activityUpdate.Since, activityUpdate.Ready, activity.CURRENT_ACTIVITY.Name)
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Println(res.RowsAffected())
	return true
}

func GetCurrentActivityFromDb() activity.Activity {
	var CURRENT_ACTIVITY activity.Activity
	err := DB.QueryRow("SELECT * FROM activity").Scan(&CURRENT_ACTIVITY.Name, &CURRENT_ACTIVITY.Website, &CURRENT_ACTIVITY.Since, &CURRENT_ACTIVITY.Ready)
	if err != nil {
		log.Fatal(err)
	}
	return CURRENT_ACTIVITY
}
