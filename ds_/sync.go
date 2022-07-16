package ds_

import (
	"github.com/darksubmarine/versatile/sync"
	"io"
)

// SyncReader returns a synchronous implementation for io.Reader
func SyncReader(reader io.Reader) *sync.Read {
	return sync.NewRead(reader)
}

// SyncWriter returns a synchronous implementation for io.Writer
func SyncWriter(writer io.Writer) *sync.Write {
	return sync.NewWrite(writer)
}
