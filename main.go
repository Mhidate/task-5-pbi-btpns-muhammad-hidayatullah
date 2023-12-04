package main

import (

	"task5/database"
	// "task5/router"
	

)


func main() {
	
	db, err := database.ConnectDB()
	if err != nil {
	
		panic(err)
	}
	defer db.Close()

	
	database.AutoMigrateDB(db)

	


}
