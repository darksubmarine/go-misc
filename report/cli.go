package report

import (
	"github.com/darksubmarine/versatile/sync"
	"os"
)

func NewStdOut() *BaseReport {
	stdout := sync.NewWrite(os.Stdout)
	return NewBaseReport("", stdout)
}
