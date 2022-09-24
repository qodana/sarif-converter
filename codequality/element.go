package codequality

type CodeQualityLocationLine struct {
	Begin int `json:"begin"`
}

type CodeQualityLocation struct {
	Path  *string                 `json:"path"`
	Lines CodeQualityLocationLine `json:"lines"`
}

type CodeQualityElement struct {
	Description *string             `json:"description"`
	Fingerprint string              `json:"fingerprint"`
	Severity    string              `json:"severity"`
	Location    CodeQualityLocation `json:"location"`
}
