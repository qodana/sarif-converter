package argument

type Arguments struct {
	args []string
}

func (a Arguments) IsValid() bool {
	if a.RequireShowVersion() {
		return true
	}

	if a.Input() != "" && a.Output() != "" {
		return true
	}

	return false
}

func (a Arguments) RequireShowVersion() bool {
	for _, arg := range a.args {
		if arg == "--version" {
			return true
		}
		if arg == "-v" {
			return true
		}
	}
	return false
}

func (a Arguments) Input() string {
	if len(a.args) > 1 {
		return a.args[1]
	}
	return ""
}

func (a Arguments) Output() string {
	if len(a.args) > 2 {
		return a.args[2]
	}
	return ""
}

func Parse(args []string) Arguments {
	return Arguments{args: args}
}
