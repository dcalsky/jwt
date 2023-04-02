package jwt

// Base64Encoder is an interface that allows to implement custom Base64 encoding/decoding algorithms.
type Base64Encoder interface {
	EncodeToString(src []byte) string
	DecodeString(s string) ([]byte, error)
}

// JSONEncoder is an interface that allows to implement custom JSON encoding/decoding algorithms.
type JSONEncoder interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}
