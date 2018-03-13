package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Task struct {
	Title string    `json:"titler"`
	Done  bool      `json:"-"`
	Date  time.Time `json:"created_at"`
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	t := Task{
		Title: "Title of my task",
		Done:  false,
		Date:  time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(&t)
	if err != nil {
		panic(err)
	}
}
func main() {
	http.HandleFunc("/api/tasks", GetTask)
	http.ListenAndServe(":8000", nil)
}
