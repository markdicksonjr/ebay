package ebay

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Authorized) Post(path string, body []byte) (Response, error) {
	return c.doPost(path, body, 0)
}

func (c *Authorized) BuildPostRequest(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if c.clientId != "" && c.clientSecret != "" {
		req.Header.Set("Authorization", "Bearer " + c.accessToken)
	}

	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func (c *Authorized) doPost(path string, body []byte, retryCount int) (Response, error) {

	// build the request, do not retry if an error happened
	req, err := c.BuildPostRequest(path, body)
	if err != nil {
		return Response{}, err
	}

	// send the initial request
	client := &http.Client{}
	if resp, err := client.Do(req); err != nil {

		// if an error happened, attempt a retry, if we're configured and haven't done so too often
		if retryCount < c.retrySettings.MaxRetries {
			c.applyRetryWait(retryCount + 1)
			return c.doPost(path, body, retryCount+1)
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
				return c.doPost(path, body, retryCount+1)
			}

			// otherwise, return the error with whatever contents could be interpreted
			return responseContents, err
		} else {

			// if no error occurred

			// if we are configured to retry when we don't get JSON back,
			// try to parse the contents to be sure they're JSON
			//if c.retrySettings.RetryIfPostResponseNotJSON {
				var mapContents map[string]interface{}

				// if we can't parse the body to JSON, attempt a retry
				if err := json.Unmarshal(responseContents.Body, &mapContents); err != nil {

					// if we haven't exceeded our r
					//etry count, retry
					if retryCount < c.retrySettings.MaxRetries {
						c.applyRetryWait(retryCount + 1)
						return c.doPost(path, body, retryCount+1)
					}

					// otherwise, return the error
					return responseContents, err
				}
			//}

			// return the response
			return responseContents, err
		}
	}
}
