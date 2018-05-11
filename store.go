package sdk

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type StoreQueryParams struct {
	Limit   int    `url:"limit"`
	Offset  int    `url:"offset"`
	Labels  string `url:"labels"`
	Search  string `url:"search"`
	Type    string `url:"type"`
	Name    string `url:"name"`
	City    string `url:"city"`
	Country string `url:"country"`
	Address string `url:"address"`
	Sort    string `url:"sort"`
}

type Store struct {
	Nr                string   `json:"nr"`
	City              string   `json:"city"`
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	AdditionalAddress string   `json:"additional_address"`
	ZipCode           string   `json:"zip_code"`
	County            string   `json:"county"`
	Phone             string   `json:"phone"`
	ShopType          string   `json:"shop_type"`
	Services          string   `json:"services"`
	Labels            []string `json:"labels"`
	OpeningHours      []string `json:"opening_hours"`
	RT90X             int      `json:"RT90x"`
	RT90Y             int      `json:"RT90y"`
}

func GetStore(nr string) (Store, error) {
	var data Store
	url := URL + "/stores/" + nr
	body, err := request(url)
	if err != nil {
		return data, err
	}

	data = Store{}
	e := json.Unmarshal(body, &data)
	return data, e
}

func GetStores(params StoreQueryParams) ([]Store, error) {
	v, _ := query.Values(params)
	url := URL + "/stores/?" + v.Encode()

	body, err := request(url)
	if err != nil {
		return nil, err
	}

	stores := make([]Store, 0)
	e := json.Unmarshal(body, &stores)
	return stores, e
}
