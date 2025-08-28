package content

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const sniffLen = 512

type Detector struct {
}

func NewDetector() *Detector {
	return &Detector{}
}

func (d *Detector) DetectType(contentReader io.Reader) (string, error) {
	r := bufio.NewReaderSize(contentReader, sniffLen)
	headerBytes, err := r.Peek(sniffLen)
	if err != nil {
		return "", fmt.Errorf("detecting content type: %w", err)
	}

	contentType := strings.ToLower(http.DetectContentType(headerBytes))
	return contentType, nil
}
