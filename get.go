package ebay

import (
	"net/http"
)

func (c *Authorized) Get(path string) (Response, error) {
	return c.doGet(path, 0)
}

func (c *Authorized) buildGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.clientId != "" && c.clientSecret != "" {
		req.Header.Set("Authorization", "Bearer " + c.accessToken)
	}

	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func (c *Authorized) doGet(path string, retryCount int) (Response, error) {

	// build the request, do not retry if an error happened
	req, err := c.buildGetRequest(path)
	if err != nil {
		return Response{}, err
	}

	// send the initial request
	client := &http.Client{}
	if resp, err := client.Do(req); err != nil {

		// if an error happened, attempt a retry, if we're configured and haven't done so too often
		if retryCount < c.retrySettings.MaxRetries {
			c.applyRetryWait(retryCount + 1)
			return c.doGet(path, retryCount+1)
		}

		// otherwise, return the error
		return Response{}, err
	} else {

		// if no error occurred with the request, start processing the response

		// if an error happened while processing the response
		if responseContents, err := GetResponseContents(resp); err != nil {

			// if we're configured to do so, and haven't tried too many times already, do a retry
			if retryCount < c.retrySettings.MaxRetries {
				c.applyRetryWait(retryCount + 1)
				return c.doGet(path, retryCount+1)
			}

			// otherwise, return the error with whatever contents could be interpreted
			return responseContents, err
		} else {
			return responseContents, err
		}
	}
}
