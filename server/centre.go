package server

import (
	"encoding/json"
	"net/http"

	"github.com/MxHonesty/change4backend/db"
	"github.com/MxHonesty/change4backend/logging"
)

type CentreHandler struct {
	Repo *db.Mongodb
}

func (c *CentreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logging.InfoLogger.Println("Request CentreHandle")
	enableCors(&w)
	centre := c.Repo.FindAllCentre()
	json.NewEncoder(w).Encode(centre)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
