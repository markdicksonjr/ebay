package ebay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AuthorizationResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

const ScopeApi = "https://api.ebay.com/oauth/api_scope"
const ScopeBuyItemFeed = "https://api.ebay.com/oauth/api_scope/buy.item.feed"
const ScopeSellInventory = "https://api.ebay.com/oauth/api_scope/sell.inventory"
const ScopeSellInventoryReadOnly = "https://api.ebay.com/oauth/api_scope/sell.inventory.readonly"

// authUrl is likely either https://api.sandbox.ebay.com/identity/v1/oauth2/token or the live one
func (c *Authorized) Authorize(authUrl string, scopes []string) error {
	if len(c.clientId) == 0 {
		return errors.New("client id not provided to Authorized client")
	}

	if len(c.clientSecret) == 0 {
		return errors.New("client secret not provided to Authorized client")
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	if len(scopes) == 0 {
		scopes = []string{ScopeApi}
	}

	data.Set("scope", strings.Join(scopes, " "))

	req, err := http.NewRequest("POST", authUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.clientId+":"+c.clientSecret)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	if resp, err := client.Do(req); err != nil {
		return err
	} else {
		var contents Response
		contents, err = GetResponseContents(resp)
		if err != nil {
			return err
		}

		if contents.StatusCode != 200 {
			return errors.New("got " + strconv.Itoa(contents.StatusCode) + " status code with body = " + string(contents.Body))
		}

		result := AuthorizationResponse{}
		if err = json.Unmarshal(contents.Body, &result); err != nil {
			return err
		}

		c.accessToken = result.AccessToken
		c.TokenExpiration = time.Now().Add(time.Duration(result.ExpiresIn) * time.Millisecond)
		c.TokenType = result.TokenType

		if len(c.accessToken) == 0 {
			return errors.New("could not get access token from response")
		}

		return nil
	}
}
