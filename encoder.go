package jwt

import "io"

type Base64Encoding interface {
	EncodeToString(src []byte) string
	DecodeString(s string) ([]byte, error)
}

type Stricter[T Base64Encoding] interface {
	Strict() T
}

// JSONMarshalFunc is an function type that allows to implement custom JSON
// encoding algorithms.
type JSONMarshalFunc func(v any) ([]byte, error)

// JSONUnmarshalFunc is an function type that allows to implement custom JSON
// unmarshal algorithms.
type JSONUnmarshalFunc func(data []byte, v any) error

type JSONDecoder interface {
	UseNumber()
	Decode(v any) error
}

type JSONNewDecoderFunc[T JSONDecoder] func(r io.Reader) T