package taskrepo

import (
	"todo-backend/models"
	util "todo-backend/utils"
)

func Create(task models.Task) (int64, error) {
	db := util.DBConn()
	defer db.Close()

	query := "INSERT INTO tasks (description) VALUES(?);"
	stmt, stmtErr := db.Prepare(query)
	util.PanicError(stmtErr)

	res, queryErr := stmt.Exec(task.Description)
	util.PanicError(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	util.PanicError(getLastInsertIdErr)

	return id, queryErr
}

func GetAll() ([]models.Task, error) {
	var task []models.Task
	db := util.DBConn()
	defer db.Close()

	query := "SELECT id, description, create_date FROM tasks"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		timeline := models.Task{}
		err = rows.Scan(&timeline.Id, &timeline.Description)
		if err != nil {
			panic(err)
		}
		task = append(task, timeline)
	}
	return task, nil
}
