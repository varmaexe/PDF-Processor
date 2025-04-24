package converter

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// SVGToPNG converts raw SVG data to a PNG file using rsvg-convert.
// Returns error if conversion fails at any step.
func SVGToPNG(svgData []byte, outputPath string) error {
	// Validate input SVG data
	if len(svgData) == 0 {
		return fmt.Errorf("empty SVG data")
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create temporary SVG file for conversion
	tmpSVGFile, err := os.CreateTemp("", "temp-*.svg")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer os.Remove(tmpSVGFile.Name()) // Clean up temp file
	defer tmpSVGFile.Close()

	// Write SVG data to temp file
	if _, err := tmpSVGFile.Write(svgData); err != nil {
		return fmt.Errorf("error writing SVG: %w", err)
	}

	// Ensure data is flushed to disk before conversion
	if err := tmpSVGFile.Sync(); err != nil {
		return fmt.Errorf("error syncing temp file: %w", err)
	}

	// Verify rsvg-convert is available in PATH
	if _, err := exec.LookPath("rsvg-convert"); err != nil {
		return fmt.Errorf("rsvg-convert not found in PATH: %w", err)
	}

	// Execute conversion command
	cmd := exec.Command("rsvg-convert", "-f", "png", "-o", outputPath, tmpSVGFile.Name())
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("conversion error: %w, output: %s", err, string(output))
	}

	return nil
}

func AddBackgroundToSVG(svg []byte, hexColor string) []byte {
	svgStr := string(svg)
	if idx := bytes.Index(svg, []byte("<svg")); idx != -1 {
		if endIdx := bytes.Index(svg[idx:], []byte(">")); endIdx != -1 {
			insertPos := idx + endIdx + 1
			injection := fmt.Sprintf(`<rect width="100%%" height="100%%" fill="%s"/>`, hexColor)
			modified := svgStr[:insertPos] + injection + svgStr[insertPos:]
			return []byte(modified)
		}
	}
	return svg
}
