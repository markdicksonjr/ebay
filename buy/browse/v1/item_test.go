package v1

import (
	"github.com/markdicksonjr/ebay"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	c := ebay.NewClient(os.Getenv("CLIENT_ID"), os.Getenv("APP_AUTH_TOKEN"), ebay.RetrySettings{})

	res, err := Search(*c, SearchParams{
		IsProduction: true,
	})
	_ = res
	if err != nil {
		t.Fatal(err)
	}
}
