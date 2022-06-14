package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	JsonContentType = "application/json"
	FormContentType = "application/x-www-form-urlencoded"
)

func Post(url string, data interface{}, headers map[string]string, contentType string, needToMarshall bool) ([]byte, int, error) {
	var body []byte

	var err error

	if needToMarshall {
		body, err = json.Marshal(data)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	} else {
		// assume that we receive bytes
		body = data.([]byte)
	}

	var req, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	req.Header.Add("content-type", contentType)

	getHeaders(req, headers)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer func() {
		res.Body.Close()
	}()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	if res.StatusCode != http.StatusOK {
		return response, res.StatusCode, errors.New(res.Status)
	}

	return response, res.StatusCode, nil
}

func PostForm(url string, data map[string]string, headers map[string]string) ([]byte, int, error) {
	var req, _ = http.NewRequest("POST", url, strings.NewReader(getEncodedData(data)))

	req.Header.Add("content-type", FormContentType)

	getHeaders(req, headers)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	defer func() {
		res.Body.Close()
	}()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	if res.StatusCode != http.StatusOK {
		return response, res.StatusCode, errors.New(res.Status)
	}

	return response, res.StatusCode, nil
}

func getEncodedData(data map[string]string) string {

	var strData = ""

	for name, value := range data {
		strData += name + "=" + value + "&"
	}

	return strings.Trim(strData, "&")
}

func getHeaders(req *http.Request, headers map[string]string) {
	for name, value := range headers {
		req.Header.Add(name, value)
	}
}
