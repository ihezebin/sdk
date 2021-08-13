package jwt

type Signature struct {
	Header  *Header  `json:"header"`
	Payload *Payload `json:"payload"`
	Secret  string   `json:"secret"`
}

func NewSignature(header *Header, payload *Payload, secret string) *Signature {
	return &Signature{Header: header, Payload: payload, Secret: secret}
}

func (signature *Signature) Encrypt() (string, error) {
	return signature.Header.Algorithm.Handler(signature.Header, signature.Payload, signature.Secret)
}
