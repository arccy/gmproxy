package gmproxy

import "net/http"

var DefaultClient = Client{
	Client: &http.Client{},
	Config: &RequestConfig{},
}

func GetProxy() (Proxy, error) {
	return DefaultClient.GetProxy()
}
