
# SVG to PDF Conversion Service

This project is a backend API that accepts SVG files and converts them to PDF format. It utilizes Go (Golang) with the `Gin` framework for routing and the `pdfcpu` library for PDF generation. The service uses `rsvg-convert` to convert SVG to PNG as an intermediate step, and then creates a PDF from the PNG image.

---

## Table of Contents

- [Project Overview](#project-overview)
- [Requirements](#requirements)
- [Installation](#installation)
- [Running the Application](#running-the-application)
  - [Using Docker](#using-docker)
  - [Without Docker](#without-docker)
- [API Endpoints](#api-endpoints)
- [Example Scenarios](#example-scenarios)
- [License](#license)

---

## Project Overview

This service provides a single API endpoint that accepts an SVG file, converts it to a PNG image, and then creates a PDF document. The PDF is then returned as a downloadable file.

---

## Requirements

- **Go 1.24+**
- **Docker** (for running the application in containers)
- **rsvg-convert** (used for converting SVG files to PNG)
- **pdfcpu** (used to create a PDF from the PNG)

If you choose to run the service locally, you will need to install `rsvg-convert` and `pdfcpu` dependencies.

---

## Installation

### 1. Clone the repository:

```bash
git clone https://github.com/yourusername/svg-to-pdf.git
cd svg-to-pdf
```

### 2. Install Go dependencies:

If you have Go installed locally, you can use `go mod` to download the necessary dependencies.

```bash
go mod tidy
```

### 3. Install `rsvg-convert` (if running without Docker):

On **Ubuntu/Debian**, you can install the required libraries:

```bash
sudo apt-get update
sudo apt-get install librsvg2-bin
```

---

## Running the Application

### Using Docker (Recommended)

The project is dockerized, and you can run it easily using Docker. Follow these steps:

#### 1. Build the Docker Image

```bash
docker build -t svg-to-pdf .
```

#### 2. Run the Docker Container

```bash
docker run -p 8080:8080 svg-to-pdf
```

This will start the application on port `8080`. You can now interact with the API.

### Without Docker

If you prefer to run the application locally without Docker:

#### 1. Run the Go application:

```bash
go run main.go
```

This will start the application on port `8080`.

---

## API Endpoints

### POST `/svg-to-pdf`

This endpoint accepts a POST request with the SVG file in the body and returns the generated PDF.

#### Request Body:

- **Content-Type:** `image/svg+xml`
- The body should contain the raw SVG data.

#### Response:

- **Content-Type:** `application/pdf`
- **Content-Disposition:** `attachment; filename="converted.pdf"`

The response will be a PDF file containing the converted content.

#### Example cURL Request:

```bash
curl -X POST http://localhost:8080/svg-to-pdf -H "Content-Type: image/svg+xml" --data-binary "@example.svg" --output result.pdf
```

---

## Example Scenarios

### Scenario 1: Convert an SVG to PDF via cURL

1. Save your SVG file as `example.svg`.
2. Run the following command:

```bash
curl -X POST http://localhost:8080/svg-to-pdf -H "Content-Type: image/svg+xml" --data-binary "@example.svg" --output result.pdf
```

3. The server will return the converted PDF as `result.pdf`.

### Scenario 2: Frontend Upload and Convert SVG to PDF

1. Open the `index.html` in your browser.
2. Upload an SVG file using the file input.
3. Click on **Convert**.
4. Once the conversion is complete, the PDF will be downloaded automatically.

### Scenario 3: Running with Docker on Production

1. Build and run the Docker container:

```bash
docker build -t svg-to-pdf .
docker run -p 8080:8080 svg-to-pdf
```

2. The server will be running on `http://localhost:8080`, and you can use either the frontend or API to upload SVGs and download the resulting PDFs.

---

## Dockerfile Explanation

This project uses a multi-stage Dockerfile to build and run the application.

- **Stage 1: Build Stage**
  - Uses `golang:1.24-alpine` as the base image.
  - Installs Go dependencies and builds the Go binary.

- **Stage 2: Final Stage**
  - Uses the `alpine:latest` image to keep the container lightweight.
  - Installs dependencies like `librsvg` and `ca-certificates`.
  - Copies the built Go binary and HTML file to the container.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
