package sync

import (
	"io"
	"sync"
)

// Write sync struct that implements io.Writer
type Write struct {
	mu sync.Mutex
	w  io.Writer
}

// NewWrite returns a pointer to Write instance
//  Ex. stdout := sync.NewWrite(os.Stdout)
func NewWrite(w io.Writer) *Write {
	return &Write{w: w}
}

// Write synchronous io.Writer implementation
func (w *Write) Write(b []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.w.Write(b)
}
