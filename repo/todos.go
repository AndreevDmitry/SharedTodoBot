package repo

import (
	"github.com/AndreevDmitry/gobitcask"
	"os"
)

type TodoItem struct {
	UserId  string
	Message string
}

var homedir, _ = os.UserHomeDir()
var db = gobitcask.New(homedir + "/.SharedTodoBotDb")

//var todos []TodoItem

func Add(userId string, message string) {
	oldmessage, _ := db.Get(userId)
	db.Put(userId, oldmessage+message)
}

func GetAllByUserId(userId string) []TodoItem {
	var result []TodoItem

	message, _ := db.Get(userId)
	item := TodoItem{
		UserId:  userId,
		Message: message,
	}
	result = append(result, item)
	return result
}

func GetAll() []TodoItem {
	return nil
}
