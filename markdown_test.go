package gotenberg

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thecodingmachine/gotenberg-go-client/v7/test"
)

func TestMarkdown(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.Nil(t, err)
	markdown1, err := NewDocumentFromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.Nil(t, err)
	markdown2, err := NewDocumentFromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.Nil(t, err)
	markdown3, err := NewDocumentFromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.Nil(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestMarkdownComplete(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.Nil(t, err)
	markdown1, err := NewDocumentFromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.Nil(t, err)
	markdown2, err := NewDocumentFromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.Nil(t, err)
	markdown3, err := NewDocumentFromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.Nil(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	header, err := NewDocumentFromPath("header.html", test.MarkdownTestFilePath(t, "header.html"))
	require.Nil(t, err)
	req.Header(header)
	footer, err := NewDocumentFromPath("footer.html", test.MarkdownTestFilePath(t, "footer.html"))
	require.Nil(t, err)
	req.Footer(footer)
	font, err := NewDocumentFromPath("font.woff", test.MarkdownTestFilePath(t, "font.woff"))
	require.Nil(t, err)
	img, err := NewDocumentFromPath("img.gif", test.MarkdownTestFilePath(t, "img.gif"))
	require.Nil(t, err)
	style, err := NewDocumentFromPath("style.css", test.MarkdownTestFilePath(t, "style.css"))
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

func TestMarkdownPageRanges(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.Nil(t, err)
	markdown1, err := NewDocumentFromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.Nil(t, err)
	markdown2, err := NewDocumentFromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.Nil(t, err)
	markdown3, err := NewDocumentFromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.Nil(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	req.PageRanges("1-1")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestMarkdownWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.Nil(t, err)
	markdown1, err := NewDocumentFromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.Nil(t, err)
	markdown2, err := NewDocumentFromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.Nil(t, err)
	markdown3, err := NewDocumentFromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.Nil(t, err)
	req := NewMarkdownRequest(index, markdown1, markdown2, markdown3)
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
