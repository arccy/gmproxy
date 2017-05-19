package gmproxy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Api endpoint
const Api = "http://gimmeproxy.com/api/getProxy"

// Wrap http.Client and RequestConfig together
type Client struct {
	Client *http.Client
	Config *RequestConfig
}

// request proxy and parse results
func (c *Client) GetProxy() (Proxy, error) {
	req, err := http.NewRequest("GET", Api, nil)
	if err != nil {
		return Proxy{}, err
	}
	req.URL.RawQuery = c.Config.ToUrl()

	resp, err := c.Client.Do(req)
	if err != nil {
		return Proxy{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Proxy{}, err
		}
		return Proxy{}, errors.New(string(t))
	}

	newProxy, err := decodeJson(resp)
	if err != nil {
		return Proxy{}, err
	}
	return newProxy, nil
}

// Marshal json into struxt
func decodeJson(resp *http.Response) (Proxy, error) {
	var rawProxy protoProxy
	err := json.NewDecoder(resp.Body).Decode(&rawProxy)
	if err != nil {
		return Proxy{}, err
	}

	var newProxy Proxy
	newProxy.commonProxy = rawProxy.commonProxy

	// Deal with single/array of OtherProtocols
	var toTest interface{}
	json.Unmarshal(*rawProxy.OtherProtocols, &toTest)
	switch toTest.(type) {
	case map[string]interface{}:
		var single Protocol
		err = json.Unmarshal(*rawProxy.OtherProtocols, &single)
		if err != nil {
			return Proxy{}, err
		}

		newProxy.OtherProtocols = []Protocol{single}
	case []map[string]interface{}:
		err = json.Unmarshal(*rawProxy.OtherProtocols, &newProxy.OtherProtocols)
		if err != nil {
			return Proxy{}, err
		}

	}

	return newProxy, nil
}
