package alg

type Algorithm interface {
	String() string
	Encrypt(header, payload, secret string) ([]byte, error)
}
