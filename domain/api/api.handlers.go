package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type TimeResponse struct {
	CurrentTime time.Time `json:"current_time"`
}

func ApiTime(w http.ResponseWriter, r *http.Request) {
	tz := r.URL.Query().Get("tz")
	w.Header().Add("Content-Type", "application/json")

	if tz == "" {
		timeNow := time.Now()
		json.NewEncoder(w).Encode(TimeResponse{timeNow})
		return
	}

	results := make(map[string]time.Time)
	timeZones := strings.Split(tz, ",")

	for _, location := range timeZones {
		loc, err := time.LoadLocation(location)

		if err != nil {
			log.Fatal(err)
		}

		results[location] = time.Now().In(loc)
	}

	json.NewEncoder(w).Encode(results)
}
