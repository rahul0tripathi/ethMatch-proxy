package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethMatch/proxy/config"
	"net/http"
)

const (
	endpoint = "https://api.pinata.cloud/pinning/pinJSONToIPFS"
)

type (
	requestBody struct {
		PinataContent interface{} `json:"pinataContent"`
	}
	responseBody struct {
		IpfsHash string `json:"IpfsHash"`
	}
)

var (
	headers http.Header
	client  http.Client
)

func init() {
	headers = http.Header{
		"pinata_api_key":        []string{config.ENV.Pinata.ApiKey},
		"pinata_secret_api_key": []string{config.ENV.Pinata.ApiSecret},
	}
	client = http.Client{}
}
func PinJson(gameLog interface{}) (CID string, err error) {
	var body []byte
	reqBody := requestBody{PinataContent: gameLog}
	fmt.Println(reqBody)
	body, err = json.Marshal(reqBody)
	fmt.Println(string(body),err)
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header = headers
	resp, err := client.Do(req)
	fmt.Println(err)
	if err != nil {
		return
	}
	response := responseBody{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}
	CID = response.IpfsHash
	return
}
