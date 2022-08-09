package common

import "io"

type Reader struct {
	R           io.Reader
	ReadedBytes int
}

func (r *Reader) Read(p []byte) (int, error) {
	length, err := r.R.Read(p)

	r.ReadedBytes += length

	return length, err
}
