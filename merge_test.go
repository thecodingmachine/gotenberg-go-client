package gotenberg

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thecodingmachine/gotenberg-go-client/v4/test"
)

func TestMerge(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req, err := NewMergeRequest(
		test.PDFTestFilePath(t, "gotenberg.pdf"),
		test.PDFTestFilePath(t, "gotenberg.pdf"),
	)
	require.Nil(t, err)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}
