package element

type LocationLine struct {
	Begin int `json:"begin"`
}

type Location struct {
	Path  *string       `json:"path"`
	Lines *LocationLine `json:"lines,omitempty"`
}

type Element struct {
	CheckName   *string  `json:"check_name"`
	Description *string  `json:"description"`
	Fingerprint string   `json:"fingerprint"`
	Severity    string   `json:"severity"`
	Location    Location `json:"location"`
}
