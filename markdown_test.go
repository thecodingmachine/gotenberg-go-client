package gotenberg

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thecodingmachine/gotenberg-go-client/v6/test"
)

func TestMarkdown(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewMarkdownRequest(
		test.MarkdownTestFilePath(t, "index.html"),
		test.MarkdownTestFilePath(t, "paragraph1.md"),
		test.MarkdownTestFilePath(t, "paragraph2.md"),
		test.MarkdownTestFilePath(t, "paragraph3.md"),
	)
	require.Nil(t, err)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.WaitDelay(1)
	err = req.Header(test.MarkdownTestFilePath(t, "header.html"))
	require.Nil(t, err)
	err = req.Footer(test.MarkdownTestFilePath(t, "footer.html"))
	require.Nil(t, err)
	err = req.Assets(
		test.MarkdownTestFilePath(t, "font.woff"),
		test.MarkdownTestFilePath(t, "img.gif"),
		test.MarkdownTestFilePath(t, "style.css"),
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

func TestMarkdownWithoutHeaderFooter(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewMarkdownRequest(
		test.MarkdownTestFilePath(t, "index.html"),
		test.MarkdownTestFilePath(t, "paragraph1.md"),
		test.MarkdownTestFilePath(t, "paragraph2.md"),
		test.MarkdownTestFilePath(t, "paragraph3.md"),
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

func TestMarkdownWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewMarkdownRequest(
		test.MarkdownTestFilePath(t, "index.html"),
		test.MarkdownTestFilePath(t, "paragraph1.md"),
		test.MarkdownTestFilePath(t, "paragraph2.md"),
		test.MarkdownTestFilePath(t, "paragraph3.md"),
	)
	require.Nil(t, err)
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
