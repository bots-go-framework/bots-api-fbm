package fbm_bot_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const ACCESS_TOKEN = "CAAGdsU6rzXgBAIVVFCQKsmZCue0PUpzuZA4BaZA80UfhxnRH2Nbf5Ri9K66tkXwkLuPa2WhN53MsiAngUUcE2wZBuhb5ZBO0DV5hVAQbOFuCuL5rP35FFQuf2NCkSs0IVwmhpkXkeAAt3a4yn4ZCnBkfearPByE4gbvSD4WfswvZBb6GrtTJ2ZAEgvawDfUWKKdcm8yXsuz2ZBAZDZD"

func TestGenericTemplateMessage(t *testing.T) {
	//var i int
	//i = 3/(3-3)
	//t.Error(i)
	fmt.Println("Test 1")
	request := Request{
		Recipient: RequestRecipient{Id: 1038405329567665},
		Message: RequestMessage{
			Attachment: &RequestAttachment{
				Type: RequestAttachmentTypeTemplate,
				Payload: NewGenericTemplate(
					RequestElement{
						Title:    "Записать долг",
						Subtitle: "Здесь можно создать новую операцию.\n\nПролистайте карточки влево-впрво чтобы увидеть другие опции.",
						Buttons: []RequestButton{
							NewRequestPostbackButton("Взял", "take"),
							NewRequestPostbackButton("Дал", "give"),
							NewRequestPostbackButton("Вернул", "return"),
						},
					},
					RequestElement{
						Title:    "Как идут дела?",
						Subtitle: "Тут смотрим баланс, историю и ближайшие платежи.",
						Buttons: []RequestButton{
							NewRequestPostbackButton("Баланс", "give"),
							NewRequestPostbackButton("История", "history"),
							NewRequestPostbackButton("Платежи", "payments"),
						},
					},
					RequestElement{
						Title:    "Картинка",
						Subtitle: "Пробная карточка",
						ImageUrl: "https://debtstracker.io/img/debtstracker-512x512.png",
						Buttons: []RequestButton{
							NewRequestWebUrlButton("https://debtstracker.io/", "https://debtstracker.io/"),
						},
					},
				),
			},
		},
		NotificationType: RequestNotificationTypeNoPush,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
	t.Log(string(data))
	resp, err := http.Post(
		"https://graph.facebook.com/v2.6/me/messages?access_token="+ACCESS_TOKEN,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(respData))
	t.Log(string(respData))
}
