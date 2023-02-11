package codequality

type LocationLine struct {
	Begin int `json:"begin"`
}

type Location struct {
	Path  *string      `json:"path"`
	Lines LocationLine `json:"lines"`
}

type Element struct {
	Description *string  `json:"description"`
	Fingerprint string   `json:"Fingerprint"`
	Severity    string   `json:"severity"`
	Location    Location `json:"location"`
}
