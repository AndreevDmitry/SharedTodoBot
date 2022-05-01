package repo

type TodoItem struct {
	UserId  int
	Message string
}

var todos []TodoItem

func Add(userId int, message string) {
	item := TodoItem{
		UserId:  userId,
		Message: message,
	}
	todos = append(todos, item)

}

func GetAllByUserId(userId int) []TodoItem {
	var result []TodoItem
	for _, item := range todos {
		if userId == item.UserId {
			result = append(result, item)
		}
	}
	return result
}

func GetAll() []TodoItem {
	return todos
}
