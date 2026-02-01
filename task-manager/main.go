package main

import (
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

type TaskStore struct {
	///here is a comment
	mu     sync.RWMutex
	tasks  []Task
	nextID int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  []Task{},
		nextID: 1,
	}
}

	s.mu.Lock()
	defer s.mu.Unlock()

task := Task{
	id:        s.nextid,
	title:     title,
	Completed: false,
}

	s.tasks = append(s.tasks, task)
	s.nextID++
	return task
}
func (s *TaskStore) All() []Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Task, len(s.tasks))
	copy(result, s.tasks)
	return result
}


func (s *TaskStore) Toggle(id int) *Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Completed = !s.tasks[i].Completed
			return &s.tasks[i]
		}
	}
	return nil
}

func (s *TaskStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}

var templates *template.Template
var store = NewTaskStore()

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/tasks", handleAddTask)
	http.HandleFunc("/tasks/toggle/", handleToggleTask)
	http.HandleFunc("/tasks/delete/", handleDeleteTask)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Task Manager server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Tasks": store.All(),
	}
	templates.ExecuteTemplate(w, "index.html", data)
}

func handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Title required", http.StatusBadRequest)
		return
	}

	task := store.Add(title)
	templates.ExecuteTemplate(w, "task.html", task)
}

func handleToggleTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasks/toggle/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	task := store.Toggle(id)
	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	templates.ExecuteTemplate(w, "task.html", task)
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasks/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if !store.Delete(id) {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
