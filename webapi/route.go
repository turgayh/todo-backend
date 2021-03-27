package webapi

import (
	taskhandle "todo-backend/webhandler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1TodoApp := r.Group("/v1/todo")
	{
		v1TodoApp.POST("/create", taskhandle.CreateHandler)
		v1TodoApp.GET("/get-all", taskhandle.ShowHandler)
		//v1TodoApp.PUT("/:id", taskhandle.PutHandler)
		//v1TodoApp.DELETE("/:id", taskhandle.DeleteHandler)
	}

	return r
}
