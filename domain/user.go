package domain

import "time"

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

func (user *User) Done(number int) {
	user.Todos[number].IsDone = true
}

func (user *User) AddTodo(message string) {
	currentTodo := TodoItem{
		UserId:  user.UserId,
		Time:    time.Now(),
		Message: message,
	}
	user.Todos = append(user.Todos, currentTodo)
}
