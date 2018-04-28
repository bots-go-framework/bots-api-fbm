package fbmbotapi

import (
	"github.com/pquerna/ffjson/ffjson"
	"testing"
)

func TestNewMenuItemWebUrl(t *testing.T) {
	menuItem := NewMenuItemWebUrl("Item #1", "https://debtstracker.io/", WebviewHeightRatioTall, true, false)

	bytes, err := ffjson.Marshal(menuItem)
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `{"type":"web_url","title":"Item #1","url":"https://debtstracker.io/","webview_height_ratio":"tall"}` {
		t.Errorf("Unexpected JSON: %v", string(bytes))
	}

	menuItem = NewMenuItemWebUrl("Item #2", "https://splitbill.co/", WebviewHeightRatioTall, false, true)

	bytes, err = ffjson.Marshal(menuItem)
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `{"type":"web_url","title":"Item #2","url":"https://splitbill.co/","webview_height_ratio":"tall","webview_share_button":"hide","messenger_extensions":true}` {
		t.Errorf("Unexpected JSON: %v", string(bytes))
	}
}
