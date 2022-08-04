package gotenberg

// HTMLRequest facilitates HTML conversion
// with the Gotenberg API.
type HTMLRequest struct {
	index  Document
	assets []Document

	*chromeRequest
}

// NewHTMLRequest create HTMLRequest.
func NewHTMLRequest(index Document) *HTMLRequest {
	return &HTMLRequest{index, []Document{}, newChromeRequest()}
}

// Assets sets assets form files.
func (req *HTMLRequest) Assets(assets ...Document) {
	req.assets = assets
}

func (req *HTMLRequest) postURL() string {
	return "/forms/chromium/convert/html"
}

func (req *HTMLRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	files["index.html"] = req.index
	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}
	for _, asset := range req.assets {
		files[asset.Filename()] = asset
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(HTMLRequest))
)
