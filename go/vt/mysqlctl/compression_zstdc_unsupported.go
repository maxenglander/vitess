//go:build !linux
// +build !linux

package mysqlctl

import (
	"fmt"
	"io"
)

func isZstdcSupported() bool {
	return false
}

func newZstdcCompressor(w io.Writer) (io.WriteCloser, error) {
	return nil, fmt.Errorf("zstdc is not supported on this platform")
}

func newZstdcDecompressor(r io.Reader) (io.ReadCloser, error) {
	return nil, fmt.Errorf("zstdc is not supported on this platform")
}
