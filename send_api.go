package fbm_api

//go:generate ffjson $GOFILE

import "fmt"

type RequestNotificationType string

const (
	RequestAttachmentTypeImage    = "image"
	RequestAttachmentTypeTemplate = "template"
)
const (
	RequestNotificationTypeRegular    = RequestNotificationType("REGULAR")
	RequestNotificationTypeSilentPush = RequestNotificationType("SILENT_PUSH")
	RequestNotificationTypeNoPush     = RequestNotificationType("NO_PUSH")
)

const (
	RequestAttachmentTemplateTypeGeneric = "generic"
	RequestAttachmentTemplateTypeList    = "list"
	RequestAttachmentTemplateTypeButton  = "button"
	RequestAttachmentTemplateTypeReceipt = "receipt"
)

const (
	RequestButtonTypePostback = "postback"
	RequestButtonTypeWebUrl   = "web_url"
)

type Request struct {
	Recipient        RequestRecipient        `json:"recipient"`
	Message          RequestMessage          `json:"message"`
	NotificationType RequestNotificationType `json:"notification_type,omitempty"` // Optional: Push notification type: REGULAR, SILENT_PUSH, NO_PUSH
}

type RequestRecipient struct { // phone_number or id must be set
	Id          string `json:"id,omitempty"` // Alex: changed from int64 on 21st November
	PhoneNumber string `json:"phone_number,omitempty"` // Phone number of the recipient with the format +1(212)555-2368
}

type RequestMessage struct { // text or attachment must be set
	Text       string             `json:"text,omitempty"`       // Message text, must be UTF-8, 320 character limit
	Attachment *RequestAttachment `json:"attachment,omitempty"` // Attachment data
}

type RequestAttachment struct {
	Type    string                   `json:"type,omitempty"`
	Payload RequestAttachmentPayload `json:"payload,omitempty"`
}

type RequestAttachmentImage struct {
	Url string `json:"url,omitempty"`
}

type RequestAttachmentGenericTemplate struct {
	Elements []RequestElement `json:"elements,omitempty"` // Data for each bubble in message
}

type RequestAttachmentListTemplate struct {
	RequestAttachmentGenericTemplate
	TopElementStyle string `json:"top_element_style,omitempty"`
}

type TopElementStyle string

const (
	TopElementStyleCompact TopElementStyle = "compact"
	TopElementStyleLarge TopElementStyle = "large"
)

type RequestAttachmentButtonTemplate struct {
	Text string `json:"text,omitempty"`
	Buttons []RequestButton `json:"buttons,omitempty"`
}

type RequestAttachmentReceiptTemplate struct {
}

type RequestAttachmentTemplate struct {
	TemplateType string `json:"template_type,omitempty"`
	RequestAttachmentListTemplate // Used for both List & Generic templates
	// RequestAttachmentGenericTemplate - conflicts with RequestAttachmentListTemplate on json.Marshal - silently not marshalling Elements.
	RequestAttachmentButtonTemplate
	RequestAttachmentReceiptTemplate
}

type RequestAttachmentPayload struct {
	RequestAttachmentImage
	RequestAttachmentTemplate
}

type RequestWebUrlAction struct {
	Url                 string `json:"url,omitempty"` // For web_url buttons, this URL is opened in a mobile browser when the button is tapped
	FallbackUrl         string `json:"fallback_url,omitempty"`
	WebviewHeightRatio  WebviewHeightRatio `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool `json:"messenger_extensions,omitempty"`
}

type RequestPostbackAction struct {
	Payload string `json:"payload,omitempty"` // For postback buttons, this data will be sent back to you via webhook
}

type RequestButton struct {
	Type  string `json:"type"`  // Value is web_url or postback
	Title string `json:"title"` // Button title
	RequestWebUrlAction
	RequestPostbackAction
}

func NewRequestPostbackButton(title, payload string) RequestButton {
	return RequestButton{
		Type:  RequestButtonTypePostback,
		Title: title,
		RequestPostbackAction: RequestPostbackAction{Payload: payload},
	}
}

func NewRequestWebUrlButton(title, url string) RequestButton {
	return RequestButton{
		Type:                RequestButtonTypeWebUrl,
		Title:               title,
		RequestWebUrlAction: RequestWebUrlAction{Url: url},
	}
}

func NewRequestWebUrlButtonWithRatio(title, url string, webviewHeightRatio WebviewHeightRatio) RequestButton {
	return RequestButton{
		Type:                RequestButtonTypeWebUrl,
		Title:               title,
		RequestWebUrlAction: RequestWebUrlAction{
			Url: url,
			WebviewHeightRatio: webviewHeightRatio,
		},
	}
}

func NewRequestWebExtentionUrlButtonWithRatio(title, url string, webviewHeightRatio WebviewHeightRatio) RequestButton {
	return RequestButton{
		Type:                RequestButtonTypeWebUrl,
		Title:               title,
		RequestWebUrlAction: RequestWebUrlAction{
			Url: url,
			MessengerExtensions: true,
			WebviewHeightRatio: webviewHeightRatio,
		},
	}
}

type RequestDefaultAction struct {
	Type  string `json:"type"`  // Value is web_url or postback
	RequestWebUrlAction
	RequestPostbackAction
}

func NewDefaultActionWithWebUrl(action RequestWebUrlAction) RequestDefaultAction {
	return RequestDefaultAction{
		Type: "web_url",
		RequestWebUrlAction: action,
	}
}

type RequestElement struct {
	Title         string          `json:"title"`               // Required: Bubble title
	Subtitle      string          `json:"subtitle,omitempty"`  // Optional: Bubble subtitle
	ImageUrl      string          `json:"image_url,omitempty"` // Optional: Bubble image
	ItemUrl       string          `json:"item_url,omitempty"`  // Optional: URL that is opened when bubble is tapped
	DefaultAction *RequestDefaultAction `json:"default_action,omitempty"`
	Buttons       []RequestButton `json:"buttons,omitempty"`   // Optional: Set of buttons that appear as call-to-actions
}

func NewRequestElementWithDefaultAction(title, subtitle string, defaultAction RequestDefaultAction, buttons ...RequestButton) RequestElement {
	return RequestElement{
		Title: title,
		Subtitle: subtitle,
		DefaultAction: &defaultAction,
		Buttons: buttons,
	}
}

func NewRequestElement(title, subtitle string, buttons ...RequestButton) RequestElement {
	return RequestElement{
		Title: title,
		Subtitle: subtitle,
		Buttons: buttons,
	}
}

func NewGenericTemplate(elements ...RequestElement) RequestAttachmentPayload {
	return RequestAttachmentPayload{
		RequestAttachmentTemplate: RequestAttachmentTemplate{
			TemplateType: RequestAttachmentTemplateTypeGeneric,
			RequestAttachmentListTemplate: RequestAttachmentListTemplate{
				RequestAttachmentGenericTemplate: RequestAttachmentGenericTemplate{Elements: elements},
			},
		},
	}
}

func NewButtonTemplate(text string, buttons ...RequestButton) RequestAttachmentPayload {
	return RequestAttachmentPayload{
		RequestAttachmentTemplate: RequestAttachmentTemplate{
			TemplateType: RequestAttachmentTemplateTypeButton,
			RequestAttachmentButtonTemplate: RequestAttachmentButtonTemplate{
				Text: text,
				Buttons: buttons,
			},
		},
	}
}

func NewListTemplate(topElementStyle TopElementStyle, elements ...RequestElement) RequestAttachmentPayload {
	switch topElementStyle {
	case TopElementStyleCompact:
	case TopElementStyleLarge:
	default:
		panic(fmt.Sprintf("Unknown topElementStyle: %v. Expected either 'large' or 'compact'.", topElementStyle))
	}

	if elementsLen := len(elements); elementsLen == 0 {
		panic("len(elements) == 0")
	} else if elementsLen > LIST_TEMPLATE_MAX_BUTTONS_COUNT {
		panic(fmt.Sprintf("len(elements)=%v > LIST_TEMPLATE_MAX_BUTTONS_COUNT=%v", elementsLen, LIST_TEMPLATE_MAX_BUTTONS_COUNT))
	} else {
		//els := make([]RequestElement, elementsLen)
		//copy(els, elements)
		return RequestAttachmentPayload{
			RequestAttachmentTemplate: RequestAttachmentTemplate{
				TemplateType: RequestAttachmentTemplateTypeList,
				RequestAttachmentListTemplate: RequestAttachmentListTemplate{
					RequestAttachmentGenericTemplate: RequestAttachmentGenericTemplate{Elements: elements},
					TopElementStyle: string(topElementStyle),
				},
			},
		}
	}
}

// https://developers.facebook.com/docs/messenger-platform/thread-settings/domain-whitelisting
type RequestWhitelistDomain struct {
	SettingType        string `json:"setting_type"` // Must be domain_whitelisting.
	DomainActionType   string `json:"domain_action_type"`
	WhitelistedDomains []string `json:"whitelisted_domains"`
}

func NewRequestWhitelistDomain(domainActionType string, domains... string) RequestWhitelistDomain {
	return RequestWhitelistDomain{
		SettingType: "domain_whitelisting",
		DomainActionType: domainActionType,
		WhitelistedDomains: domains,

	}
}