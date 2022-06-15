package main

import (
	"SharedTodoBot/domain"
	"SharedTodoBot/repo"
	"SharedTodoBot/telegrambot"
	"fmt"
	"os"
	"strconv"
	"strings"
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

		if result.Message.Text == "/delete_all" {
			handleDeleteAll(chatId)
			continue
		}

		if result.Message.Text == "/list" {
			handleList(chatId, bot)
			continue
		}

		if strings.HasPrefix(result.Message.Text, "/done ") {
			if number, err := strconv.Atoi(result.Message.Text[len("/done "):len(result.Message.Text)]); err == nil {
				doneStatus := handleDone(chatId, bot, number, true)
				bot.SendMessage(chatId, doneStatus)
			} else {
				bot.SendMessage(chatId, "Please, pass the Todo number from /list")
			}
			continue
		}

		if strings.HasPrefix(result.Message.Text, "/undone ") {
			if number, err := strconv.Atoi(result.Message.Text[len("/undone "):len(result.Message.Text)]); err == nil {
				doneStatus := handleDone(chatId, bot, number, false)
				bot.SendMessage(chatId, doneStatus)
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

func handleList(chatId string, bot telegrambot.Bot) {
	user := repo.Get(chatId)
	for i, todo := range user.Todos {
		//message := fmt.Sprintf("%s Todo %d: %s", time.Now().Format(time.ANSIC), i, todo.Message)
		var message string
		if todo.IsDone {
			message = fmt.Sprintf("%s ✅ %d: %s", todo.Time.Format(time.Kitchen), i, todo.Message)
		} else {
			message = fmt.Sprintf("%s  %d: %s", todo.Time.Format(time.Kitchen), i, todo.Message)
		}
		bot.SendMessage(chatId, message)
	}
}

func handleDeleteAll(chatId string) {
	user := repo.Get(chatId)
	user.Todos = []domain.TodoItem{}
	repo.Save(chatId, user)
}

func handleDone(chatId string, bot telegrambot.Bot, number int, done bool) string {
	user := repo.Get(chatId)
	result := user.SetStatus(number, done)
	repo.Save(chatId, user)
	return result
}
