package ebay

import (
	"os"
	"testing"
)

func TestAuthorized_Authorize(t *testing.T) {
	c := NewClient(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), RetrySettings{})

	err := c.Authorize("https://api.sandbox.ebay.com/identity/v1/oauth2/token", nil)
	if err != nil {
		t.Fatal(err)
	}
}
