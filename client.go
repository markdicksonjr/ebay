package ebay

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Authorized struct {
	clientId        string
	clientSecret    string
	accessToken     string
	retrySettings   RetrySettings
	TokenExpiration time.Time
	TokenType       string
}

type RetrySettings struct {
	MaxRetries          int
	MinMsBetweenRetries int
}

func NewClient(clientId, clientSecret string, retrySettings RetrySettings) *Authorized {
	return &Authorized{
		clientId:      clientId,
		clientSecret:  clientSecret,
		retrySettings: retrySettings,
	}
}

// NOTE: retryCount provided in case we want to implement exponential backoff or similar algorithm
func (c *Authorized) applyRetryWait(retryCount int) {
	time.Sleep(time.Duration(c.retrySettings.MinMsBetweenRetries) * time.Millisecond)
}

func GetResponseContents(resp *http.Response) (Response, error) {
	var body []byte
	var err error

	if resp.Body != nil {
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
	}

	return Response{
		Body:       body,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
	}, err
}
