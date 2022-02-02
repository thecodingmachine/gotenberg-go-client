package gotenberg

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"log"
)

const (
	resultFilename              string = "resultFilename"
	waitTimeout                 string = "waitTimeout"
	webhookURL                  string = "webhookURL"
	webhookURLTimeout           string = "webhookURLTimeout"
	webhookURLBaseHTTPHeaderKey string = "Gotenberg-Webhookurl-"
)

// Client facilitates interacting with
// the Gotenberg API.
type Client struct {
	Hostname   string
	HTTPClient *http.Client
}

// Request is a type for sending
// form values and form files to
// the Gotenberg API.
type Request interface {
	postURL() string
	customHTTPHeaders() map[string]string
	formValues() map[string]string
	formFiles() map[string]Document
}

type request struct {
	httpHeaders map[string]string
	values      map[string]string
}

func newRequest() *request {
	return &request{
		httpHeaders: make(map[string]string),
		values:      make(map[string]string),
	}
}

// ResultFilename sets resultFilename form field.
func (req *request) ResultFilename(filename string) {
	req.values[resultFilename] = filename
}

// WaitTimeout sets waitTimeout form field.
func (req *request) WaitTimeout(timeout float64) {
	req.values[waitTimeout] = strconv.FormatFloat(timeout, 'f', 2, 64)
}

// WebhookURL sets webhookURL form field.
func (req *request) WebhookURL(url string) {
	req.values[webhookURL] = url
}

// WebhookURLTimeout sets webhookURLTimeout form field.
func (req *request) WebhookURLTimeout(timeout float64) {
	req.values[webhookURLTimeout] = strconv.FormatFloat(timeout, 'f', 2, 64)
}

// AddWebhookURLHTTPHeader add a webhook custom HTTP header.
func (req *request) AddWebhookURLHTTPHeader(key, value string) {
	key = fmt.Sprintf("%s%s", webhookURLBaseHTTPHeaderKey, key)
	req.httpHeaders[key] = value
}

func (req *request) customHTTPHeaders() map[string]string {
	return req.httpHeaders
}

func (req *request) formValues() map[string]string {
	return req.values
}

// Post sends a request to the Gotenberg API
// and returns the response.
func (c *Client) Post(req Request) (*http.Response, error) {
	return c.PostContext(context.Background(), req)
}

// PostContext sends a request to the Gotenberg API
// and returns the response.
// The created HTTP request can be canceled by the passed context.
func (c *Client) PostContext(ctx context.Context, req Request) (*http.Response, error) {
	body, contentType, err := multipartForm(req)
	if err != nil {
		return nil, err
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{}
	}
	URL := fmt.Sprintf("%s%s", c.Hostname, req.postURL())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, URL, body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", contentType)
	for key, value := range req.customHTTPHeaders() {
		httpReq.Header.Set(key, value)
	}
	resp, err := c.HTTPClient.Do(httpReq) /* #nosec */
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Store creates the resulting PDF to given destination.
func (c *Client) Store(req Request, dest string) error {
	return c.StoreContext(context.Background(), req, dest)
}

// StoreContext creates the resulting PDF to given destination.
// The created HTTP request can be canceled by the passed context.
func (c *Client) StoreContext(ctx context.Context, req Request, dest string) error {
	if hasWebhook(req) {
		return errors.New("cannot use Store method with a webhook")
	}
	resp, err := c.PostContext(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("failed to generate the result PDF - http status code: [%d]", resp.StatusCode) // more details here now!
	}
	
        if resp.StatusCode < http.StatusOK {
		return log.Printf("http status code: [%d]") // maybe the gotenburg server is still processing the request etc - this could be passed through as a bool maybe.
	}
	
	return writeNewFile(dest, resp.Body)
}

func hasWebhook(req Request) bool {
	webhookURL, ok := req.formValues()[webhookURL]
	if !ok {
		return false
	}
	return webhookURL != ""
}

func writeNewFile(fpath string, in io.Reader) error {
	if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
		return fmt.Errorf("%s: making directory for file: %v", fpath, err)
	}
	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", fpath, err)
	}
	defer out.Close() // nolint: errcheck
	err = out.Chmod(0644)
	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", fpath, err)
	}
	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("%s: writing file: %v", fpath, err)
	}
	return nil
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func multipartForm(req Request) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close() // nolint: errcheck
	for filename, document := range req.formFiles() {
		in, err := document.Reader()
		if err != nil {
			return nil, "", fmt.Errorf("%s: creating reader: %v", filename, err)
		}
		defer in.Close() // nolint: errcheck
		part, err := writer.CreateFormFile("files", filename)
		if err != nil {
			return nil, "", fmt.Errorf("%s: creating form file: %v", filename, err)
		}
		_, err = io.Copy(part, in)
		if err != nil {
			return nil, "", fmt.Errorf("%s: copying data: %v", filename, err)
		}
	}
	for name, value := range req.formValues() {
		if err := writer.WriteField(name, value); err != nil {
			return nil, "", fmt.Errorf("%s: writing form field: %v", name, err)
		}
	}
	return body, writer.FormDataContentType(), nil
}
