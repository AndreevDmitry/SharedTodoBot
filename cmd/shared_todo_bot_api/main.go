package main

import (
	"SharedTodoBot/domain"
	"SharedTodoBot/repo"
	"SharedTodoBot/telegrambot"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func main() {
	fmt.Println("SharedTodoBot running...")
	lastOffset := 0

	botToken := os.Getenv("BOT_TOKEN")

	bot := telegrambot.NewWithTimeout(botToken, 30)

	for {
		updates := bot.GetUpdates(lastOffset)
		fmt.Println(updates)
		if len(updates.Result) == 0 {
			continue
		}

		result := updates.Result[0]
		lastOffset = result.UpdateId + 1
		chatId := strconv.Itoa(result.Message.Chat.Id)

		if strings.HasPrefix(result.Message.Text, "/delete ") {
			if number, err := strconv.Atoi(result.Message.Text[len("/delete "):len(result.Message.Text)]); err == nil {
				deleteStatus := handleDelete(chatId, bot, number)
				bot.SendMessage(chatId, deleteStatus)
			} else {
				bot.SendMessage(chatId, "Please, pass the Todo number from /list")
			}
			continue
		}

		if strings.HasPrefix(result.Message.Text, "/restore ") {
			if number, err := strconv.Atoi(result.Message.Text[len("/restore "):len(result.Message.Text)]); err == nil {
				restoreStatus := handleRestore(chatId, bot, number)
				bot.SendMessage(chatId, restoreStatus)
			} else {
				bot.SendMessage(chatId, "Please, pass the Todo number from /list")
			}
			continue
		}

		if result.Message.Text == "/delete_all" {
			handleDeleteAll(chatId)
			continue
		}

		if result.Message.Text == "/list" {
			handleList(chatId, bot, true)
			continue
		}

		if result.Message.Text == "/list_deleted" {
			handleList(chatId, bot, false)
			continue
		}

		if strings.HasPrefix(result.Message.Text, "/done ") {
			if number, err := strconv.Atoi(result.Message.Text[len("/done "):len(result.Message.Text)]); err == nil {
				doneStatus := handleDone(chatId, bot, number)
				bot.SendMessage(chatId, doneStatus)
			} else {
				bot.SendMessage(chatId, "Please, pass the Todo number from /list")
			}
			continue
		}

		if strings.HasPrefix(result.Message.Text, "/undone ") {
			if number, err := strconv.Atoi(result.Message.Text[len("/undone "):len(result.Message.Text)]); err == nil {
				undoneStatus := handleUndone(chatId, bot, number)
				bot.SendMessage(chatId, undoneStatus)
			} else {
				bot.SendMessage(chatId, "Please, pass the Todo number from /list")
			}
			continue
		}

		handleSave(chatId, result)
	}
}

func handleSave(chatId string, result telegrambot.TelegramUpdate) {
	user := repo.Get(chatId)
	user.AddTodo(result.Message.Text)
	repo.Save(chatId, user)
}

func handleList(chatId string, bot telegrambot.Bot, active bool) {
	type Message struct {
		Time       string
		Todo       domain.TodoItem
		TodoNumber int
	}
	text := `{{.Time}} {{.TodoNumber}}: {{.Todo.Message}} {{if .Todo.IsDone}} ✅ {{end}}
`
	t, err := template.New("Todos").Parse(text)
	if err != nil {
		panic(err)
	}

	user := repo.Get(chatId)
	var output bytes.Buffer
	var message Message
	for i, todo := range user.Todos {
		if active != todo.IsActive {
			continue
		}
		message.Todo = todo
		message.Time = todo.Time.Format(time.Kitchen)
		message.TodoNumber = i + 1
		t.Execute(&output, message)
	}
	bot.SendMessage(chatId, output.String())
}

func handleDelete(chatId string, bot telegrambot.Bot, number int) string {
	user := repo.Get(chatId)
	result := user.DeleteTodo(number - 1)
	repo.Save(chatId, user)
	return result
}

func handleRestore(chatId string, bot telegrambot.Bot, number int) string {
	user := repo.Get(chatId)
	result := user.RestoreTodo(number - 1)
	repo.Save(chatId, user)
	return result
}

func handleDeleteAll(chatId string) {
	user := repo.Get(chatId)
	user.Todos = []domain.TodoItem{}
	repo.Save(chatId, user)
}

func handleDone(chatId string, bot telegrambot.Bot, number int) string {
	user := repo.Get(chatId)
	result := user.DoneTodo(number - 1)
	repo.Save(chatId, user)
	return result
}

func handleUndone(chatId string, bot telegrambot.Bot, number int) string {
	user := repo.Get(chatId)
	result := user.UndoneTodo(number - 1)
	repo.Save(chatId, user)
	return result
}
