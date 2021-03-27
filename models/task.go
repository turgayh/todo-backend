package models

type Task struct {
	Id          int64  `json:"id"`
	Description string `json:"description" binding:"required"`
}
