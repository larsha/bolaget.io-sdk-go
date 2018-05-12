package sdk

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
)

const productsUrl = "/products/"

type ProductQueryParams struct {
	Limit          int     `url:"limit"`
	Offset         int     `url:"offset"`
	Ecological     bool    `url:"ecological"`
	Koscher        bool    `url:"koscher"`
	Ethical        bool    `url:"ethical"`
	YearFrom       int     `url:"year_from"`
	YearTo         int     `url:"year_to"`
	PriceFrom      float64 `url:"price_from"`
	PriceTo        float64 `url:"price_to"`
	VolumeFrom     float64 `url:"volume_from"`
	VolumeTo       float64 `url:"volume_to"`
	Assortment     string  `url:"assortment"`
	Sort           string  `url:"sort"`
	Name           string  `url:"name"`
	Type           string  `url:"type"`
	Style          string  `url:"style"`
	Provider       string  `url:"provider"`
	Producer       string  `url:"producer"`
	Origin         string  `url:"origin"`
	OriginCountry  string  `url:"origin_country"`
	Packaging      string  `url:"packaging"`
	ProductGroup   string  `url:"product_group"`
	Sealing        string  `url:"sealing"`
	Search         string  `url:"search"`
	SalesStartFrom string  `url:"sales_start_from"`
	SalesStartTo   string  `url:"sales_start_to"`
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type Product struct {
	Nr                 string  `json:"nr"`
	ArticleId          int     `json:"article_id"`
	ArticleNr          int     `json:"article_nr"`
	Name               string  `json:"name"`
	AdditionalName     string  `json:"additional_name"`
	Price              Price   `json:"price"`
	VolumeInMilliliter float64 `json:"volume_in_milliliter"`
	PricePerLiter      float64 `json:"price_per_liter"`
	SalesStart         string  `json:"sales_start"`
	Expired            bool    `json:"expired"`
	ProductGroup       string  `json:"product_group"`
	Type               string  `json:"type"`
	Style              string  `json:"style"`
	Packaging          string  `json:"packaging"`
	Sealing            string  `json:"sealing"`
	Origin             string  `json:"origin"`
	Producer           string  `json:"producer"`
	Provider           string  `json:"provider"`
	Year               int     `json:"year"`
	YearTested         string  `json:"year_tested"`
	Alcohol            string  `json:"alcohol"`
	Assortment         string  `json:"assortment"`
	Ecological         bool    `json:"ecological"`
	Ethical            bool    `json:"ethical"`
	Koscher            bool    `json:"koscher"`
}

func GetProduct(nr string) (Product, error) {
	var data Product
	url := URL + productsUrl + nr
	body, err := request(url)
	if err != nil {
		return data, err
	}

	data = Product{}
	e := json.Unmarshal(body, &data)
	return data, e
}

func GetProducts(p ProductQueryParams) ([]Product, error) {
	v, _ := query.Values(p)
	url := URL + productsUrl + "?" + v.Encode()

	body, err := request(url)
	if err != nil {
		return nil, err
	}

	products := make([]Product, 0)
	e := json.Unmarshal(body, &products)
	return products, e
}
