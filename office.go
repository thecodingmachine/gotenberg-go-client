package gotenberg

import (
	"fmt"
	"path/filepath"
	"strconv"
)

const landscapeOffice string = "landscape"

// OfficeRequest facilitates Office documents
// conversion with the Gotenberg API.
type OfficeRequest struct {
	filePaths []string

	*request
}

// NewOfficeRequest create OfficeRequest.
func NewOfficeRequest(fpaths ...string) (*OfficeRequest, error) {
	for _, fpath := range fpaths {
		if !fileExists(fpath) {
			return nil, fmt.Errorf("%s: file does not exist", fpath)
		}
	}
	return &OfficeRequest{fpaths, newRequest()}, nil
}

// Landscape sets landscape form field.
func (req *OfficeRequest) Landscape(isLandscape bool) {
	req.values[landscapeOffice] = strconv.FormatBool(isLandscape)
}

func (req *OfficeRequest) postURL() string {
	return "/convert/office"
}

func (req *OfficeRequest) formFiles() map[string]string {
	files := make(map[string]string)
	for _, fpath := range req.filePaths {
		files[filepath.Base(fpath)] = fpath
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(OfficeRequest))
)
