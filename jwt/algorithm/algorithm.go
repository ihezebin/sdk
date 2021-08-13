package algorithm

import (
	"github.com/whereabouts/sdk/jwt"
)

type AlgHandler func(header *jwt.Header, payload *jwt.Payload, secret string) (string, error)

type Algorithm struct {
	Name    string
	Handler AlgHandler
}

func NewAlgorithm(name string, handler AlgHandler) *Algorithm {
	return &Algorithm{name, handler}
}

func (algorithm *Algorithm) String() string {
	return algorithm.Name
}
