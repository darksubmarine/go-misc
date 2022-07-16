package bdd

type FeatureReport struct {
	Name            string
	ScenarioReports []*ScenarioReport
}

type ScenarioReport struct {
	TestName   string
	TestStatus string
	PanicMsg   string
	Trace      string
	Feature    string
	Scenario   string
	Steps      []string
	Errors     []string
}
