package toa_api

import (
	"io/ioutil"
	"net/http"
)

type TOAClient struct {
	APIKey  string
	AppName string
	client  *http.Client
}

func NewTOAClient(APIKey string, name string) *TOAClient {
	return &TOAClient{APIKey: APIKey, AppName: name, client: &http.Client{}}
}

func (c *TOAClient) Get(APIUri string) ([]byte, error) {
	req, _ := http.NewRequest("GET", "https://theorangealliance.org/api/"+APIUri, nil)
	// Set TOA headers as listed in https://theorangealliance.org/apidocs/get
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TOA-Key", c.APIKey)
	req.Header.Set("X-Application-Origin", c.AppName)
	// send request
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	// Read body
	var body []byte
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
