package fbm_api

//go:generate ffjson $GOFILE

type GetStartedMessage struct {
	GetStarted struct {
		Payload string `json:"payload"`
	} `json:"get_started"`
}
