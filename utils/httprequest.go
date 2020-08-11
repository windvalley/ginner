package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func HttpRequest(method, url string, postData interface{}) (string, error) {
	jsonStr, err := json.Marshal(postData)
	if err != nil {
		return "", err
	}
	body := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// PostWithUrlencoded
//     e.g. payload := url.Values{}
// payload.Set("key", value1)
// payload.Add("key", value2)
func PostWithUrlencoded(api string, payload url.Values) (string, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		api,
		strings.NewReader(payload.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Add(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), err
}

// PostWithFormdata
//     e.g. postData := map[string]string{"key":"value"}
func PostWithFormdata(url string, postData map[string]string) (string, error) {
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	for k, v := range postData {
		mw.WriteField(k, v)
	}
	defer mw.Close()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", mw.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), err
}
