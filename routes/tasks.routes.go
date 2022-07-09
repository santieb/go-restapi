package routes

import (
	"encoding/json"
	"net/http"
	"github.com/santieb/go-restapi/db"
	"github.com/santieb/go-restapi/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []*models.Task
	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(tasks)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	idTask := params["id"]

	db.DB.First(&task, idTask)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	idTask := params["id"]

	db.DB.First(&task, idTask)

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Unscoped().Delete(&task)

	w.WriteHeader(http.StatusNoContent)
}