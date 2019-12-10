package gotenberg

import "fmt"

const (
	remoteURL                  string = "remoteURL"
	remoteURLBaseHTTPHeaderKey string = "Gotenberg-Remoteurl-"
)

// URLRequest facilitates remote URL conversion
// with the Gotenberg API.
type URLRequest struct {
	*chromeRequest
}

// NewURLRequest create URLRequest.
func NewURLRequest(url string) *URLRequest {
	req := &URLRequest{newChromeRequest()}
	req.values[remoteURL] = url
	return req
}

func (url *URLRequest) postURL() string {
	return "/convert/url"
}

// AddRemoteURLHTTPHeader add a remote URL custom HTTP header.
func (url *URLRequest) AddRemoteURLHTTPHeader(key, value string) {
	key = fmt.Sprintf("%s%s", remoteURLBaseHTTPHeaderKey, key)
	url.httpHeaders[key] = value
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(URLRequest))
)
