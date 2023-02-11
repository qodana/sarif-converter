package codequality

func severity(level string) string {
	switch level {
	case "error":
		return "critical"
	case "warning":
		return "major"
	case "note":
		return "minor"
	case "none":
		return "info"
	}

	return "unknown"
}
