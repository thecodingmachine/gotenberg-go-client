package gotenberg

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thecodingmachine/gotenberg-go-client/v6/test"
)

func TestHTML(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.WaitDelay(1)
	err = req.Header(test.HTMLTestFilePath(t, "header.html"))
	require.Nil(t, err)
	err = req.Footer(test.HTMLTestFilePath(t, "footer.html"))
	require.Nil(t, err)
	err = req.Assets(
		test.HTMLTestFilePath(t, "font.woff"),
		test.HTMLTestFilePath(t, "img.gif"),
		test.HTMLTestFilePath(t, "style.css"),
	)
	require.Nil(t, err)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Landscape(false)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLWithoutHeaderFooter(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Landscape(false)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}
