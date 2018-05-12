package sdk

import (
	"io/ioutil"
	"net/http"
	"time"
)

const URL = "https://bolaget.io/v1"

func request(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), err
}
