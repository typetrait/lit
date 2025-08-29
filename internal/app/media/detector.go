package media

type Detector interface {
	DetectType(headerBytes []byte) (string, error)
}
