package gmproxy

import "net/http"

// Client with no reuest params using default http.Client
var DefaultClient = Client{
	Client: &http.Client{},
	Config: &RequestConfig{},
}

// Wrapper function using DefaultClient
func GetProxy() (Proxy, error) {
	return DefaultClient.GetProxy()
}
