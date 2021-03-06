package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/MxHonesty/change4backend/authentication"
	"github.com/MxHonesty/change4backend/db"
	"github.com/MxHonesty/change4backend/logging"
	"github.com/MxHonesty/change4backend/server"
)

func main() {
	logging.InitLoggers()
	logging.InfoLogger.Println("Server Start")

	dbConn := db.NewMongodb()
	SetupCloseHandler(dbConn)
	defer dbConn.CloseConnection()

	authentication.SetupGoGuardian()
	http.Handle("/centre", &server.CentreHandler{Repo: dbConn})
	http.HandleFunc("/token", authentication.Middleware(http.HandlerFunc(authentication.CreateToken)))
	logging.InfoLogger.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func CleanUp(dbConn *db.Mongodb) {
	dbConn.CloseConnection()
}

func SetupCloseHandler(dbConn *db.Mongodb) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logging.InfoLogger.Println("Ctrl+C pressed in Terminal")
		CleanUp(dbConn)
		os.Exit(0)
	}()
}
