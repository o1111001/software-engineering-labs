package time

import (
	"encoding/json"
	"net/http"
	"software-engineering-labs/services/time"
)

type timeResponse struct {
	Time string `json:"time"`
}

func GetJson(w http.ResponseWriter, r *http.Request) {
	currentTime := time.GetTime()
	response := timeResponse{ Time: currentTime }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
