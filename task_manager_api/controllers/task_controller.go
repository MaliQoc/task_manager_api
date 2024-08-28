package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi_base/models"
)

func GetHome(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"message": "Welcome to the Task Manager"},
	)
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"status": http.StatusOK, "tasks": models.Tasks},
	)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := models.FindTaskByID(id)

	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			gin.H{"status": http.StatusNotFound},
		)
		return
	}
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"status": http.StatusOK, "task": task},
	)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	if newTask.TaskID == "" || newTask.TaskName == "" {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest},
		)
	} else {
		n := len(models.Tasks)
		exists := false

		for i := 0; i < n; i++ {
			if newTask.TaskID == models.Tasks[i].TaskID {
				exists = true
				break
			}
		}

		if exists {
			c.IndentedJSON(
				http.StatusBadRequest,
				gin.H{"status": http.StatusBadRequest},
			)
		} else {
			models.Tasks = append(models.Tasks, newTask)
			c.IndentedJSON(
				http.StatusCreated,
				gin.H{"status": http.StatusCreated, "task": newTask},
			)
		}
	}
}

