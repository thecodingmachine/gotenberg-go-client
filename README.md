# Gotenberg Go client

A simple Go client for interacting with a Gotenberg API.

## Install

```bash
$ go get -u github.com/thecodingmachine/gotenberg-go-client/v4
```

## Usage

```golang
import "github.com/thecodingmachine/gotenberg-go-client/v4"

func main() {
    // HTML conversion example.
    c := &gotenberg.Client{Hostname: "http://localhost:3000"}
    req, _ := gotenberg.NewHTMLRequest("index.html")
    req.SetHeader("header.html")
    req.SetFooter("footer.html")
    req.SetAssets(
        "font.woff",
        "img.gif",
        "style.css",
    )
    req.SetPaperSize(gotenberg.A4)
    req.SetMargins(gotenberg.NormalMargins)
    req.SetLandscape(false)
    dest := "foo.pdf"
    c.Store(req, dest)
}
```

For more complete usages, head to the [documentation](https://thecodingmachine.github.io/gotenberg).


## Badges

[![Travis CI](https://travis-ci.org/thecodingmachine/gotenberg-go-client.svg?branch=master)](https://travis-ci.org/thecodingmachine/gotenberg-go-client)
[![GoDoc](https://godoc.org/github.com/thecodingmachine/gotenberg-go-client?status.svg)](https://godoc.org/github.com/thecodingmachine/gotenberg-go-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/thecodingmachine/gotenberg-go-client)](https://goreportcard.com/report/thecodingmachine/gotenberg-go-client)