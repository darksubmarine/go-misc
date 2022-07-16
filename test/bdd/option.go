package bdd

import "github.com/darksubmarine/versatile/report"

const (
	optReport = iota + 1
	optDecorator
	optEmojiDecorator
)

type Option struct {
	key int
	val interface{}
}

func WithEmojiDecorator() Option {
	return Option{key: optEmojiDecorator, val: "emoji"}
}

func WithReport(r report.Reporter) Option {
	return Option{key: optReport, val: r}
}

func WithDecorator(d Decorator) Option {
	return Option{key: optDecorator, val: d}
}

func OptReport(m map[int]interface{}) report.Reporter {
	if rep, ok := m[optReport]; ok {
		if r, ok := rep.(report.Reporter); ok {
			return r
		}
	}

	return nil
}

func OptDecorator(m map[int]interface{}) Decorator {
	if dec, ok := m[optDecorator]; ok {
		if d, ok := dec.(Decorator); ok {
			return d
		}
	}

	return nil
}
