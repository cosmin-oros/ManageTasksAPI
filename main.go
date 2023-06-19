package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
	"strconv"
)

type Task struct {
	ID       	int     `json:"id"`
	Title    	string  `json:"title"`
	Description string  `json:"description"`
	Complete 	bool    `json:"complete"`
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

func getTasks(c* gin.Context) {
	// return tasks as JSON
	c.IndentedJSON(http.StatusOK, tasks)
}

func createTask(c* gin.Context) {
	var newTask Task

	// bind the JSON payload to a new task and then append it
	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
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

	tasks[taskID-1].Title = updatedTask.Title
	tasks[taskID-1].Description = updatedTask.Description
	tasks[taskID-1].Complete = updatedTask.Complete
	c.Status(http.StatusOK)
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID := parseID(id)
	if taskID == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	tasks = append(tasks[:taskID-1], tasks[taskID:]...)
	c.Status(http.StatusOK)
}

func parseID(id string) int {
	taskID, err := strconv.Atoi(id)
	if err != nil || taskID <= 0 || taskID > len(tasks) {
		return -1
	}
	return taskID
}

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.POST("/tasks", createTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run("localhost:8080")

}