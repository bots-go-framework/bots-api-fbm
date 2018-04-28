package fbmbotapi

//go:generate ffjson $GOFILE

// GetStartedMessage is get started message
type GetStartedMessage struct {
	GetStarted struct {
		Payload string `json:"payload"`
	} `json:"get_started"`
}
