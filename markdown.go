package gotenberg

import (
	"fmt"
	"path/filepath"
)

// MarkdownRequest facilitates Markdown conversion
// with the Gotenberg API.
type MarkdownRequest struct {
	indexFilePath     string
	markdownFilePaths []string
	assetFilePaths    []string

	*chromeRequest
}

// NewMarkdownRequest create MarkdownRequest.
func NewMarkdownRequest(indexFilePath string, markdownFilePaths ...string) (*MarkdownRequest, error) {
	if !fileExists(indexFilePath) {
		return nil, fmt.Errorf("%s: index file does not exist", indexFilePath)
	}
	for _, fpath := range markdownFilePaths {
		if !fileExists(fpath) {
			return nil, fmt.Errorf("%s: markdown file does not exist", fpath)
		}
	}
	return &MarkdownRequest{indexFilePath, markdownFilePaths, nil, newChromeRequest()}, nil
}

// Assets sets assets form files.
func (req *MarkdownRequest) Assets(fpaths ...string) error {
	for _, fpath := range fpaths {
		if !fileExists(fpath) {
			return fmt.Errorf("%s: file does not exist", fpath)
		}
	}
	req.assetFilePaths = fpaths
	return nil
}

func (req *MarkdownRequest) postURL() string {
	return "/convert/markdown"
}

func (req *MarkdownRequest) formFiles() map[string]string {
	files := make(map[string]string)
	files["index.html"] = req.indexFilePath
	files["header.html"] = req.headerFilePath
	files["footer.html"] = req.footerFilePath
	for _, fpath := range req.markdownFilePaths {
		files[filepath.Base(fpath)] = fpath
	}
	for _, fpath := range req.assetFilePaths {
		files[filepath.Base(fpath)] = fpath
	}
	return files
}

func (req *MarkdownRequest) formData() map[string]string {
	files := make(map[string]string)
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(MarkdownRequest))
)
