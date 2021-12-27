package http_client

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"strconv"
)

var client *resty.Client
var authToken = "BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F"

func Setup() {
	client = resty.New()
}

func GetRequest(queryParams map[string]string, url string) ([]byte, error) {
	resp, err := client.R().
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		SetAuthToken(authToken).
		Get(url)
	return validateAndReturnResponse(resp, err)
}

func validateAndReturnResponse(res *resty.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	status := res.StatusCode()
	if status >= 200 && status < 300 {
		return res.Body(), nil
	} else {
		return nil, errors.New("http request failed with status code" + strconv.Itoa(status))
	}
}
