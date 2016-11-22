package fbm_bot_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const ACCESS_TOKEN = "EAAIOtyFmtbsBAA4CuLiZALf4R4voPZBg3AySB63XB8SdRsid7FB2dWwHJgAgONJ0olWEGcOEVYXEjsZBeQ1M124keNAWhgWj3XwIDJ4mfYCl1m1DUuwaZCaOZCm7BZCY6TWwAKTRL5Uv0BSilWVhwGZBDcVmUg8Cm5na19KrFUVOAZDZD"

func TestGenericTemplateMessage(t *testing.T) {
	//var i int
	//i = 3/(3-3)
	//t.Error(i)
	fmt.Println("Test 1")
	request := Request{
		//Recipient: RequestRecipient{Id: "1038405329567665"},
		Recipient: RequestRecipient{Id: "1315593905118354"}, // DebtsTracker.DEV
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
							NewRequestPostbackButton("Возврат", "return"),
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
