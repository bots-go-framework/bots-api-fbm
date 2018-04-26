package fbm_api

import (
	"encoding/json"
	//"github.com/strongo/strongo-bots"
	"testing"
	//"github.com/strongo/bots-framework/core"
)

func TestTextMessage(t *testing.T) {
	s := `{
  "object":"page",
  "entry":[
    {
      "id": 123456790,
      "time":1457764198246,
      "messaging":[
        {
          "sender":{
            "id":987654321
          },
          "recipient":{
            "id":123456790
          },
          "timestamp":1457764197627,
          "message":{
            "mid":"mid.1457764197618:41d102a3e1ae206a38",
            "seq":73,
            "text":"hello, world!"
          }
        }
      ]
    }
  ]
}`
	var receivedMessage ReceivedMessage
	err := json.Unmarshal([]byte(s), &receivedMessage)
	if err != nil {
		t.Error(err)
		return
	}
	if len(receivedMessage.Entries) != 1 {
		t.Errorf("Invalid length of receivedMessage.Entry: %v, expected 1", len(receivedMessage.Entries))
	}
}

//func TestEntryIsInterfaceOfWebhookEntry(t *testing.T) {
//	_ = bots.WebhookEntry(Entry{})
//}
//
//func TestEntriesIsInterfaceOfArrayOfWebhookEntry(t *testing.T) {
//	entries := []Entry{Entry{}}
//	webhookEntries := []bots.WebhookEntry{}
//	for _, entry := range entries {
//		_ = append(webhookEntries, bots.WebhookEntry(entry))
//	}
//}
//
//func TestSenderIsInterfaceOfWebhookSender(t *testing.T) {
//	_ = bots.WebhookSender(Sender{})
//}
//
//func TestSenderIsInterfaceOfWebhookRecipient(t *testing.T) {
//	_ = bots.WebhookRecipient(Recipient{})
//}
//
//func TestMessagingIsInterfaceOfWebhookInput(t *testing.T) {
//	//_ = bots.webhookInput(Messaging{})
//}
