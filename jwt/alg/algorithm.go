package alg

import (
	"sync"
)

type Algorithm interface {
	// Name The name of the algorithm
	Name() string
	// Encrypt Only pay attention to how to encrypt, return the encrypted original Signature data.
	// Don't do encoding processing; signing is a string that connects the header and the payload after encoding through the point ".".
	// 只关注如何加密, 返回加密后的原始Signature数据, 不要做编码处理; signing为将header和payload编码后的通过点连接起来的字符串.
	Encrypt(signing, secret string) ([]byte, error)
}

type manager struct {
	algorithmMap map[string]Algorithm
	lock         *sync.RWMutex
}

var gManager = manager{
	algorithmMap: make(map[string]Algorithm, 0),
	lock:         new(sync.RWMutex),
}

func RegisterAlg(algorithm Algorithm) {
	gManager.lock.Lock()
	defer gManager.lock.Unlock()
	gManager.algorithmMap[algorithm.Name()] = algorithm
}

func GetAlg(name string) Algorithm {
	gManager.lock.Lock()
	defer gManager.lock.Unlock()
	return gManager.algorithmMap[name]
}
