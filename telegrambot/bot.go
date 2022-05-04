package telegrambot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func telegramGetMe() {
	resp, err := http.Get("https://api.telegram.org/bot5167317855:AAEWC1JzKxk7Wof8W51QcOgKB675vVRAVx4/getMe")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

type TelegramUpdateResult struct {
	Ok     bool
	Result []TelegramUpdate
}

type TelegramMessage struct {
	MessageId int `json:"message_id"`
	Text      string
	Chat      TelegramMessageChat
}

type TelegramMessageChat struct {
	Id int
}

type TelegramUpdate struct {
	UpdateId int             `json:"update_id"`
	Message  TelegramMessage // chat
	// message
}

func TelegramGetUpdates(offset int) TelegramUpdateResult {
	resp, err := http.Get("https://api.telegram.org/bot5167317855:AAEWC1JzKxk7Wof8W51QcOgKB675vVRAVx4/getUpdates?timeout=30&offset=" + strconv.Itoa(offset))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var updateResult TelegramUpdateResult
	json.Unmarshal(body, &updateResult)

	return updateResult
}

func TelegramSendMessage(chatId int, message string) string {
	// FIXME: Исправить проблему с амперсандами и вопросами
	resp, err := http.Get("https://api.telegram.org/bot5167317855:AAEWC1JzKxk7Wof8W51QcOgKB675vVRAVx4/SendMessage?text=" + message + "&chat_id=" + strconv.Itoa(chatId))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
