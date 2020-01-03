package ebay

import (
	"os"
	"strconv"
	"testing"
)

func TestAuthorized_Get(t *testing.T) {
	c := NewClient(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), RetrySettings{})

	err := c.Authorize("https://api.sandbox.ebay.com/identity/v1/oauth2/token", []string{
		ScopeApi,
		ScopeBuyItemFeed,
	})
	if err != nil {
		t.Fatal(err)
	}

	if r, err := c.Get("https://api.sandbox.ebay.com/sell/inventory/v1/inventory_item"); err != nil {
		t.Fatal(err)
	} else {
		if r.StatusCode != 200 {
			t.Fatal("response failed with status code ", strconv.Itoa(r.StatusCode))
		}
	}
}
