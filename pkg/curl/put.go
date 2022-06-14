package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Put(url string, data interface{}, username, password, contentType string, needToMarshall bool) ([]byte, int, error) {
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

	var req, _ = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))

	req.Header.Add("content-type", contentType)

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

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
