package database

import (
	"database/sql"

	"os"

	"log"

	_ "github.com/lib/pq"

	"wyd/activity"
)

// Global database connection
var (
	DB *sql.DB
)

func InitDatabase() {
	connStr := os.Getenv("DATABASE_URL")

	// If server is running locally, use localhost for database connection.
	if connStr == "" {
		log.Println("No database connection string found. Using default.")
		connStr = "postgres://puffer:puffer@localhost:5432/wyd?sslmode=disable" // Avoids the "SSL is not enabled on the server." error
	}

	DB, _ = sql.Open("postgres", connStr)
	log.Println(DB)
	
	CreateTableIfNotExists()
}

func CreateTableIfNotExists() {
	log.Println("Creating table...")
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS activity (name TEXT, website TEXT, since TEXT, ready BOOLEAN)")
	if err != nil {
		log.Fatal(err)
	}

	CheckIfInitalDataPresent()
}

/* 
When the app first starts, like...literally the first time ever, there is no data in the database. This function ensures that 
there is always some default value to read when the stream endpoint is accessed.
*/
func CheckIfInitalDataPresent() {
	currentActivity := activity.Activity{}
	err := DB.QueryRow("SELECT * FROM activity").Scan(&currentActivity.Name, &currentActivity.Website, &currentActivity.Since, &currentActivity.Ready)
	if err != nil {
		log.Println("No inital data found. Inserting default data.")

		_, err := DB.Exec("INSERT INTO activity (name, website, since, ready) VALUES ($1, $2, $3, $4)", "the unknown", "https://nimatullo.com", "08-20-1999", true)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func UpdateCurrentActivityInDb(activityUpdate activity.Activity) bool { // Returns a boolean to let the caller know if the update was successful.
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
