package functions

import (
	"encoding/json"
	"log"
	"time"
)

type timeInfo struct {
	DayOfWeek  string
	DayOfMonth int
	Month      string
	Year       int
	Hour       int
	Minute     int
	Second     int
}

func JSON_Time() string {
	currentTime := time.Now()

	info := timeInfo{
		DayOfWeek:  currentTime.Weekday().String(),
		DayOfMonth: currentTime.Day(),
		Month:      currentTime.Month().String(),
		Year:       currentTime.Year(),
		Hour:       currentTime.Hour(),
		Minute:     currentTime.Minute(),
		Second:     currentTime.Second(),
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		log.Fatal("Error encodign JSON: ", err)
	}
	return string(jsonData)
}
