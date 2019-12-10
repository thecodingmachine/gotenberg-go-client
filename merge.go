package gotenberg

// MergeRequest facilitates merging PDF
// with the Gotenberg API.
type MergeRequest struct {
	pdfs []Document

	*request
}

// NewMergeRequest create MergeRequest.
func NewMergeRequest(pdfs ...Document) *MergeRequest {
	return &MergeRequest{pdfs, newRequest()}
}

func (req *MergeRequest) postURL() string {
	return "/merge"
}

func (req *MergeRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(MergeRequest))
)
