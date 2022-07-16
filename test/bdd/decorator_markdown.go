package bdd

import "fmt"

type MarkdownDecorator struct{}

func NewMarkdownDecorator() *MarkdownDecorator {
	return &MarkdownDecorator{}
}

func (d *MarkdownDecorator) _title(s string) string { return fmt.Sprintf("## %s\n", s) }

func (d *MarkdownDecorator) _scenario(s *ScenarioReport) string {
	return parseReport(scenarioMarkdownDecoratorTpl, s)
}

func (d *MarkdownDecorator) _feature(f *FeatureReport) string {
	return parseReport(featureMarkdownDecoratorTpl, f)
}

var scenarioMarkdownDecoratorTpl = `
{{if eq .TestStatus "PANIC"}}
#### PANIC! {{.TestName}}
Recovered: {{.PanicMsg}}
Trace: {{.Trace}}{{else}}
#### <span style="color:{{if eq .TestStatus "FAIL"}}red{{else}}green{{end}}">{{.TestStatus}}</span> {{.TestName}}
    Feature: {{.Feature}}
      Scenario: {{.Scenario}}
        {{range .Steps}}{{ . }}
        {{end}}

{{if .Errors}}**Errors:** {{range .Errors}}
 - <code>{{ . }}</code>{{end}}{{end}}{{end}}
`

var featureMarkdownDecoratorTpl = `
### Feature: {{ .Name }}{{range .ScenarioReports}}
{{if eq .TestStatus "PANIC"}}
#### PANIC! {{.TestName}}
Recovered: {{.PanicMsg}}
Trace: {{.Trace}}{{else}}
#### <span style="color:{{if eq .TestStatus "FAIL"}}red{{else}}green{{end}}">{{.TestStatus}}</span> {{.TestName}}
    Feature: {{.Feature}}
      Scenario: {{.Scenario}}
        {{range .Steps}}{{ . }}
        {{end}}

{{if .Errors}}**Errors:** {{range .Errors}}
 - <code>{{ . }}</code>{{end}}{{end}}{{end}}
{{end}}
`
