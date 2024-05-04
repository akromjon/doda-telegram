package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const baseURL string = "https://api.telegram.org/bot"

type Telegram struct {
	BotToken string
	ChatID   string
}

func NewTelegram(c *Config) Telegram {

	return Telegram{
		BotToken: c.BotToken,
		ChatID:   c.ChatID,
	}

}

func (t *Telegram) base(Command *string, body []byte) *http.Response {

	// Send the POST request to Telegram Bot API
	apiUrl := baseURL + *&t.BotToken + "/" + *Command

	res, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(body))

	if err != nil {
		DD(err)
	}

	if res.StatusCode != 200 {

		var buf bytes.Buffer

		if _, err := io.Copy(&buf, res.Body); err != nil {

			DD(err)

		}

		DD(buf.String())
	}

	return res
}
func (t *Telegram) SendFile(f *string) (bool, error) {

	// Open the document file
	file, err := os.Open(*f)

	if err != nil {

		DD(err)

	}

	defer file.Close()

	// Prepare the request body
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("document", *f)

	if err != nil {

		DD(err)

	}
	_, err = io.Copy(part, file)

	if err != nil {

		DD(err)

	}

	err = writer.Close()

	if err != nil {

		DD(err)

	}

	apiUrl := fmt.Sprintf("%s%s/sendDocument?chat_id=%s", baseURL, t.BotToken, t.ChatID)

	req, err := http.NewRequest("POST", apiUrl, body)

	if err != nil {

		DD(err)

	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		DD(err)

	}

	msg := "File was sent successfully!"

	Echo(msg)

	defer resp.Body.Close()

	return true, nil

}

func (t *Telegram) SendMessage(m *string) (bool, error) {

	requestBody, _ := json.Marshal(map[string]string{
		"chat_id": t.ChatID,
		"text":    *m,
	})

	msg := "sendMessage"

	res := t.base(&msg, requestBody)

	if res.StatusCode == 200 {

		Echo("Message was sent successfully!")

		return true, nil
	}

	return false, errors.New("cannot send a message")

}
