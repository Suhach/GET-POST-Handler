package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// -------------------STRUCT----------------------------\\
type requestTask struct {
	Task string `json"task"`
}

// -------------------GLOBAL-VARABLE---------------------\\
var task string

// -------------------HANDLER-GET-------------------------\\
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello %s", task)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод 'GET'")
	}
}

// -------------------HANDLER-POST-------------------------\\
func SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req requestTask
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		task = req.Task
		fmt.Fprintf(w, "Task updated to %s", task)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод 'POST'")
	}
}

// -------------------MAIN-FUNC----------------------------\\
func main() {
	// router := mux.NewRouter()

	// router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	// router.HandleFunc("/api/task", SetTaskHandler).Methods("POST")

	//http.ListenAndServe(":8080", router)
	http.HandleFunc("/api/hello", HelloHandler)
	http.HandleFunc("/api/task", SetTaskHandler)
	http.ListenAndServe(":8080", nil)
}
