package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

func (e *Event) MarshalJSON() ([]byte, error) {
	type Alias Event
	return json.Marshal(&struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		*Alias
	}{
		StartTime: e.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:   e.EndTime.Format("2006-01-02 15:04:05"),
		Alias:     (*Alias)(e),
	})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Result string `json:"result"`
}

func main() {
	http.HandleFunc("/create_event", createEvent)
	http.HandleFunc("/update_event", updateEvent)
	http.HandleFunc("/delete_event", deleteEvent)
	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	// Реализация создания события
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	// Реализация обновления события
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	// Реализация удаления события
}

func eventsForDay(w http.ResponseWriter, r *http.Request) {
	// Реализация получения событий на день
}

func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	// Реализация получения событий на неделю
}

func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	// Реализация получения событий на месяц
}
