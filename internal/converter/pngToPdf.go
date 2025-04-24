package converter

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// GeneratePDF takes an image (PNG) as an io.Reader and returns a PDF buffer.
func GeneratePDF(image io.Reader) (*bytes.Buffer, error) {
	var pdfBuf bytes.Buffer

	fmt.Println("[INFO] Importing PNG into PDF...")
	err := api.ImportImages(nil, &pdfBuf, []io.Reader{image}, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error importing image to PDF: %w", err)
	}

	fmt.Println("[SUCCESS] PDF creation successful")
	return &pdfBuf, nil
}
