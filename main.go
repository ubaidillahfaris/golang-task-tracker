package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "not started", "in progress", "done"
	CreatedAt   time.Time `json:"created_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

var (
	taskList TaskList
	mu       sync.Mutex
)

const filename = "tasks.json"

// Load tasks from JSON file
func loadTasks(filename string) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File tidak ada, kita akan membuatnya nanti
		}
		return err
	}

	return json.Unmarshal(file, &taskList)
}

// Save tasks to JSON file
func saveTasks(filename string) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// Add a task
func addTask(description string) {
	mu.Lock()
	defer mu.Unlock()

	newTask := Task{
		ID:          len(taskList.Tasks) + 1,
		Description: description,
		Status:      "not started",
		CreatedAt:   time.Now(),
	}
	taskList.Tasks = append(taskList.Tasks, newTask)
}

// Update a task
func updateTask(id int, description string, status string) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range taskList.Tasks {
		if task.ID == id {
			// Update deskripsi
			if description != "" {
				taskList.Tasks[i].Description = description
			}
			// Update status hanya jika ada
			if status != "" {
				taskList.Tasks[i].Status = status
			}
			break
		}
	}
}

// Delete a task
func deleteTask(id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks = append(taskList.Tasks[:i], taskList.Tasks[i+1:]...)
			break
		}
	}
}

// Mark a task as in progress
func markTaskInProgress(id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Status = "in progress"
			break
		}
	}
}

// Mark a task as done
func markTaskDone(id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range taskList.Tasks {
		if task.ID == id {
			taskList.Tasks[i].Status = "done"
			break
		}
	}
}

// List tasks based on status
func listTasks(status string) {
	mu.Lock()
	defer mu.Unlock()

	for _, task := range taskList.Tasks {
		if status == "" || task.Status == status {
			fmt.Printf("ID: %d, Description: %s, Status: %s, Created At: %s",
				task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339))
		}
	}
}

func main() {
	// Load existing tasks
	if err := loadTasks(filename); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: task-tracker <command> [options]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		// Menambahkan tugas
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker add <description>")
			return
		}
		addTask(os.Args[2])
		if err := saveTasks(filename); err != nil {
			fmt.Println("Error saving tasks:", err)
		}

	case "update":
		// Memperbarui tugas
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-tracker update <id> <description> [status]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		description := os.Args[3]
		var status string
		if len(os.Args) > 4 {
			status = os.Args[4]
		}
		updateTask(id, description, status)
		if err := saveTasks(filename); err != nil {
			fmt.Println("Error saving tasks:", err)
		}
		fmt.Println("Task updated successfully")

	case "delete":
		// Menghapus tugas
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		deleteTask(id)
		if err := saveTasks(filename); err != nil {
			fmt.Println("Error saving tasks:", err)
		}

	case "list":
		args := os.Args
		if len(args) < 3 {
			listTasks("")
			return
		} else {
			listTasks(args[2])
			return
		}

	case "mark-in-progress":
		// Menandai tugas sebagai "in progress"
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		markTaskInProgress(id)
		if err := saveTasks(filename); err != nil {
			fmt.Println("Error saving tasks:", err)
		}
		fmt.Println("Task marked as in progress")

	case "mark-done":
		// Menandai tugas sebagai "done"
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-tracker mark-done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		markTaskDone(id)
		if err := saveTasks(filename); err != nil {
			fmt.Println("Error saving tasks:", err)
		}
		fmt.Println("Task marked as done")

	default:
		fmt.Println("Unknown command:", command)
	}
}
