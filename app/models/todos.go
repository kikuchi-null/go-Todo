package models

import (
	"time"
)

type Todo struct {
	Todo_ID   int
	User_ID   int
	Content   string
	Create_At interface{}
	Update_At interface{}
}

func (u *User) CreateTodo(content string) (err error) {
	DB := gormConnect()
	defer DB.Close()

	t := Todo{
		User_ID:   u.User_ID,
		Content:   content,
		Create_At: time.Now(),
		Update_At: time.Now(),
	}
	result := DB.Create(&t)
	return result.Error
}

func GetTodo(todo_id int) (todo Todo, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("todo_id = ?", todo_id).First(&todo)

	todo.Create_At = todo.Create_At.(time.Time).Format("2006/01/02")
	todo.Update_At = todo.Update_At.(time.Time).Format("2006/01/02 15:04:05")

	return todo, result.Error
}

func GetTodos() (todos []Todo, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Find(&todos)

	for i, v := range todos {
		todo := &v
		todo.Create_At = todo.Create_At.(time.Time).Format("2006/01/02")
		todo.Update_At = todo.Update_At.(time.Time).Format("2006/01/02 15:04:05")
		todos[i] = *todo
	}

	return todos, result.Error

}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("user_id = ?", u.User_ID).Find(&todos)

	for i, v := range todos {
		todo := &v
		todo.Create_At = todo.Create_At.(time.Time).Format("2006/01/02")
		todo.Update_At = todo.Update_At.(time.Time).Format("2006/01/02 15:04:05")
		todos[i] = *todo
	}

	return todos, result.Error
}

func (t *Todo) UpdateTodo(newTodo string) (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Model(&t).Where("todo_id = ?", t.Todo_ID).Update(
		map[string]interface{}{"content": newTodo, "update_at": time.Now()})

	return result.Error
}

func (t *Todo) DeleteTodo() (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("todo_id = ?", t.Todo_ID).Delete(&t)

	return result.Error
}
