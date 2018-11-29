package fileutil

import (
	"io"
	"os"

	"github.com/blend/go-sdk/exception"
)

// ReadChunkHandler is a receiver for a chunk of a file.
type ReadChunkHandler func(line []byte) error

// ReadByChunks reads a file in `chunkSize` pieces, dispatched to the handler.
func ReadByChunks(filePath string, chunkSize int, handler ReadChunkHandler) error {
	f, err := os.Open(filePath)

	if err != nil {
		return exception.New(err)
	}
	defer f.Close()

	chunk := make([]byte, chunkSize)
	for {
		readBytes, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		readData := chunk[:readBytes]
		err = handler(readData)
		if err != nil {
			return exception.New(err)
		}
	}
	return nil
}
