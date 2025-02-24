package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// -------------------STRUCT----------------------------\\
type requestTask struct {
	Task string `json:"task"`
}

// -------------------GLOBAL-VARABLE---------------------\\
var task string

// -------------------HANDLER-GET-------------------------\\
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", task)
}

// -------------------HANDLER-POST-------------------------\\
func SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req requestTask
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	task = req.Task
	fmt.Fprintf(w, "Task updated to: %s", task)
}

// -------------------MAIN-FUNC----------------------------\\
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", SetTaskHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
