package bdd

import "fmt"

type HTMLDecorator struct{}

func NewHTMLDecorator() *HTMLDecorator {
	return &HTMLDecorator{}
}

func (d *HTMLDecorator) _title(s string) string { return fmt.Sprintf("<h2>%s</h2>\n", s) }

func (d *HTMLDecorator) _scenario(s *ScenarioReport) string {
	return parseReport(scenarioHTMLDecoratorTpl, s)
}

func (d *HTMLDecorator) _feature(f *FeatureReport) string {
	return parseHtmlReport(featureHTMLDecoratorTpl, f)
}

var scenarioHTMLDecoratorTpl = `
{{if eq .TestStatus "PANIC"}}
<h4>PANIC! {{.TestName}}</h4>
<p>Recovered: {{.PanicMsg}}</p>
<p>Trace: {{.Trace}}</p>{{else}}
<h4><span style="color:{{if eq .TestStatus "FAIL"}}red{{else}}green{{end}}">{{.TestStatus}}</span> {{.TestName}}</h4>
<code style="background-color: #F5F5F5; display: block; padding:5px;"><ul><li>
Feature: {{.Feature}}
<ul>
	<li>Scenario: {{.Scenario}}
		<ul>
			{{range .Steps}}<li>{{ . }}</li>{{end}}
		</ul>
	</li>
</ul>
</li></ul></code>
{{if .Errors}}<p><strong>Errors:</strong></p>
<ul>
{{range .Errors}}
<li><code style="background-color: #F5F5F5; display: block; padding:5px;">{{ . }}</code></li>
{{end}}
</ul>
{{end}}{{end}}
`

var featureHTMLDecoratorTpl = `
<h3>Feature: {{ .Name }}</h3>{{range .ScenarioReports}}
{{if eq .TestStatus "PANIC"}}
<h4>PANIC! {{.TestName}}</h4>
<p>Recovered: {{.PanicMsg}}</p>
<p>Trace: {{.Trace}}</p>{{else}}
<h4><span style="color:{{if eq .TestStatus "FAIL"}}red{{else}}green{{end}}">{{.TestStatus}}</span> {{.TestName}}</h4>
<code style="background-color: #F5F5F5; display: block; padding:5px;"><ul><li>
Feature: {{.Feature}}
<ul>
	<li>Scenario: {{.Scenario}}
		<ul>
			{{range .Steps}}<li>{{ . }}</li>{{end}}
		</ul>
	</li>
</ul>
</li></ul></code>
{{if .Errors}}<p><strong>Errors:</strong></p>
<ul>
{{range .Errors}}
<li><code style="background-color: #F5F5F5; display: block; padding:5px;">{{ . }}</code></li>
{{end}}
</ul>
{{end}}{{end}}
{{end}}
`
