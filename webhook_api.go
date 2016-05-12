package fbm_bot_api

import (
	"github.com/strongo/bots-framework/core"
	"time"
)

// ReceivedMessage ...
type ReceivedMessage struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

// Entry ...
type Entry struct {
	ID        int64       `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

func (e Entry) GetID() interface{} {
	return e.ID
}

func (e Entry) GetTime() time.Time {
	return time.Unix(e.Time, 0)
}

// Messaging ...
type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   *Message  `json:"message"`
	Postback  *Postback `json:"postback"`
	Delivery  *Delivery `json:"delivery"`
}

func (m Messaging) GetSender() bots.WebhookSender {
	return m.Sender
}
func (m Messaging) GetRecipient() bots.WebhookRecipient {
	return m.Recipient
}
func (m Messaging) GetTime() time.Time {
	return time.Unix(m.Timestamp, 0)
}

func (m Messaging) InputMessage() bots.WebhookMessage {
	return nil
}
func (m Messaging) InputPostback() bots.WebhookPostback {
	return nil
}
func (m Messaging) InputDelivery() bots.WebhookDelivery {
	return nil
}

func (m Messaging) InputInlineQuery() bots.WebhookInlineQuery {
	panic("Not implemented")
}

func (m Messaging) InputType() bots.WebhookInputType {
	switch {
	case m.Message != nil:
		if len(m.Message.Attachments) > 0 {
			return bots.WebhookInputAttachment
		} else if len(m.Message.Text) > 0 {
			return bots.WebhookInputMessage
		}
	case m.Postback != nil:
		return bots.WebhookInputPostback
	case m.Delivery != nil:
		return bots.WebhookInputDelivery
	}
	return bots.WebhookInputUnknown
}

type Actor struct {
	ID int64 `json:"id"`
}

func (a Actor) GetID() interface{} {
	return a.ID
}

func (a Actor) GetFirstName() string {
	return ""
}

func (a Actor) GetLastName() string {
	return ""
}

func (a Actor) GetUserName() string {
	return ""
}

// Sender ...
type Sender struct {
	Actor
}

// Recipient ...
type Recipient struct {
	Actor
}

type Postback struct {
	Payload Payload `json:"payload"`
}

type Delivery struct {
	Watermark  int64      `json:"watermark"`
	MessageIDs MessageIDs `json:"mids"`
}

type MessageID string

type MessageIDs []MessageID

// Message ...
type Message struct {
	MID         string       `json:"mid"`
	Seq         int64        `json:"seq"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"paylaod"`
}

type Payload struct {
	Url string `json:"url"`
}

// SendMessage ...
type SendMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   struct {
		Text string `json:"text"`
	} `json:"message"`
}
