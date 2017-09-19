package fbm_api

//go:generate ffjson $GOFILE

type PersistentMenuMessage struct {
	PersistentMenus []PersistentMenu `json:"persistent_menu"`
}

type PersistentMenu struct {
	Locale                string `json:"locale"`
	ComposerInputDisabled bool `json:"composer_input_disabled"`
	CallToActions         []MenuItem `json:"call_to_actions,omitempty"`
}

func NewPersistentMenu(locale string, composerInputDisabled bool, callToActions ...MenuItem) PersistentMenu {
	menu := PersistentMenu{
		Locale: locale,
		ComposerInputDisabled: composerInputDisabled,
		CallToActions: callToActions,
	}
	return menu
}

type MenuItemType string

const (
	MenuItemTypeWebUrl MenuItemType = "web_url"
	MenuItemTypePostback MenuItemType = "postback"
	MenuItemTypeNested MenuItemType = "nested"
)

type WebviewHeightRatio string

const (
	WebviewHeightRatioCompact WebviewHeightRatio = "compact"
	WebviewHeightRatioTall WebviewHeightRatio = "tall"
	WebviewHeightRatioFull WebviewHeightRatio = "full"
)

const WebviewShareButtonHide = "hide"

type MenuItem interface {
	MenuItemType() MenuItemType
}

type MenuItemBase struct {
	Type MenuItemType `json:"type"`
	Title string `json:"title"`
}

func (mi MenuItemBase) MenuItemType() MenuItemType {
	return mi.Type
}

type MenuItemWebUrl struct {
	MenuItemBase
	Url string `json:"url"`
	WebviewHeightRatio WebviewHeightRatio `json:"webview_height_ratio"`
	WebviewShareButton string `json:"webview_share_button,omitempty"`
	MessengerExtensions bool `json:"messenger_extensions,omitempty"`
}

func NewMenuItemWebUrl(title, url string, webviewHeightRatio WebviewHeightRatio, shareButton, messengerExtensions bool) MenuItemWebUrl {
	menuItem := MenuItemWebUrl{
		MenuItemBase: MenuItemBase{
			Type: MenuItemTypeWebUrl,
			Title: title,
		},
		Url: url,
		WebviewHeightRatio: webviewHeightRatio,
		MessengerExtensions: messengerExtensions,
	}
	if !shareButton {
		menuItem.WebviewShareButton = WebviewShareButtonHide
	}
	return menuItem
}

type MenuItemPostback struct {
	MenuItemBase
	Payload string `json:"payload"`
}

type MenuItemNested struct {
	MenuItemBase
	CallToActions         []MenuItem `json:"call_to_actions"`
}

func NewMenuItemNested(title string, callToActions ...MenuItem) MenuItemNested {
	menuItem := MenuItemNested{
		MenuItemBase: MenuItemBase{
			Type: MenuItemTypeNested,
			Title: title,
		},
		CallToActions: callToActions,
	}
	return menuItem
}

func NewMenuItemPostback(title, payload string) MenuItemPostback {
	menuItem := MenuItemPostback{
		MenuItemBase: MenuItemBase{
			Type: MenuItemTypePostback,
			Title: title,
		},
		Payload: payload,
	}
	return menuItem
}