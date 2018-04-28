package fbmbotapi

//go:generate ffjson $GOFILE

import "fmt"

// RequestNotificationType defines request notification type
type RequestNotificationType string

const (
	// RequestNotificationTypeRegular is for regular notifications
	RequestNotificationTypeRegular = RequestNotificationType("REGULAR")

	// RequestNotificationTypeSilentPush is for silent notifications
	RequestNotificationTypeSilentPush = RequestNotificationType("SILENT_PUSH")

	// RequestNotificationTypeNoPush is for no push notifications
	RequestNotificationTypeNoPush = RequestNotificationType("NO_PUSH")
)

const (
	// RequestAttachmentTypeImage is for image attachment
	RequestAttachmentTypeImage = "image"

	// RequestAttachmentTypeTemplate is for template attachment
	RequestAttachmentTypeTemplate = "template"
)

const (
	// RequestAttachmentTemplateTypeGeneric is generic
	RequestAttachmentTemplateTypeGeneric = "generic"
	// RequestAttachmentTemplateTypeList is list
	RequestAttachmentTemplateTypeList = "list"
	// RequestAttachmentTemplateTypeButton is button
	RequestAttachmentTemplateTypeButton = "button"
	// RequestAttachmentTemplateTypeReceipt is receipt
	RequestAttachmentTemplateTypeReceipt = "receipt"
)

const (
	// RequestButtonTypePostback os postback
	RequestButtonTypePostback = "postback"

	// RequestButtonTypeWebURL is for web url
	RequestButtonTypeWebURL = "web_url"
)

// Request is request
type Request struct {
	Recipient        RequestRecipient        `json:"recipient"`
	Message          RequestMessage          `json:"message"`
	NotificationType RequestNotificationType `json:"notification_type,omitempty"` // Optional: Push notification type: REGULAR, SILENT_PUSH, NO_PUSH
}

// RequestRecipient describes recipient
type RequestRecipient struct { // phone_number or id must be set
	ID          string `json:"id,omitempty"`           // Alex: changed from int64 on 21st November
	PhoneNumber string `json:"phone_number,omitempty"` // Phone number of the recipient with the format +1(212)555-2368
}

// RequestMessage is request message
type RequestMessage struct { // text or attachment must be set
	Text       string             `json:"text,omitempty"`       // Message text, must be UTF-8, 320 character limit
	Attachment *RequestAttachment `json:"attachment,omitempty"` // Attachment data
}

// RequestAttachment is request attachment
type RequestAttachment struct {
	Type    string                   `json:"type,omitempty"`
	Payload RequestAttachmentPayload `json:"payload,omitempty"`
}

// RequestAttachmentImage is request attachment image
type RequestAttachmentImage struct {
	URL string `json:"url,omitempty"`
}

// RequestAttachmentGenericTemplate is request attachment generic template
type RequestAttachmentGenericTemplate struct {
	Elements []RequestElement `json:"elements,omitempty"` // Data for each bubble in message
}

// RequestAttachmentListTemplate is request attachment list template
type RequestAttachmentListTemplate struct {
	RequestAttachmentGenericTemplate
	TopElementStyle string `json:"top_element_style,omitempty"`
}

// TopElementStyle defines top element style
type TopElementStyle string

const (
	// TopElementStyleCompact for compact top element
	TopElementStyleCompact TopElementStyle = "compact"

	// TopElementStyleLarge for large top element
	TopElementStyleLarge TopElementStyle = "large"
)

// RequestAttachmentButtonTemplate defines attachment button template
type RequestAttachmentButtonTemplate struct {
	Text    string          `json:"text,omitempty"`
	Buttons []RequestButton `json:"buttons,omitempty"`
}

// RequestAttachmentReceiptTemplate defines receipt template
type RequestAttachmentReceiptTemplate struct {
}

// RequestAttachmentTemplate defines attachment template
type RequestAttachmentTemplate struct {
	TemplateType                  string `json:"template_type,omitempty"`
	RequestAttachmentListTemplate        // Used for both List & Generic templates
	// RequestAttachmentGenericTemplate - conflicts with RequestAttachmentListTemplate on json.Marshal - silently not marshalling Elements.
	RequestAttachmentButtonTemplate
	RequestAttachmentReceiptTemplate
}

// RequestAttachmentPayload defines attachment payload
type RequestAttachmentPayload struct {
	RequestAttachmentImage
	RequestAttachmentTemplate
}

// RequestWebURLAction defines web URL action
type RequestWebURLAction struct {
	URL                 string             `json:"url,omitempty"` // For web_url buttons, this URL is opened in a mobile browser when the button is tapped
	FallbackURL         string             `json:"fallback_url,omitempty"`
	WebviewHeightRatio  WebviewHeightRatio `json:"webview_height_ratio,omitempty"`
	MessengerExtensions bool               `json:"messenger_extensions,omitempty"`
}

// RequestPostbackAction defines postback action
type RequestPostbackAction struct {
	Payload string `json:"payload,omitempty"` // For postback buttons, this data will be sent back to you via webhook
}

// RequestButton defines request button
type RequestButton struct {
	Type  string `json:"type"`  // Value is web_url or postback
	Title string `json:"title"` // Button title
	RequestWebURLAction
	RequestPostbackAction
}

// NewRequestPostbackButton creates new postback button
func NewRequestPostbackButton(title, payload string) RequestButton {
	return RequestButton{
		Type:  RequestButtonTypePostback,
		Title: title,
		RequestPostbackAction: RequestPostbackAction{Payload: payload},
	}
}

// NewRequestWebURLButton creates new URL button
func NewRequestWebURLButton(title, url string) RequestButton {
	return RequestButton{
		Type:                RequestButtonTypeWebURL,
		Title:               title,
		RequestWebURLAction: RequestWebURLAction{URL: url},
	}
}

// NewRequestWebURLButtonWithRatio create new web url button with ratio
func NewRequestWebURLButtonWithRatio(title, url string, webviewHeightRatio WebviewHeightRatio) RequestButton {
	return RequestButton{
		Type:  RequestButtonTypeWebURL,
		Title: title,
		RequestWebURLAction: RequestWebURLAction{
			URL:                url,
			WebviewHeightRatio: webviewHeightRatio,
		},
	}
}

// NewRequestWebExtentionURLButtonWithRatio creates new URL extention button
func NewRequestWebExtentionURLButtonWithRatio(title, url string, webviewHeightRatio WebviewHeightRatio) RequestButton {
	return RequestButton{
		Type:  RequestButtonTypeWebURL,
		Title: title,
		RequestWebURLAction: RequestWebURLAction{
			URL:                 url,
			MessengerExtensions: true,
			WebviewHeightRatio:  webviewHeightRatio,
		},
	}
}

// RequestDefaultAction defines default action
type RequestDefaultAction struct {
	Type string `json:"type"` // Value is web_url or postback
	RequestWebURLAction
	RequestPostbackAction
}

// NewDefaultActionWithWebURL creates new default action
func NewDefaultActionWithWebURL(action RequestWebURLAction) RequestDefaultAction {
	return RequestDefaultAction{
		Type:                RequestButtonTypeWebURL,
		RequestWebURLAction: action,
	}
}

// RequestElement defines request element
type RequestElement struct {
	Title         string                `json:"title"`               // Required: Bubble title
	Subtitle      string                `json:"subtitle,omitempty"`  // Optional: Bubble subtitle
	ImageURL      string                `json:"image_url,omitempty"` // Optional: Bubble image
	ItemURL       string                `json:"item_url,omitempty"`  // Optional: URL that is opened when bubble is tapped
	DefaultAction *RequestDefaultAction `json:"default_action,omitempty"`
	Buttons       []RequestButton       `json:"buttons,omitempty"` // Optional: Set of buttons that appear as call-to-actions
}

// NewRequestElementWithDefaultAction creates new request element with default action
func NewRequestElementWithDefaultAction(title, subtitle string, defaultAction RequestDefaultAction, buttons ...RequestButton) RequestElement {
	return RequestElement{
		Title:         title,
		Subtitle:      subtitle,
		DefaultAction: &defaultAction,
		Buttons:       buttons,
	}
}

// NewRequestElement creates new request element
func NewRequestElement(title, subtitle string, buttons ...RequestButton) RequestElement {
	return RequestElement{
		Title:    title,
		Subtitle: subtitle,
		Buttons:  buttons,
	}
}

// NewGenericTemplate creates new generic template
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

// NewButtonTemplate creates new button template
func NewButtonTemplate(text string, buttons ...RequestButton) RequestAttachmentPayload {
	return RequestAttachmentPayload{
		RequestAttachmentTemplate: RequestAttachmentTemplate{
			TemplateType: RequestAttachmentTemplateTypeButton,
			RequestAttachmentButtonTemplate: RequestAttachmentButtonTemplate{
				Text:    text,
				Buttons: buttons,
			},
		},
	}
}

// NewListTemplate creates new list template
func NewListTemplate(topElementStyle TopElementStyle, elements ...RequestElement) RequestAttachmentPayload {
	switch topElementStyle {
	case TopElementStyleCompact:
	case TopElementStyleLarge:
	default:
		panic(fmt.Sprintf("Unknown topElementStyle: %v. Expected either 'large' or 'compact'.", topElementStyle))
	}

	if elementsLen := len(elements); elementsLen == 0 {
		panic("len(elements) == 0")
	} else if elementsLen > listTemplateMaxButtonsCount {
		panic(fmt.Sprintf("len(elements)=%v > listTemplateMaxButtonsCount=%v", elementsLen, listTemplateMaxButtonsCount))
	} else {
		//els := make([]RequestElement, elementsLen)
		//copy(els, elements)
		return RequestAttachmentPayload{
			RequestAttachmentTemplate: RequestAttachmentTemplate{
				TemplateType: RequestAttachmentTemplateTypeList,
				RequestAttachmentListTemplate: RequestAttachmentListTemplate{
					RequestAttachmentGenericTemplate: RequestAttachmentGenericTemplate{Elements: elements},
					TopElementStyle:                  string(topElementStyle),
				},
			},
		}
	}
}

// RequestWhitelistDomain see https://developers.facebook.com/docs/messenger-platform/thread-settings/domain-whitelisting
type RequestWhitelistDomain struct {
	SettingType        string   `json:"setting_type"` // Must be domain_whitelisting.
	DomainActionType   string   `json:"domain_action_type"`
	WhitelistedDomains []string `json:"whitelisted_domains"`
}

// NewRequestWhitelistDomain creates new request for whitelisted domains
func NewRequestWhitelistDomain(domainActionType string, domains ...string) RequestWhitelistDomain {
	return RequestWhitelistDomain{
		SettingType:        "domain_whitelisting",
		DomainActionType:   domainActionType,
		WhitelistedDomains: domains,
	}
}
