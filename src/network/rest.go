package network

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type RestInterface interface {
}

// CallRestService - Generic function to fetch data from a rest service
func CallRestService(requestMethod, httpEndPoint, requestType, responseType string,
	requestHeaders, queryParams map[string]string, requestCookies []*http.Cookie, body interface{}, client http.Client) ([]byte, error, int) {

	var (
		responseData    []byte
		mainError, err  error
		request         *http.Request
		statusCode      int
		ioBody          io.Reader
		requestBodyJson string
	)

	if body != nil && requestMethod == "POST" {
		if requestType == "Proto" {
			ioBody = bytes.NewReader(body.([]byte))
		} else if requestType == "Json" {
			requestBodyJson = body.(string)
			ioBody = strings.NewReader(requestBodyJson)
		}
	}

	request, err = http.NewRequest(requestMethod, httpEndPoint, ioBody)

	if err != nil {
		// log error in creating request
		if responseType == "Proto" {
			responseData = make([]byte, 0)
		} else if responseType == "JSON" {
			responseData = []byte("{}")
		}
		return responseData, err, statusCode
	}
	for key, val := range requestHeaders {

		request.Header.Add(key, val)
	}

	// Adding query params into the request URL
	if len(queryParams) > 0 {
		// Query params container
		q := request.URL.Query()
		for pKey, pVal := range queryParams {
			q.Add(pKey, pVal)
		}

		// Encoding (stringify) query params
		request.URL.RawQuery = q.Encode()
	}

	for _, cookie := range requestCookies {
		request.AddCookie(cookie)
	}

	var resp *http.Response
	resp, err = client.Do(request)
	if err == nil {
		defer resp.Body.Close()
	}

	if resp != nil {
		statusCode = resp.StatusCode
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err == nil {
			responseData, err = ioutil.ReadAll(reader)
			defer reader.Close()
		}

	default:
		responseData, err = ioutil.ReadAll(resp.Body)
	}
	return responseData, mainError, statusCode
}
