package main

import (
	"github.com/MxHonesty/change4backend/db"
	"github.com/MxHonesty/change4backend/logging"
)

func main() {
	logging.InitLoggers()
	logging.InfoLogger.Println("Server Start")

	dbConn := db.NewMongodb()
	defer dbConn.CloseConnection()
	// id, _ := dbConn.AddUser("root", "root", 1)
	// logging.InfoLogger.Println(id)
}
