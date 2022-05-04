package main

import (
	"SharedTodoBot/repo"
	"SharedTodoBot/telegrambot"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello world")
	lastOffset := 0

	go func() {
		for {
			updates := telegrambot.TelegramGetUpdates(lastOffset)
			fmt.Println((updates))
			if len(updates.Result) == 0 {
				continue
			}

			result := updates.Result[0]
			lastOffset = result.UpdateId + 1
			chatId := result.Message.Chat.Id

			repo.Add(chatId, result.Message.Text)
			items := repo.GetAllByUserId(chatId)
			for i, item := range items {
				message := fmt.Sprintf("%s Todo %d: %s", time.Now(), i, item.Message)
				telegrambot.TelegramSendMessage(chatId, message)
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
