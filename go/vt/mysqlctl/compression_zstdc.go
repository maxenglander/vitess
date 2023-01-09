//go:build linux
// +build linux

package mysqlctl

import (
	"io"

	"github.com/valyala/gozstd"
)

func isZstdcSupported() bool {
	return true
}

func newZstdcCompressor(w io.Writer) (io.WriteCloser, error) {
	return gozstd.NewWriterLevel(w, compressionLevel), nil
}

func newZstdcDecompressor(r io.Reader) (io.ReadCloser, error) {
	return io.NopCloser(gozstd.NewReader(r)), nil
}
