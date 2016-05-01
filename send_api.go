package fbm_bot_api

type RequestNotificationType string

const (
	RequestAttachmentTypeImage = "image"
	RequestAttachmentTypeTemplate = "template"
)
const (
	RequestNotificationTypeRegular = RequestNotificationType("REGULAR")
	RequestNotificationTypeSilentPush = RequestNotificationType("SILENT_PUSH")
	RequestNotificationTypeNoPush = RequestNotificationType("NO_PUSH")
)

const (
	RequestAttachmentTemplateTypeGeneric = "generic"
	RequestAttachmentTemplateTypeButton = "button"
	RequestAttachmentTemplateTypeReceipt = "receipt"
)

const (
	RequestButtonTypePostback = "postback"
	RequestButtonTypeWebUrl = "web_url"
)


type Request struct {
	Recipient RequestRecipient `json:"recipient"`
	Message RequestMessage `json:"message"`
	NotificationType RequestNotificationType `json:"notification_type,omitempty"` // Optional: Push notification type: REGULAR, SILENT_PUSH, NO_PUSH
}

type RequestRecipient struct { // phone_number or id must be set
	Id int64 `json:"id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"` // Phone number of the recipient with the format +1(212)555-2368
}

type RequestMessage struct {  				// text or attachment must be set
	Text string `json:"text,omitempty"`  				// Message text, must be UTF-8, 320 character limit
	Attachment *RequestAttachment `json:"attachment,omitempty"`  // Attachment data
}

type RequestAttachment struct {
	Type string `json:"type,omitempty"`
	Payload RequestAttachmentPayload `json:"payload,omitempty"`
}

type RequestAttachmentImage struct {
	Url string `json:"url,omitempty"`
}

type RequestAttachmentGenericTemplate struct {
	Elements []RequestElement `json:"elements,omitempty"` // Data for each bubble in message
}

type RequestAttachmentButtonTemplate struct {
}

type RequestAttachmentReceiptTemplate struct {
}

type RequestAttachmentTemplate struct {
	TemplateType string `json:"template_type,omitempty"`
	RequestAttachmentGenericTemplate
	RequestAttachmentButtonTemplate
	RequestAttachmentReceiptTemplate
}

type RequestAttachmentPayload struct {
	RequestAttachmentImage
	RequestAttachmentTemplate
}

type RequestWebUrlButton struct {
	Url string `json:"url,omitempty"` // For web_url buttons, this URL is opened in a mobile browser when the button is tapped
}

type RequestPostbackButton struct {
	Payload string `json:"payload,omitempty"` // For postback buttons, this data will be sent back to you via webhook
}

type RequestButton struct {
	Type string `json:"type"` // Value is web_url or postback
	Title string `json:"title"` // Button title
	RequestWebUrlButton
	RequestPostbackButton
}

func NewRequestPostbackButton(title, payload string) RequestButton {
	return RequestButton{
		Type: RequestButtonTypePostback,
		Title: title,
		RequestPostbackButton: RequestPostbackButton{Payload: payload},
	}
}

func NewRequestWebUrlButton(title, url string) RequestButton {
	return RequestButton{
		Type: RequestButtonTypeWebUrl,
		Title: title,
		RequestWebUrlButton: RequestWebUrlButton{Url: url},
	}
}

type RequestElement struct {
	Title string 				`json:"title"` 		// Required: Bubble title
	Subtitle string 			`json:"subtitle,omitempty"` 	// Optional: Bubble subtitle
	ImageUrl string 			`json:"image_url,omitempty"` 	// Optional: Bubble image
	ItemUrl string 			`json:"item_url,omitempty"` 	// Optional: URL that is opened when bubble is tapped
	Buttons []RequestButton `json:"buttons,omitempty"`		// Optional: Set of buttons that appear as call-to-actions
}

func NewGenericTemplate(elements ...RequestElement) RequestAttachmentPayload {
	return RequestAttachmentPayload{
		RequestAttachmentTemplate: RequestAttachmentTemplate{
			TemplateType: RequestAttachmentTemplateTypeGeneric,
			RequestAttachmentGenericTemplate: RequestAttachmentGenericTemplate{
				Elements: elements,
			},
		},
	}
}