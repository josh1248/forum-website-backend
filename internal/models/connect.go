package models

import (
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

// package scope variable representing DB connection.
var db *sqlx.DB

func ConnectToDB() {
	//declare here to prevent overshadowing of db variable:
	//https://stackoverflow.com/questions/34195360/how-to-use-global-var-across-files-in-a-package
	var err error

	//switch between :memory: and a file directory for transient/permanent DBs.
	//sqlx.Connect combines sql.Open with sql.Ping
	db, err = sqlx.Connect("sqlite", "internal/db/forumdb")
	//not needed, since it will be closed upon interruption.
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	//uncomment to clear out inputted data and reset to dummy data.
	resetDB()

	log.Println("Database connection established.")
}
