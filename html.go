package gotenberg

import (
	"fmt"
	"path/filepath"
)

const (
	//PathOption is used to denote that index.html is from path
	PathOption string = "path"
	//RawHTMLOption is used to denote that index.html is a raw string
	RawHTMLOption string = "raw"
)

// HTMLRequest facilitates HTML conversion
// with the Gotenberg API.
type HTMLRequest struct {
	indexFilePath  string
	assetFilePaths []string
	*chromeRequest
	indexFileData string
}

// NewHTMLRequest create HTMLRequest.
func NewHTMLRequest(fileInfo string, filePass string) (*HTMLRequest, error) {
	if filePass == "path" {
		if !fileExists(fileInfo) {
			return nil, fmt.Errorf("%s: index file does not exist", fileInfo)
		}
		return &HTMLRequest{fileInfo, nil, newChromeRequest(), ""}, nil
	}
	if len(fileInfo) == 0 {
		return nil, fmt.Errorf("index data is empty")
	}
	return &HTMLRequest{"", nil, newChromeRequest(), fileInfo}, nil
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

func (req *HTMLRequest) formData() map[string]string {
	files := make(map[string]string)
	files["index.html"] = req.indexFileData
	files["header.html"] = req.headerData
	files["footer.html"] = req.footerData

	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(HTMLRequest))
)
