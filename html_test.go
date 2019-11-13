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
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, "index.html"), PathOption)
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
	req.GoogleChromeRpccBufferSize(1048576)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLRaw(t *testing.T) {

	indexString := "<html><body><p>go filler text</p></body></html>"
	newHeader := "<html><body><div>headtext</div></body></html>"

	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, indexString), RawHTMLOption)
	require.Nil(t, err)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.WaitDelay(1)
	err = req.HeaderRaw(test.HTMLTestFilePath(t, newHeader))
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
	req.GoogleChromeRpccBufferSize(1048576)
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
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, "index.html"), PathOption)
	require.Nil(t, err)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Landscape(false)
	req.GoogleChromeRpccBufferSize(1048576)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLWithGetByte(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, "index.html"), PathOption)
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
	req.GoogleChromeRpccBufferSize(1048576)
	pdfBytes, err := c.GetBytes(req)
	assert.Nil(t, err)
	assert.NotEmpty(t, pdfBytes)
	assert.True(t, (len(pdfBytes) > 120800))
	assert.Contains(t, string(pdfBytes), "%PDF-1.4")
	assert.Contains(t, string(pdfBytes), "/Width 287")
	assert.Contains(t, string(pdfBytes), "/Height 320")
	assert.Nil(t, err)
}

func TestHTMLStoreAndGetByte(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, "index.html"), PathOption)
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
	req.GoogleChromeRpccBufferSize(1048576)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
	pdfBytes, err := c.GetBytes(req)
	assert.Nil(t, err)
	assert.NotEmpty(t, pdfBytes)
	assert.Nil(t, err)
}

func TestHTMLWithGetByteWithoutHeaderFooter(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewHTMLRequest(test.HTMLTestFilePath(t, "index.html"), PathOption)
	require.Nil(t, err)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Landscape(false)
	req.GoogleChromeRpccBufferSize(1048576)
	pdfBytes, err := c.GetBytes(req)
	assert.Nil(t, err)
	assert.NotEmpty(t, pdfBytes)
	assert.Nil(t, err)
}
