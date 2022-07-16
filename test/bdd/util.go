package bdd

import (
	"bytes"
	"fmt"
	htmlTemplate "html/template"
	"strings"
	"testing"
	"text/template"
)

func parseStack(str string) []string {
	fullStack := strings.Split(str, "\n")
	var stack = []string{}
	for _, s := range fullStack {
		if strings.HasPrefix(s, "\t") {
			stack = append(stack, s)
		}
	}

	return stack
}

func stack_(stackTrace []byte) string {
	stack := string(stackTrace)
	return fmt.Sprintln(strings.Join(parseStack(stack), "\n"))
}

func recover_(msg interface{}, stack string, t *testing.T) {
	fmt.Printf("Test %s ends with panic!\nRecovered from %s\n", t.Name(), msg)
	fmt.Println("Trace:")
	fmt.Println(stack)
	t.FailNow()
}

func parseReport(tpl string, data interface{}) string {
	tp, _ := template.New("").Parse(tpl)
	var buf = bytes.NewBuffer([]byte{})
	_ = tp.Execute(buf, data)

	return buf.String()
}

func parseHtmlReport(tpl string, data interface{}) string {
	tp, _ := htmlTemplate.New("").Parse(tpl)
	var buf = bytes.NewBuffer([]byte{})
	_ = tp.Execute(buf, data)

	return buf.String()
}
