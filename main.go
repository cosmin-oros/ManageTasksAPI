package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

var tasks = []Task{
	{
		ID:          1,
		Title:       "Task 1",
		Description: "Description for Task 1",
		Complete:    false,
	},
	{
		ID:          2,
		Title:       "Task 2",
		Description: "Description for Task 2",
		Complete:    true,
	},
	{
		ID:          3,
		Title:       "Task 3",
		Description: "Description for Task 3",
		Complete:    false,
	},
	{
		ID:          4,
		Title:       "Task 4",
		Description: "Description for Task 4",
		Complete:    false,
	},
	{
		ID:          5,
		Title:       "Task 5",
		Description: "Description for Task 5",
		Complete:    true,
	},
	{
		ID:          6,
		Title:       "Task 6",
		Description: "Description for Task 6",
		Complete:    false,
	},
	{
		ID:          11,
		Title:       "Task 11",
		Description: "Description for Task 11",
		Complete:    false,
	},
	{
		ID:          12,
		Title:       "Task 12",
		Description: "Description for Task 12",
		Complete:    true,
	},
	{
		ID:          7,
		Title:       "Task 7",
		Description: "Description for Task 7",
		Complete:    false,
	},
	{
		ID:          8,
		Title:       "Task 8",
		Description: "Description for Task 8",
		Complete:    false,
	},
	{
		ID:          9,
		Title:       "Task 9",
		Description: "Description for Task 9",
		Complete:    true,
	},
	{
		ID:          10,
		Title:       "Task 10",
		Description: "Description for Task 10",
		Complete:    false,
	},
}

type MySQLConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

var db *sql.DB

func getTasks(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Complete)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
			return
		}
		tasks = append(tasks, task)
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO tasks (title, description, complete) VALUES (?, ?, ?)",
		newTask.Title, newTask.Description, newTask.Complete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	taskID, _ := result.LastInsertId()
	newTask.ID = int(taskID)

	c.IndentedJSON(http.StatusCreated, newTask)
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	taskID := parseID(id)
	if taskID == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE tasks SET title=?, description=?, complete=? WHERE id=?",
		updatedTask.Title, updatedTask.Description, updatedTask.Complete, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.Status(http.StatusOK)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID := parseID(id)
	if taskID == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	_, err := db.Exec("DELETE FROM tasks WHERE id=?", taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.Status(http.StatusOK)
}

func parseID(id string) int {
	taskID, err := strconv.Atoi(id)
	if err != nil || taskID <= 0 {
		return -1
	}
	return taskID
}

func loadMySQLConfig() (*MySQLConfig, error) {
	file, err := os.Open(".secrets")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config MySQLConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	router := gin.Default()

	config, err := loadMySQLConfig()
	if err != nil {
		log.Fatal("Failed to load MySQL config: ", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to MySQL: ", err)
	}
	defer db.Close()

	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run("localhost:8080")

}
