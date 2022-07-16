package bdd

type Decorator interface {
	_title(s string) string
	_scenario(s *ScenarioReport) string
	_feature(f *FeatureReport) string
}
