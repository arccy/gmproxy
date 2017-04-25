package gmproxy

import "encoding/json"

type commonProxy struct {
	Protocol
	IpPort    string
	TsChecked int
	Curl      string
	Type      string
}

type protoProxy struct {
	commonProxy
	OtherProtocols *json.RawMessage
}

type Proxy struct {
	commonProxy
	OtherProtocols []Protocol
}

type Protocol struct {
	Get            bool
	Post           bool
	Cookies        bool
	Referer        bool
	UserAgent      bool `json:"user-agent"`
	AnonymityLevel int
	SupportsHttps  bool
	Protocol       string
	Ip             string
	Port           string
	Country        string
	Websites       map[string]bool
}
