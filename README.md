# go-gimmeproxy
go-gimmeproxy is a Go client library for accessing the [GimmeProxy API](http://256cats.com/gimmeproxy-com-free-rotating-proxy-api/)

## Installation
To install the library, simple:
```
go get github.com/arccy/gmproxy
```

## Usage
### Default Client
For situations where you just need a proxy, a package level wrapper function is provided:
```
gmproxy.GetProxy()
```

### Optional Parameters
To customize the request, create your own `Client`:
```
var CustomClient = gmproxy.Client {
        Client: &http.Client{},
        Config: &gmproxy.RequestConfig{},
}
```

The `http.Client` is used to make the request

`RequestConfig` specifies the query parameters

### Results
__Or why this package exists__

The results are parsed into structs! see `response.go`


## License
MIT, see LICENCSE file
