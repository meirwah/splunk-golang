package splunk

import (
        "bytes"
        "net/http"
        "net/url"
        "io"
        "io/ioutil"
        "crypto/tls"
)

/*
 * HTTP helper methods
 */

func (conn SplunkConnection) httpGet(url string, data *url.Values) (string, error) {
	if response, err := conn.httpCall(url,"GET",data);err != nil {
		return "", err
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(body), nil
	}
}

func (conn SplunkConnection) httpPost(url string, data *url.Values) (string, error) {
	if response, err := conn.httpCall(url,"POST",data);err != nil {
		return "", err
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		return string(body), nil
	}
}

func (conn SplunkConnection) httpCall(url string, method string, data *url.Values) (*http.Response, error) {
        var payload io.Reader
        if data != nil {
          payload = bytes.NewBufferString(data.Encode())
        }

        request, err := http.NewRequest(method, url, payload)
        conn.addAuthHeader(request)
        response, err := conn.HttpClient.Do(request)

        if err != nil {
                return nil, err
        }
	return response, err
}

func (conn SplunkConnection) addAuthHeader(request *http.Request) {
        request.SetBasicAuth(conn.Username, conn.Password)
}


