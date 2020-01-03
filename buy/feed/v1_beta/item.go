package v1_beta

import (
	"github.com/markdicksonjr/ebay"
	"os"
)

type GetItemFeedParameters struct {
	IsProduction  bool
	FeedScope     string // NEWLY_LISTED, ALL_ACTIVE
	CategoryId    string
	Date          string
	MarketplaceId string // https://developer.ebay.com/api-docs/buy/feed/overview.html#API
	Range         string // bytes=startpos-endpos (e.g. Range bytes=0-10485760)
}

func GetItemFeed(client ebay.Authorized, params GetItemFeedParameters) ([]byte, error) {

}
