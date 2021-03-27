package taskhandle

import (
	"fmt"
	"net/http"
	taskrepo "todo-backend/data-layel"
	"todo-backend/models"
	util "todo-backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
	var task models.Task

	// App level validation
	bindErr := c.BindJSON(&task)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Inserting data
	id, insertErr := taskrepo.Create(task)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		util.PanicError(insertErr)
	} else {
		task.Id = id
		c.JSON(http.StatusCreated, task)
	}
}

func ShowHandler(c *gin.Context) {
	tasks, _ := taskrepo.GetAll()

	// Check if resource exist
	if tasks == nil {
		c.JSON(http.StatusNotFound, "Not found")
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}
