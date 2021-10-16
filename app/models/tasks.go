package models

import (
	"time"
)

type Task struct {
	Task_ID   int
	User_ID   int
	Content   string
	Create_At interface{}
	Update_At interface{}
	// Deadline Interface{}
}

func (u *User) CreateTask(content string) (err error) {
	DB := gormConnect()
	defer DB.Close()

	t := Task{
		User_ID:   u.User_ID,
		Content:   content,
		Create_At: time.Now(),
		Update_At: time.Now(),
		// Deadline:  stringToTime(deadline),
	}

	result := DB.Create(&t)
	return result.Error
}

func GetTask(Task_id int) (Task Task, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("Task_id = ?", Task_id).First(&Task)

	Task.Create_At = Task.Create_At.(time.Time).Format("2006/01/02")
	Task.Update_At = Task.Update_At.(time.Time).Format("2006/01/02 15:04:05")

	return Task, result.Error
}

func GetTasks() (Tasks []Task, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Find(&Tasks)

	for i, v := range Tasks {
		Task := &v
		Task.Create_At = Task.Create_At.(time.Time).Format("2006/01/02")
		Task.Update_At = Task.Update_At.(time.Time).Format("2006/01/02 15:04:05")
		Tasks[i] = *Task
	}

	return Tasks, result.Error

}

func (u *User) GetTasksByUser() (Tasks []Task, err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("user_id = ?", u.User_ID).Find(&Tasks)

	for i, v := range Tasks {
		Task := &v
		Task.Create_At = Task.Create_At.(time.Time).Format("2006/01/02")
		Task.Update_At = Task.Update_At.(time.Time).Format("2006/01/02 15:04:05")
		Tasks[i] = *Task
	}

	return Tasks, result.Error
}

func (t *Task) UpdateTask(newTask string) (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Model(&t).Where("Task_id = ?", t.Task_ID).Update(
		map[string]interface{}{"content": newTask, "update_at": time.Now()})

	return result.Error
}

func (t *Task) DeleteTask() (err error) {
	DB := gormConnect()
	defer DB.Close()

	result := DB.Where("Task_id = ?", t.Task_ID).Delete(&t)

	return result.Error
}
