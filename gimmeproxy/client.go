package gmproxy

import (
	"encoding/json"
	"net/http"
)

const Api = "http://gimmeproxy.com/api/getProxy"

type Client struct {
	Client *http.Client
	Config *RequestConfig
}

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

	newProxy, err := decodeJson(resp)
	if err != nil {
		return Proxy{}, err
	}
	return newProxy, nil
}

func decodeJson(resp *http.Response) (Proxy, error) {
	var rawProxy protoProxy
	err := json.NewDecoder(resp.Body).Decode(&rawProxy)
	if err != nil {
		return Proxy{}, err
	}

	var newProxy Proxy
	newProxy.commonProxy = rawProxy.commonProxy

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
