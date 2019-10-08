package gotenberg

const remoteURL string = "remoteURL"

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

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(URLRequest))
)
