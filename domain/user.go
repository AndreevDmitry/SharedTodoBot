package domain

type TodoItem struct {
	UserId  string
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

func (user *User) AddTodo(message string) {
	currentTodo := TodoItem{
		UserId:  user.UserId,
		Message: message,
	}
	user.Todos = append(user.Todos, currentTodo)
}
