package sarifreport

type LevelWrapper struct {
	level interface{}
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

func NewLevel(r *SarifResultWrapper) LevelWrapper {
	return LevelWrapper{level: r.rule().DefaultLevel()}
}
