package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type HasNil interface {
	IsNil() bool
}

type Task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

func (t *Task) isNil() bool {
	return t == nil
}

type allTasks []Task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "TASK 1",
		Content: "Some content",
	},
}

func createTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	var newTask Task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "Insert a valid taks")
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 2
	tasks = append(tasks, newTask)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func deleteTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
	} else {

		for i, t := range tasks {
			println("position:", i)
			if t.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}

		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "dATOS ELMINADOS CORRECTAMENTE DE LA TAREA %v", id)
		json.NewEncoder(w).Encode(tasks)
	}

}

func getTask(w http.ResponseWriter, r *http.Request) {
	//extrae las de la url
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	println("id:", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
	} else {
		var task Task
		for index, taskI := range tasks {
			println("position:", index)

			if taskI.ID == id {
				task = taskI
				println("nombre: ", taskI.Name)
				break
			}

		}

		if task.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Objecto no encontrado")
		} else {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
		}
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func updateTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	var updatedTask Task
	println("id:", id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return

	}

	json.Unmarshal(reqBody, &updatedTask)
	fmt.Println("reqBody: ", updatedTask)
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Name = updatedTask.Name
			tasks[i].Content = updatedTask.Content
		}
	}

	fmt.Println(tasks)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func findPersonaByID(tasks []Task, id int) *Task {
	for _, t := range tasks {
		if t.ID == id {
			return &t // Return a pointer to the persona
		}
	}
	return nil // Persona not found
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the home pages 2")
}

func mains() {
	fmt.Println("Hello Worlds")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTasks).Methods("Delete")
	router.HandleFunc("/tasks/{id}", updateTasks).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))

}
