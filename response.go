package gmproxy

import "encoding/json"

// A proxy from the Api
type Proxy struct {
	commonProxy
	OtherProtocols []Protocol
}

// Attributes common to the primary proxy and other Protocols
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

// Split out because DRY between protoProxy and Proxy
type commonProxy struct {
	Protocol
	IpPort    string
	TsChecked int
	Curl      string
	Type      string
}

// necessary because OtherProtocols could a single item or a list
type protoProxy struct {
	commonProxy
	OtherProtocols *json.RawMessage
}
