package telegrambot

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"fmt"
)

type Bot struct {
	Token   string
	Timeout int
}

func New(token string) Bot {
	return NewWithTimeout(token, 30)
}

func NewWithTimeout(token string, timeout int) Bot {
	return Bot{
		Token:   token,
		Timeout: timeout,
	}
}

func (bot Bot) SendMessage(chatId string, message string) {
	bot.get("/SendMessage", []string{"text", "chat_id"}, []string{message, chatId})
}

type TelegramMessageChat struct {
	Id int
}

type TelegramMessage struct {
	MessageId int `json:"message_id"`
	Text      string
	Chat      TelegramMessageChat
}

type TelegramUpdate struct {
	UpdateId int             `json:"update_id"`
	Message  TelegramMessage // chat
	// message
}

type TelegramUpdateResult struct {
	Ok     bool
	Result []TelegramUpdate
	Description string  
}

func (bot Bot) GetUpdates(offset int) TelegramUpdateResult {
	body := bot.get("/getUpdates", []string{"timeout", "offset"}, []string{strconv.Itoa(bot.Timeout), strconv.Itoa(offset)})
	var updateResult TelegramUpdateResult
	json.Unmarshal(body, &updateResult)
	if !updateResult.Ok {
		fmt.Println("Telegram getUpdates result:", updateResult.Description)
	}
	
	return updateResult
}

func (bot Bot) get(method string, keys []string, values []string) []byte {
	if len(keys) != len(values) {
		panic("Then number of keys is not equal with the number of values")
	}

	req, err := http.NewRequest("GET", "https://api.telegram.org/bot"+bot.Token+method, nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	for i := range keys {
		q.Add(keys[i], values[i])
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}
