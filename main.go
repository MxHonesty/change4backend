package main

import (
	"fmt"

	"github.com/MxHonesty/change4backend/db"
	"github.com/MxHonesty/change4backend/logging"
)

func main() {
	logging.InitLoggers()
	logging.InfoLogger.Println("Server Start")

	dbConn := db.NewMongodb()
	defer dbConn.CloseConnection()
	fmt.Println(dbConn.FindAllUsers())
}
