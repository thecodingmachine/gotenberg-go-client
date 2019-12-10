package gotenberg

// MarkdownRequest facilitates Markdown conversion
// with the Gotenberg API.
type MarkdownRequest struct {
	index     Document
	markdowns []Document
	assets    []Document

	*chromeRequest
}

// NewMarkdownRequest create MarkdownRequest.
func NewMarkdownRequest(index Document, markdowns ...Document) *MarkdownRequest {
	return &MarkdownRequest{index, markdowns, []Document{}, newChromeRequest()}
}

// Assets sets assets form files.
func (req *MarkdownRequest) Assets(assets ...Document) {
	req.assets = assets
}

func (req *MarkdownRequest) postURL() string {
	return "/convert/markdown"
}

func (req *MarkdownRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	files["index.html"] = req.index
	for _, markdown := range req.markdowns {
		files[markdown.Filename()] = markdown
	}
	for _, asset := range req.assets {
		files[asset.Filename()] = asset
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(MarkdownRequest))
)
