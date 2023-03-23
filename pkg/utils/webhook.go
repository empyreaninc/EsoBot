package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Thumbnail struct {
	URL string `json:"url"`
}
type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type Embed struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Color     int       `json:"color"`
	Timestamp string    `json:"timestamp"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Fields    []Fields  `json:"fields"`
	Footer    Footer    `json:"footer"`
}
type Webhook struct {
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
	Embeds    []Embed `json:"embeds"`
}

func CreateWebhook() Webhook {
	Wh := Webhook{
		Username:  "",
		AvatarURL: "",
		Embeds: []Embed{
			Embed{
				Title:     "",
				URL:       "",
				Color:     16411130,
				Thumbnail: Thumbnail{URL: ""},
				Fields:    []Fields{},
			},
		},
	}

	return Wh
}

func (wh Webhook) AddField(title string, value string, inline bool) {

	newField := Fields{
		Name:   title,
		Value:  value,
		Inline: inline,
	}

	wh.Embeds[0].Fields = append(wh.Embeds[0].Fields, newField)

}

func (wh Webhook) SendWebhook(url string) *http.Response {
	client := &http.Client{}

	webhookData, err := json.Marshal(wh)

	if err != nil {
		panic("Eror encoding webhook data")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(webhookData))

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		panic("Error creating webhook request")
	}

	webhookPost, err := client.Do(req)

	if err != nil {
		fmt.Println("Error posting webhook")
	}

	if webhookPost.StatusCode == 204 {
		return webhookPost
	} else {
		return webhookPost
	}
}

func SendSuccessHookEnd(email string, size string) {
	hook := CreateWebhook()
	hook.Username = "Esotera | END"
	hook.AddField("Successfully Entered END Raffle!", ".", false)
	hook.AvatarURL = "https://scontent.ffab2-1.fna.fbcdn.net/v/t1.6435-9/186480896_4372809109419762_5239562243289338853_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=LExUucWFp8cAX96NiMV&_nc_ht=scontent.ffab2-1.fna&oh=00_AT8jRz0hZprrPUFmwjO5sd2h2T3WCAm1xRQ0-RmDHZihVw&oe=636C2F0g"
	hook.AddField("Email", email, true)
	hook.AddField("Size", string(size), true)

	webhookReq := hook.SendWebhook("redacted") // use variable to check if post request was successful

	if webhookReq.StatusCode == 204 { //204 status is successful webhook post
		fmt.Println("Webhook sent")
	} else {
		fmt.Println("Webhook failed")
		fmt.Println(webhookReq.StatusCode)

	}
}
