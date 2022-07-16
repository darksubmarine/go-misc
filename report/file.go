package report

import (
	"fmt"
	"os"
)

func NewFile(name string) *BaseReport {
	f, _ := os.Create(name)
	return NewBaseReport("", f)
}

func NewHTMLFile(name string) *HTMLReport {
	return NewHTMLReport(NewFile(fmt.Sprintf("%s.html", name)))
}
