package utils

import (
	"io"
)

type CountReader struct {
	Reader io.Reader
	buf    []byte
	Size   int64
}

func (r *CountReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.Size += int64(n)
	return n, err
}
