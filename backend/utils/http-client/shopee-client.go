package httpclient

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type shopeeClient struct {
	Header http.Header
}

// NewShopeeClient ...
func NewShopeeClient() HTTPClient {
	return shopeeClient{
		http.Header{
			"Referer": []string{"Shopee.vn"},
		},
	}
}

func (s shopeeClient) Get(url string) ([]byte, error) {
	client := http.Client{}
	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, nil
	}
	rq.Header = s.Header
	response, err := client.Do(rq)
	if err != nil {
		return []byte{}, err
	}
	if response.StatusCode != 200 {
		errMsg, _ := ioutil.ReadAll(response.Body)
		return []byte{}, errors.New(string(errMsg))
	}
	return ioutil.ReadAll(response.Body)

}
