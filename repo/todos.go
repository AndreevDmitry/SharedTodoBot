package repo

import (
	"SharedTodoBot/domain"
	"encoding/json"
	"github.com/AndreevDmitry/gobitcask"
	"os"
)

var homedir, _ = os.UserHomeDir()
var db = gobitcask.New(homedir + "/.SharedTodoBotDb")

//var todos []TodoItem

func Save(userId string, todos domain.User) {
	encoded, _ := json.Marshal(todos)
	db.Put(userId, string(encoded))
}

func Get(userId string) domain.User {
	rawUser, err := db.Get(userId)
	if err != nil && err.Error() == "Bitcask: record not found" {
		return domain.NewUser(userId)
	}
	if err != nil {
		panic(err)
	}

	var user domain.User
	err = json.Unmarshal([]byte(rawUser), &user)
	if err != nil {
		panic(err)
	}
	return user
}
