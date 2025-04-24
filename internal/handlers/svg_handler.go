package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"pdfprocessor/internal/converter"

	"github.com/gin-gonic/gin"
)

func ConvertSVGToPDF(c *gin.Context) {
	svgData, err := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil || len(svgData) == 0 {
		msg := "Invalid or empty SVG data"
		fmt.Println("[ERROR]", msg)
		c.String(http.StatusBadRequest, msg)
		return
	}

	bgColor := c.DefaultQuery("bg", "#080634")
	svgData = converter.AddBackgroundToSVG(svgData, bgColor)

	outputFilename := c.DefaultQuery("filename", "converted.pdf")

	tmpDir, err := os.MkdirTemp("", "pdfgen-*")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Temp dir error: %v", err))
		return
	}
	defer os.RemoveAll(tmpDir)

	pngPath := filepath.Join(tmpDir, "image.png")
	if err = converter.SVGToPNG(svgData, pngPath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("SVG to PNG error: %v", err))
		return
	}

	pngFile, err := os.Open(pngPath)
	defer pngFile.Close()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Open PNG error: %v", err))
		return
	}

	pdfBuf, err := converter.GeneratePDF(pngFile)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("PDF generation error: %v", err))
		return
	}

	c.DataFromReader(http.StatusOK, int64(pdfBuf.Len()), "application/pdf", bytes.NewReader(pdfBuf.Bytes()), map[string]string{
		"Content-Disposition": fmt.Sprintf(`attachment; filename="%s"`, outputFilename),
	})
}
