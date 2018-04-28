package fbmbotapi

//go:generate ffjson $GOFILE

import (
	"time"
)

// ReceivedMessage ...
type ReceivedMessage struct {
	Object  string  `json:"object"`
	Entries []Entry `json:"entry"`
}

// Entry hold entry data
type Entry struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

// GetID returns ID
func (e Entry) GetID() interface{} {
	return e.ID
}

// GetTime returns time of the message
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

// Actor represents actor
type Actor struct {
	ID string `json:"id"`
}

// GetID returns ID of the actor
func (a Actor) GetID() interface{} {
	return a.ID
}

// GetFirstName returns first name
func (a Actor) GetFirstName() string {
	return "First" //TODO: Make call to API
}

// GetLastName returns last name
func (a Actor) GetLastName() string {
	return "Last" //TODO: Make call to API
}

// GetUserName returns username
func (a Actor) GetUserName() string {
	return "Username" //TODO: Make call to API
}

// Platform returns "fbm"
func (a Actor) Platform() string {
	return "fbm"
}

// Sender represents sender
type Sender struct {
	Actor
}

// IsBotUser returns false
func (Actor) IsBotUser() bool {
	return false
}

// GetAvatar is not supported or not implemented yet
func (Actor) GetAvatar() string {
	return ""
}

// GetLanguage returns preferred language (not implemented yet)
func (Actor) GetLanguage() string {
	return "" // TODO: Check if we can return actual
}

// Recipient represents recipient
type Recipient struct {
	Actor
}

// Postback message
type Postback struct {
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

// Delivery message
type Delivery struct {
	Watermark  int64      `json:"watermark"`
	MessageIDs MessageIDs `json:"mids"`
}

// MessageID defines message ID
type MessageID string

// MessageIDs defiens message IDs
type MessageIDs []MessageID

// Message defines message
type Message struct {
	MID         string       `json:"mid"`
	Seq         int          `json:"seq"`
	MText       string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment defines attachment
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"paylaod"`
}

// Payload defines payload
type Payload struct {
	URL string `json:"url"`
}

// SendMessage ...
type SendMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   struct {
		Text string `json:"text"`
	} `json:"message"`
}

// IntID not supported
func (m Message) IntID() int64 {
	panic("Not supported")
}

// StringID returns message ID as string
func (m Message) StringID() string {
	return m.MID
}

// Sequence returns Seq of message
func (m Message) Sequence() int {
	return m.Seq
}

// Text returns text of the message
func (m Message) Text() string {
	return m.MText
}
