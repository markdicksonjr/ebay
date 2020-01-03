package v1

import "github.com/markdicksonjr/ebay"

// https://api.ebay.com/buy/browse/v1/item_summary/search?q=drone&limit=3

type SearchParams struct {
	IsProduction  bool
}

func Search(client ebay.Authorized, params SearchParams) ([]byte, error) {
	var url string
	if params.IsProduction {
		url = "https://api.ebay.com/buy/browse/v1/item_summary/search"
	} else {
		url = "https://api.sandbox.ebay.com/buy/browse/v1/item_summary/search"
	}

	url += "?q=Shark"

	//q=string&
	//	gtin=string&
	//	charity_ids=string&
	//	fieldgroups=string&
	//	compatibility_filter=CompatibilityFilter&
	//	category_ids=string&
	//	filter=FilterField&
	//	sort=SortField&
	//	limit=string&
	//	offset=string&
	//	aspect_filter=AspectFilter&
	//	epid=string
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}


	body := string(res.Body)
	_ = body

	if res.StatusCode != 200 {
		return nil, err
	}

	return nil, nil
}