# Gotenberg Go client

A simple Go client for interacting with a Gotenberg API.

## Install

```bash
$ go get -u github.com/thecodingmachine/gotenberg-go-client/v7
```

## Usage

```golang
import (
    "time"
    "net/http"

    "github.com/thecodingmachine/gotenberg-go-client/v7"
)

// create the client.
client := &gotenberg.Client{Hostname: "http://localhost:3000"}
// ... or use your own *http.Client.
httpClient := &http.Client{
    Timeout: time.Duration(5) * time.Second,
}
client := &gotenberg.Client{Hostname: "http://localhost:3000", HTTPClient: httpClient}

// prepare the files required for your conversion.

// from a path.
index, _ := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")
// ... or from a string.
index, _ := gotenberg.NewDocumentFromString("index.html", "<html>Foo</html>")
// ... or from bytes.
index, _ := gotenberg.NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))

header, _ := gotenberg.NewDocumentFromPath("header.html", "/path/to/file")
footer, _ := gotenberg.NewDocumentFromPath("footer.html", "/path/to/file")
style, _ := gotenberg.NewDocumentFromPath("style.css", "/path/to/file")
img, _ := gotenberg.NewDocumentFromPath("img.png", "/path/to/file")

req := gotenberg.NewHTMLRequest(index)
req.Header(header)
req.Footer(footer)
req.Assets(style, img)
req.PaperSize(gotenberg.A4)
req.Margins(gotenberg.NoMargins)
req.Scale(0.75)

// store method allows you to... store the resulting PDF in a particular destination.
client.Store(req, "path/you/want/the/pdf/to/be/stored.pdf")

// if you wish to redirect the response directly to the browser, you may also use:
resp, _ := client.Post(req)
```

For more complete usages, head to the [documentation](https://thecodingmachine.github.io/gotenberg).

## Badges

[![Travis CI](https://travis-ci.org/thecodingmachine/gotenberg-go-client.svg?branch=master)](https://travis-ci.org/thecodingmachine/gotenberg-go-client)
[![GoDoc](https://godoc.org/github.com/thecodingmachine/gotenberg-go-client?status.svg)](https://godoc.org/github.com/thecodingmachine/gotenberg-go-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/thecodingmachine/gotenberg-go-client)](https://goreportcard.com/report/thecodingmachine/gotenberg-go-client)