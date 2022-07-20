package domain

import (
	"time"
)

type TodoItem struct {
	UserId   string
	Time     time.Time
	IsDone   bool
	IsActive bool
	Message  string
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

func (user *User) DoneTodo(number int) string {
	if number >= 0 && number < len(user.Todos) {
		user.Todos[number].IsDone = true
		return "Ok"
	}
	return "Out of range"
}

func (user *User) UndoneTodo(number int) string {
	if number >= 0 && number < len(user.Todos) {
		user.Todos[number].IsDone = false
		return "Ok"
	}
	return "Out of range"
}

func (user *User) DeleteTodo(number int) string {
	if number >= 0 && number < len(user.Todos) {
		user.Todos[number].IsActive = false
		return "Ok"
	}
	return "Out of range"
}

func (user *User) RestoreTodo(number int) string {
	if number >= 0 && number < len(user.Todos) {
		user.Todos[number].IsActive = true
		return "Ok"
	}
	return "Out of range"
}

func (user *User) AddTodo(message string) {
	currentTodo := TodoItem{
		UserId:   user.UserId,
		Time:     time.Now(),
		IsActive: true,
		Message:  message,
	}
	user.Todos = append(user.Todos, currentTodo)
}
