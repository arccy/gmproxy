package gmproxy

import (
	"errors"
	"net/url"
	"strconv"
)

// Valid query params, 2017/05
const (
	FilterApi_key        = "api_key"
	FilterGet            = "get"
	FilterPost           = "post"
	FilterCookies        = "cookies"
	FilterReferer        = "referer"
	FilterUserAgent      = "user-agent"
	FilterSupportsHttps  = "supportsHttps"
	FilterAnonymityLevel = "anonymityLevel"
	FilterProtocol       = "protocol"
	FilterPort           = "port"
	FilterCountry        = "country"
	FilterMaxCheckPeriod = "maxCheckPeriod"
	FilterWebsites       = "websites"
	FilterMinSpeed       = "minSpeed"
	FilterNotCountry     = "notCountry"
)

// Key-Value pairs of request params
type RequestConfig struct {
	v url.Values
}

// Type switch and insert into RequestConfig
func (r *RequestConfig) Add(param string, val interface{}) error {
	switch val.(type) {
	default:
		return errors.New("Unkown type when adding value to RequestParams")
	case string:
		r.v.Add(param, val.(string))
	case int:
		r.v.Add(param, strconv.Itoa(val.(int)))
	case float64:
		r.v.Add(param, strconv.FormatFloat(val.(float64), 'f', -1, 64))
	case bool:
		if val.(bool) == true {
			r.v.Add(param, "1")
		} else {
			r.v.Add(param, "0")
		}
	}
	return nil
}

// converts url.Values to query string
func (r *RequestConfig) ToUrl() string {
	return r.v.Encode()
}
