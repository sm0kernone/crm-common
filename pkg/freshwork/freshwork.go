package freshwork

import (
	"github.com/sm0kernone/crm-common/pkg/curl"
)

func UpdateTicket(data map[string]interface{}, url, username, password string) ([]byte, int, error) {
	res, statusCode, err := curl.Put(url, data, username, password, curl.JsonContentType, true)

	if err != nil {
		return nil, statusCode, err
	}

	return res, statusCode, nil
}

func CreateTicket(data map[string]interface{}, url, username, password string) ([]byte, int, error) {
	res, statusCode, err := curl.PostAuth(url, data, username, password, curl.JsonContentType, true)

	if err != nil {
		return nil, statusCode, err
	}

	return res, statusCode, nil
}

func Get(url string, data map[string]interface{}, username, password string) ([]byte, int, error) {
	res, statusCode, err := curl.GetAuth(url, data, username, password, curl.JsonContentType, true)

	if err != nil {
		return nil, statusCode, err
	}

	return res, statusCode, nil
}
