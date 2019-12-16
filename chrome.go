package gotenberg

import (
	"fmt"
	"strconv"
)

const (
	waitDelay                  string = "waitDelay"
	paperWidth                 string = "paperWidth"
	paperHeight                string = "paperHeight"
	marginTop                  string = "marginTop"
	marginBottom               string = "marginBottom"
	marginLeft                 string = "marginLeft"
	marginRight                string = "marginRight"
	landscapeChrome            string = "landscape"
	pageRanges                 string = "pageRanges"
	googleChromeRpccBufferSize string = "googleChromeRpccBufferSize"
)

// nolint: gochecknoglobals
var (
	// A3 paper size.
	A3 = [2]float64{11.7, 16.5}
	// A4 paper size.
	A4 = [2]float64{8.27, 11.7}
	// A5 paper size.
	A5 = [2]float64{5.8, 8.3}
	// A6 paper size.
	A6 = [2]float64{4.1, 5.8}
	// Letter paper size.
	Letter = [2]float64{8.5, 11}
	// Legal paper size.
	Legal = [2]float64{8.5, 14}
	// Tabloid paper size.
	Tabloid = [2]float64{11, 17}
)

// nolint: gochecknoglobals
var (
	// NoMargins removes margins.
	NoMargins = [4]float64{0, 0, 0, 0}
	// NormalMargins uses 1 inche margins.
	NormalMargins = [4]float64{1, 1, 1, 1}
	// LargeMargins uses 2 inche margins.
	LargeMargins = [4]float64{2, 2, 2, 2}
)

type chromeRequest struct {
	header Document
	footer Document

	*request
}

func newChromeRequest() *chromeRequest {
	return &chromeRequest{nil, nil, newRequest()}
}

// WaitDelay sets waitDelay form field.
func (req *chromeRequest) WaitDelay(delay float64) {
	req.values[waitDelay] = strconv.FormatFloat(delay, 'f', 2, 64)
}

// Header sets header form file.
func (req *chromeRequest) Header(header Document) {
	req.header = header
}

// Footer sets footer form file.
func (req *chromeRequest) Footer(footer Document) {
	req.footer = footer
}

// PaperSize sets paperWidth and paperHeight form fields.
func (req *chromeRequest) PaperSize(size [2]float64) {
	req.values[paperWidth] = fmt.Sprintf("%f", size[0])
	req.values[paperHeight] = fmt.Sprintf("%f", size[1])
}

// Margins sets marginTop, marginBottom,
// marginLeft and marginRight form fields.
func (req *chromeRequest) Margins(margins [4]float64) {
	req.values[marginTop] = fmt.Sprintf("%f", margins[0])
	req.values[marginBottom] = fmt.Sprintf("%f", margins[1])
	req.values[marginLeft] = fmt.Sprintf("%f", margins[2])
	req.values[marginRight] = fmt.Sprintf("%f", margins[3])
}

// Landscape sets landscape form field.
func (req *chromeRequest) Landscape(isLandscape bool) {
	req.values[landscapeChrome] = strconv.FormatBool(isLandscape)
}

// PageRanges sets pageRanges form field.
func (req *chromeRequest) PageRanges(ranges string) {
	req.values[pageRanges] = ranges
}

// GoogleChromeRpccBufferSize sets googleChromeRpccBufferSize form field.
func (req *chromeRequest) GoogleChromeRpccBufferSize(bufferSize int64) {
	req.values[googleChromeRpccBufferSize] = strconv.FormatInt(bufferSize, 10)
}
