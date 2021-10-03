package oss

import (
	"io"
)

type Client interface {
	Upload(file io.Reader, filename string) (string, error)
	Download(url string) error
	Delete(filename string) error
}
