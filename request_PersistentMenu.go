package fbmbotapi

//go:generate ffjson $GOFILE

// PersistentMenuMessage is persistent menus message
type PersistentMenuMessage struct {
	PersistentMenus []PersistentMenu `json:"persistent_menu"`
}

// PersistentMenu is mersistent menu
type PersistentMenu struct {
	Locale                string     `json:"locale"`
	ComposerInputDisabled bool       `json:"composer_input_disabled"`
	CallToActions         []MenuItem `json:"call_to_actions,omitempty"`
}

// NewPersistentMenu creates new persistent menu
func NewPersistentMenu(locale string, composerInputDisabled bool, callToActions ...MenuItem) PersistentMenu {
	menu := PersistentMenu{
		Locale:                locale,
		ComposerInputDisabled: composerInputDisabled,
		CallToActions:         callToActions,
	}
	return menu
}

// MenuItemType is type of persistent menu item
type MenuItemType string

const (
	// MenuItemTypeWebURL is for URL menu items
	MenuItemTypeWebURL MenuItemType = "web_url"
	// MenuItemTypePostback is for postback menu items
	MenuItemTypePostback MenuItemType = "postback"
	// MenuItemTypeNested is for nested menu items
	MenuItemTypeNested MenuItemType = "nested"
)

// WebviewHeightRatio defines height ratio
type WebviewHeightRatio string

const (
	// WebviewHeightRatioCompact is compact height ratio
	WebviewHeightRatioCompact WebviewHeightRatio = "compact"
	// WebviewHeightRatioTall is tall height ratio
	WebviewHeightRatioTall WebviewHeightRatio = "tall"
	// WebviewHeightRatioFull is full height ratio
	WebviewHeightRatioFull WebviewHeightRatio = "full"
)

// WebviewShareButtonHide to hide button
const WebviewShareButtonHide = "hide"

// MenuItem is menu item
type MenuItem interface {
	MenuItemType() MenuItemType
}

// MenuItemBase holds menu item base properties
type MenuItemBase struct {
	Type  MenuItemType `json:"type"`
	Title string       `json:"title"`
}

// MenuItemType returns type of menu item
func (mi MenuItemBase) MenuItemType() MenuItemType {
	return mi.Type
}

// MenuItemWebURL is URL menu item
type MenuItemWebURL struct {
	MenuItemBase
	URL                 string             `json:"url"`
	WebviewHeightRatio  WebviewHeightRatio `json:"webview_height_ratio"`
	WebviewShareButton  string             `json:"webview_share_button,omitempty"`
	MessengerExtensions bool               `json:"messenger_extensions,omitempty"`
}

// NewMenuItemWebUrl creates new URL menu item
func NewMenuItemWebUrl(title, url string, webviewHeightRatio WebviewHeightRatio, shareButton, messengerExtensions bool) MenuItemWebURL {
	menuItem := MenuItemWebURL{
		MenuItemBase: MenuItemBase{
			Type:  MenuItemTypeWebURL,
			Title: title,
		},
		URL:                 url,
		WebviewHeightRatio:  webviewHeightRatio,
		MessengerExtensions: messengerExtensions,
	}
	if !shareButton {
		menuItem.WebviewShareButton = WebviewShareButtonHide
	}
	return menuItem
}

// MenuItemPostback is postback menu item
type MenuItemPostback struct {
	MenuItemBase
	Payload string `json:"payload"`
}

// MenuItemPostback is nested menu item
type MenuItemNested struct {
	MenuItemBase
	CallToActions []MenuItem `json:"call_to_actions"`
}

// NewMenuItemNested creates nested menu item
func NewMenuItemNested(title string, callToActions ...MenuItem) MenuItemNested {
	menuItem := MenuItemNested{
		MenuItemBase: MenuItemBase{
			Type:  MenuItemTypeNested,
			Title: title,
		},
		CallToActions: callToActions,
	}
	return menuItem
}

// NewMenuItemPostback creates postback menu item
func NewMenuItemPostback(title, payload string) MenuItemPostback {
	menuItem := MenuItemPostback{
		MenuItemBase: MenuItemBase{
			Type:  MenuItemTypePostback,
			Title: title,
		},
		Payload: payload,
	}
	return menuItem
}
