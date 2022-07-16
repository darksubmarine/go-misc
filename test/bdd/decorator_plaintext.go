package bdd

import "fmt"

type PlainTextDecorator struct{}

func NewPlainTextDecorator() *PlainTextDecorator {
	return &PlainTextDecorator{}
}

func (d *PlainTextDecorator) _title(s string) string { return fmt.Sprintf("* %s\n", s) }

func (d *PlainTextDecorator) _scenario(s *ScenarioReport) string {
	return parseReport(scenarioPlainTextTpl, s)
}

func (d *PlainTextDecorator) _feature(f *FeatureReport) string {
	return parseReport(featurePlainTextTpl, f)
}

var scenarioPlainTextTpl = `
{{if eq .TestStatus "PANIC"}}
--- PANIC! {{.TestName}}
Recovered: {{.PanicMsg}}
Trace: {{.Trace}}{{else}}
--- {{.TestStatus}} {{.TestName}}
Feature: {{.Feature}}
  Scenario: {{.Scenario}}{{range .Steps}}
    {{ . }}{{end}}

{{if .Errors}}Errors: {{range .Errors}}
 [x] {{ . }}{{end}}{{end}}{{end}}
`

var featurePlainTextTpl = `
--- Feature: {{ .Name }}{{range .ScenarioReports}}
{{if eq .TestStatus "PANIC"}}
--- PANIC! {{.TestName}}
Recovered: {{.PanicMsg}}
Trace: {{.Trace}}{{else}}
--- {{.TestStatus}} {{.TestName}}
Feature: {{.Feature}}
  Scenario: {{.Scenario}}{{range .Steps}}
    {{ . }}{{end}}

{{if .Errors}}Errors: {{range .Errors}}
 [x] {{ . }}{{end}}{{end}}{{end}}
{{end}}
`
