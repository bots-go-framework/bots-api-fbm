package fbm_api

import (
	"net/http"
	"github.com/pquerna/ffjson/ffjson"
	"bytes"
	"fmt"
	"io/ioutil"
	"context"
	"github.com/strongo/log"
	"github.com/pkg/errors"
)

const ContentTypeApplicationJson = "application/json"

const (
	EndpointMeMessage = "me/messages"
	EndpointMeMessengerProfile = "me/messenger_profile"
)

type GraphAPI struct {
	httpClient *http.Client
	AccessToken string
}

func NewGraphApi(httpClient *http.Client, accessToken string) GraphAPI {
	return GraphAPI{
		httpClient: httpClient,
		AccessToken: accessToken,
	}
}

func (graphAPI GraphAPI) apiUrl(endpoint string) string {
	return fmt.Sprintf("https://graph.facebook.com/v2.6/%v/?access_token=%v", endpoint, graphAPI.AccessToken)
}

func (graphAPI GraphAPI) SetGetStarted(c context.Context, message GetStartedMessage) (error) {
	return graphAPI.postMessage(c, EndpointMeMessengerProfile, &message)
}

func (graphAPI GraphAPI) SetPersistentMenu(c context.Context, message PersistentMenuMessage) (error){
	return graphAPI.postMessage(c, EndpointMeMessengerProfile, &message)
}

func (graphAPI GraphAPI) SetWhitelistedDomains(c context.Context, message WhitelistedDomainsMessage) (error){
	return graphAPI.postMessage(c, EndpointMeMessengerProfile, &message)
}

func (graphAPI GraphAPI) SendMessage(c context.Context, request Request) (error){
	return graphAPI.postMessage(c, EndpointMeMessage, &request)
}

func (graphAPI GraphAPI) postMessage(c context.Context, endpoint string, message interface{}) error {
	content, err := ffjson.MarshalFast(message)
	if err != nil {
		ffjson.Pool(content)
		return err
	}

	apiUrl := graphAPI.apiUrl(endpoint)
	log.Debugf(c, "Posting to FB API: %v\n%v", apiUrl, string(content))

	resp, err := graphAPI.httpClient.Post(apiUrl, ContentTypeApplicationJson, bytes.NewReader(content))
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