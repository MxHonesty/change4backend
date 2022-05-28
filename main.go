package main

import (
	"net/http"

	"github.com/MxHonesty/change4backend/db"
	"github.com/MxHonesty/change4backend/logging"
	"github.com/MxHonesty/change4backend/server"
)

func main() {
	logging.InitLoggers()
	logging.InfoLogger.Println("Server Start")

	dbConn := db.NewMongodb()
	defer dbConn.CloseConnection()

	http.Handle("/centre", &server.CentreHandler{Repo: dbConn})
	logging.InfoLogger.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
