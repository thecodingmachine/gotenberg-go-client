package gotenberg

import (
	"fmt"
	"path/filepath"
)

// HTMLRequest facilitates HTML conversion
// with the Gotenberg API.
type HTMLRequest struct {
	indexFilePath  string
	assetFilePaths []string

	*chromeRequest
}

// NewHTMLRequest create HTMLRequest.
func NewHTMLRequest(indexFilePath string) (*HTMLRequest, error) {
	if !fileExists(indexFilePath) {
		return nil, fmt.Errorf("%s: index file does not exist", indexFilePath)
	}
	return &HTMLRequest{indexFilePath, nil, newChromeRequest()}, nil
}

// Assets sets assets form files.
func (req *HTMLRequest) Assets(fpaths ...string) error {
	for _, fpath := range fpaths {
		if !fileExists(fpath) {
			return fmt.Errorf("%s: file does not exist", fpath)
		}
	}
	req.assetFilePaths = fpaths
	return nil
}

func (req *HTMLRequest) postURL() string {
	return "/convert/html"
}

func (req *HTMLRequest) formFiles() map[string]string {
	files := make(map[string]string)
	files["index.html"] = req.indexFilePath
	files["header.html"] = req.headerFilePath
	files["footer.html"] = req.footerFilePath
	for _, fpath := range req.assetFilePaths {
		files[filepath.Base(fpath)] = fpath
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(HTMLRequest))
)
