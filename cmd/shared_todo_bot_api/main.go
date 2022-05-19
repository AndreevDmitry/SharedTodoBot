package main

import (
	"SharedTodoBot/repo"
	"SharedTodoBot/telegrambot"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("SharedTodoBot running...")
	lastOffset := 0

	botToken := os.Getenv("BOT_TOKEN")

	bot := telegrambot.NewWithTimeout(botToken, 30)

	go func() {
		for {
			updates := bot.GetUpdates(lastOffset)
			fmt.Println((updates))
			if len(updates.Result) == 0 {
				continue
			}

			result := updates.Result[0]
			lastOffset = result.UpdateId + 1
			chatId := strconv.Itoa(result.Message.Chat.Id)

			if result.Message.Text == "/delete_all" {
				repo.DeleteAll(chatId)
				continue
			}

			repo.Add(chatId, result.Message.Text)
			items := repo.GetAllByUserId(chatId)
			for i, item := range items {
				message := fmt.Sprintf("%s Todo %d: %s", time.Now().Format(time.ANSIC), i, item.Message)
				bot.SendMessage(chatId, message)
			}
		}
	}()

	fmt.Println("http://localhost:8080/todos")
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		items := repo.GetAll()
		result, _ := json.Marshal(items)
		w.Write(result)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
