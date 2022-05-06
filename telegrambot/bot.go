package telegrambot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	botToken := os.Getenv("BOT_TOKEN")

	client := &http.Client{}
	resp, err := client.Get("https://api.telegram.org")
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", "https://api.telegram.org/bot"+botToken+"/SendMessage", nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	q.Add("text", message)
	q.Add("chat_id", strconv.Itoa(chatId))
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	resp, err = client.Do(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
