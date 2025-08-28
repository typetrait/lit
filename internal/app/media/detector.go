package media

import "io"

type Detector interface {
	DetectType(contentReader io.Reader) (string, error)
}
