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

var tasks []Task

func main() {
	
}