package sync

import (
	"io"
	"sync"
)

// Read sync struct that implements io.Reader
type Read struct {
	mu sync.Mutex
	r  io.Reader
}

// NewRead returns a pointer to Read instance
//  Ex. stdin := sync.NewRead(os.Stdin)
func NewRead(r io.Reader) *Read {
	return &Read{r: r}
}

// Read synchronous io.Reader implementation
func (r *Read) Read(b []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.r.Read(b)
}
