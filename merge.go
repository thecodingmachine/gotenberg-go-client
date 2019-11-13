package gotenberg

import (
	"fmt"
	"path/filepath"
)

// MergeRequest facilitates merging PDF
// with the Gotenberg API.
type MergeRequest struct {
	filePaths []string

	*request
}

// NewMergeRequest create MergeRequest.
func NewMergeRequest(fpaths ...string) (*MergeRequest, error) {
	for _, fpath := range fpaths {
		if !fileExists(fpath) {
			return nil, fmt.Errorf("%s: file does not exist", fpath)
		}
	}
	return &MergeRequest{fpaths, newRequest()}, nil
}

func (req *MergeRequest) postURL() string {
	return "/merge"
}

func (req *MergeRequest) formFiles() map[string]string {
	files := make(map[string]string)
	for _, fpath := range req.filePaths {
		files[filepath.Base(fpath)] = fpath
	}
	return files
}

func (req *MergeRequest) formData() map[string]string {
	files := make(map[string]string)

	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(MergeRequest))
)
