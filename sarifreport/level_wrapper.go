package sarifreport

type LevelWrapper struct {
	level string
}

func (w LevelWrapper) Severity() string {
	switch w.level {
	case "error":
		return "critical"
	case "warning":
		return "major"
	case "note":
		return "minor"
	}

	return "unknown"
}

func NewLevel(l string) LevelWrapper {
	return LevelWrapper{level: l}
}
