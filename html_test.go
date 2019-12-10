package gotenberg

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thecodingmachine/gotenberg-go-client/v7/test"
)

func TestHTML(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLFromString(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromString("index.html", "<html>Foo</html>")
	req := NewHTMLRequest(index)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLFromBytes(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))
	req := NewHTMLRequest(index)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLComplete(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	header, err := NewDocumentFromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.Nil(t, err)
	req.Header(header)
	footer, err := NewDocumentFromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.Nil(t, err)
	req.Footer(footer)
	font, err := NewDocumentFromPath("font.woff", test.HTMLTestFilePath(t, "font.woff"))
	require.Nil(t, err)
	img, err := NewDocumentFromPath("img.gif", test.HTMLTestFilePath(t, "img.gif"))
	require.Nil(t, err)
	style, err := NewDocumentFromPath("style.css", test.HTMLTestFilePath(t, "style.css"))
	req.Assets(font, img, style)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.WaitDelay(1)
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

func TestHTMLWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
