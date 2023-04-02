package jwt_test

import (
	"encoding/base64"
	"encoding/json"
)

type customJSONEncoder struct{}

func (s *customJSONEncoder) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (s *customJSONEncoder) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

type customBase64Encoder struct{}

func (s *customBase64Encoder) EncodeToString(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (s *customBase64Encoder) DecodeString(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
