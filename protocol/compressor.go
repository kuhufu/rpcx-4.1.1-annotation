package protocol

import (
	"github.com/smallnest/rpcx/util"
)

// Compressor defines a common compression interface.
type Compressor interface {
	Zip([]byte) ([]byte, error)
	Unzip([]byte) ([]byte, error)
}

// GzipCompressor implements gzip compressor.
// 使用Gzip压缩
type GzipCompressor struct {
}

func (c GzipCompressor) Zip(data []byte) ([]byte, error) {
	return util.Zip(data)
}

func (c GzipCompressor) Unzip(data []byte) ([]byte, error) {
	return util.Unzip(data)
}

// 不压缩
type RawDataCompressor struct {
}

func (c RawDataCompressor) Zip(data []byte) ([]byte, error) {
	return data, nil
}

func (c RawDataCompressor) Unzip(data []byte) ([]byte, error) {
	return data, nil
}
