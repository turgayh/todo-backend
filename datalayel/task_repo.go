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
	var tasks []models.Task
	db := util.DBConn()
	defer db.Close()

	query := "SELECT id, description, create_date FROM tasks"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		task := models.Task{}

		err = rows.Scan(&task.Id, &task.Description, &task.Create_Date)
		if err != nil {
			panic(err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
