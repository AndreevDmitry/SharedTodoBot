package domain

import (
	"time"
)

type TodoItem struct {
	UserId  string
	Time    time.Time
	IsDone  bool
	Message string
}

type User struct {
	UserId string
	Todos  []TodoItem
}

func NewUser(userId string) User {
	user := User{
		UserId: userId,
	}
	return user
}

func (user *User) SetStatus(number int, status bool) string {
	if number < len(user.Todos) {
		user.Todos[number].IsDone = status
		return "Ok"
	}
	return "Out of range"
}

func (user *User) AddTodo(message string) {
	currentTodo := TodoItem{
		UserId:  user.UserId,
		Time:    time.Now(),
		Message: message,
	}
	user.Todos = append(user.Todos, currentTodo)
}
