package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/pquerna/ffjson/ffjson"
)

// HTTPRequest get body of response that io.ReadCloser type
func HTTPRequest(
	method, url string,
	postData interface{},
) (io.ReadCloser, error) {
	jsonStr, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

// PostWithUrlencoded post with body of form-urlencoded, and get response body.
//    e.g. payload := url.Values{}
// payload.Set("key", value1)
// payload.Add("key", value2)
func PostWithUrlencoded(api string, payload url.Values) (io.ReadCloser, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		api,
		strings.NewReader(payload.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, err
}

// PostWithFormdata post with body of form-data, and get response body.
//     e.g. postData := map[string]string{"key":"value"}
func PostWithFormdata(
	url string,
	postData map[string]string,
) (io.ReadCloser, error) {
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	for k, v := range postData {
		mw.WriteField(k, v)
	}
	defer mw.Close()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", mw.FormDataContentType())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, err
}

// GetBodyStringData transfer io.ReadCloser to string.
func GetBodyStringData(resBody io.ReadCloser) (string, error) {
	defer resBody.Close()
	data, err := ioutil.ReadAll(resBody)
	if err != nil {
		return "", nil
	}
	return string(data), nil
}

// GetBodyStructData parameter structData is a Struct that contains the fields which you want from resBody.
//    e.g.:
//type Response struct {
//	Code   int               `json:"code"`
//	Msg    string            `json:"msg"`
//	Data   interface{}       `json:"data"`
//}
//response := new(Response)
//data, err := util.GetBodyStructData(res, response)
//if err != nil {
//return err
//}
//v := data.(*Response)
func GetBodyStructData(
	resBody io.ReadCloser,
	structData interface{},
) (interface{}, error) {
	defer resBody.Close()
	decoder := ffjson.NewDecoder()
	if err := decoder.DecodeReader(resBody, &structData); err != nil {
		return nil, err
	}

	return structData, nil
}
