// Package test contains useful functions used across tests.
package test

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

// Rand returns a random string.
func Rand() (string, error) {
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", fmt.Errorf("creating random string: %v", err)
	}
	return hex.EncodeToString(randBytes), nil
}

// HTMLTestFilePath returns the absolute file path
// of a file in "html" folder in test/testdata.
func HTMLTestFilePath(t *testing.T, filename string) string {
	return abs(t, "html", filename)
}

// URLTestFilePath returns the absolute file path
// of a file in "url" folder in test/testdata.
func URLTestFilePath(t *testing.T, filename string) string {
	return abs(t, "url", filename)
}

// MarkdownTestFilePath returns the absolute file path
// of a file in "markdown" folder in test/testdata.
func MarkdownTestFilePath(t *testing.T, filename string) string {
	return abs(t, "markdown", filename)
}

// OfficeTestFilePath returns the absolute file path
// of a file in "office" folder in test/testdata.
func OfficeTestFilePath(t *testing.T, filename string) string {
	return abs(t, "office", filename)
}

// PDFTestFilePath returns the absolute file path
// of a file in "pdf" folder in test/testdata.
func PDFTestFilePath(t *testing.T, filename string) string {
	return abs(t, "pdf", filename)
}

func abs(t *testing.T, kind, filename string) string {
	_, gofilename, _, ok := runtime.Caller(0)
	require.Equal(t, ok, true, "got no caller information")
	if filename == "" {
		path, err := filepath.Abs(fmt.Sprintf("%s/testdata/%s", path.Dir(gofilename), kind))
		require.Nil(t, err, `getting the absolute path of "%s"`, kind)
		return path
	}
	path, err := filepath.Abs(fmt.Sprintf("%s/testdata/%s/%s", path.Dir(gofilename), kind, filename))
	require.Nil(t, err, `getting the absolute path of "%s"`, filename)
	return path
}
