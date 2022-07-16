package report

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Reporter interface {
	Title() string
	SetTitle(title string)
	Print(msg string)
	Generate()
	Reset()
}

type BaseReport struct {
	title string
	buf   *bytes.Buffer
	w     io.Writer
}

func NewBaseReport(title string, w io.Writer) *BaseReport {
	return &BaseReport{title: title, w: w, buf: bytes.NewBuffer([]byte{})}
}

func (b *BaseReport) SetTitle(title string) {
	b.title = title
}

func (b *BaseReport) Title() string {
	return b.title
}

func (b *BaseReport) Print(msg string) {
	b.buf.WriteString(msg)
}

func (b *BaseReport) Generate() {
	_, _ = fmt.Fprint(b.w, b.title, b.buf.String())
}

func (b *BaseReport) Reset() {
	b.buf.Reset()
}

type HTMLReport struct {
	*BaseReport
}

func NewHTMLReport(bReport *BaseReport) *HTMLReport {
	return &HTMLReport{bReport}
}

func (b *HTMLReport) Generate() {
	_, _ = fmt.Fprint(b.w,
		strings.ReplaceAll(htmlReportHead, "{{_TITLE_}}", stripTags(b.title)), b.buf.String(), htmlReportFooter)
}

var htmlReportHead = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta http-equiv="x-ua-compatible" content="ie=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>{{_TITLE_}}</title>
<style>
body {
  font-family: Arial, serif;
}
code {background-color: #F5F5F5; display: block;}
</style>
  </head>
  <body>`
var htmlReportFooter = `</body>
</html>`
