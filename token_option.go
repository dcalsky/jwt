package jwt

// TokenOption is a reserved type, which provides some forward compatibility,
// if we ever want to introduce token creation-related options.
type TokenOption func(*Token)

func WithTokenJSONEncoder(enc JSONEncoder) TokenOption {
	return func(token *Token) {
		token.jsonEncoder = enc
	}
}

func WithTokenBase64Encoder(enc Base64Encoder) TokenOption {
	return func(token *Token) {
		token.base64Encoder = enc
	}
}
