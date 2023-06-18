package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
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

func main() {
	
}