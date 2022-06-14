package utils

import (
	"net/url"
	"strconv"
	"strings"
)

/**
Example
Input - STRING - quote_amount=DESC&created_time=DESC
Output - map[string]string = {"quote_amount": "DESC", "created_time": "DESC"}
*/
func GetMapFromQuery(numStr string) (map[string]string, error) {
	var resMap = make(map[string]string)
	var mapStrings []string

	if len(numStr) > 0 {
		mapStrings = strings.Split(numStr, "&")

		for _, pair := range mapStrings {
			z := strings.Split(pair, "=")

			key, err := url.QueryUnescape(z[0])
			if err != nil {
				return nil, err
			}

			value, err := url.QueryUnescape(z[1])
			if err != nil {
				return nil, err
			}

			if len(key) > 0 && len(value) > 0 {
				resMap[key] = value
			}
		}
	}

	return resMap, nil
}

func GetInt(numStr string) (int, error) {

	res, err := strconv.Atoi(numStr)

	if err != nil {
		return 0, err
	}

	return res, nil
}
