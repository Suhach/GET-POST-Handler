package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// =----------GET--Handler---------------------------\\
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, "Ошибка при получении задач", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

// -------POST---Handler-----------------------------------\\
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req Task
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	DB.Create(&req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

// ---------main----func------------------------------------\\
func main() {

	initDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")

	http.ListenAndServe(":8080", router)

}
