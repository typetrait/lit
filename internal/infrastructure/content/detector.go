package content

import (
	"net/http"
	"strings"
)

type Detector struct {
}

func NewDetector() *Detector {
	return &Detector{}
}

func (d *Detector) DetectType(headerBytes []byte) (string, error) {
	contentType := strings.ToLower(http.DetectContentType(headerBytes))
	return contentType, nil
}
