package fbmbotapi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/strongo/log"
	"io/ioutil"
	"net/http"
)

const contentTypeApplicationJSON = "application/json"

const (
	endpointMeMessage          = "me/messages"
	endpointMeMessengerProfile = "me/messenger_profile"
)

// GraphAPI is Facebook API client
type GraphAPI struct {
	httpClient  *http.Client
	AccessToken string
}

// NewGraphAPI creates new API client
func NewGraphAPI(httpClient *http.Client, accessToken string) GraphAPI {
	return GraphAPI{
		httpClient:  httpClient,
		AccessToken: accessToken,
	}
}

func (graphAPI GraphAPI) apiURL(endpoint string) string {
	return fmt.Sprintf("https://graph.facebook.com/v2.6/%v/?access_token=%v", endpoint, graphAPI.AccessToken)
}

// SetGetStarted sets get started message
func (graphAPI GraphAPI) SetGetStarted(c context.Context, message GetStartedMessage) error {
	return graphAPI.postMessage(c, endpointMeMessengerProfile, &message)
}

// SetPersistentMenu sets persistent menu
func (graphAPI GraphAPI) SetPersistentMenu(c context.Context, message PersistentMenuMessage) error {
	return graphAPI.postMessage(c, endpointMeMessengerProfile, &message)
}

// SetWhitelistedDomains sets whitelisted domains
func (graphAPI GraphAPI) SetWhitelistedDomains(c context.Context, message WhitelistedDomainsMessage) error {
	return graphAPI.postMessage(c, endpointMeMessengerProfile, &message)
}

// SendMessage sends message
func (graphAPI GraphAPI) SendMessage(c context.Context, request Request) error {
	return graphAPI.postMessage(c, endpointMeMessage, &request)
}

func (graphAPI GraphAPI) postMessage(c context.Context, endpoint string, message interface{}) error {
	content, err := ffjson.MarshalFast(message)
	if err != nil {
		ffjson.Pool(content)
		return err
	}

	apiURL := graphAPI.apiURL(endpoint)
	log.Debugf(c, "Posting to FB API: %v\n%v", apiURL, string(content))

	resp, err := graphAPI.httpClient.Post(apiURL, contentTypeApplicationJSON, bytes.NewReader(content))
	ffjson.Pool(content)
	if err != nil {
		return err
	}
	var respData []byte
	if resp.Body != nil {
		defer resp.Body.Close()
		if respData, err = ioutil.ReadAll(resp.Body); err != nil {
			err = errors.WithMessage(err, "Failed to read response body")
			return err
		}
	}
	log.Debugf(c, "Response from FB API (status=%v): %v", resp.Status, string(respData))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = fmt.Errorf("FB API returned unexpected HTTP status code: %d", resp.StatusCode)
		return err
	}
	return nil
}
