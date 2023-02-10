package argument

type files struct {
	inputs []string
	output string
}

func (f files) isValid() bool {
	return f.output != ""
}

func parseFileArguments(arguments []string) files {
	length := len(arguments)
	if length > 0 {
		return files{
			inputs: arguments[:length-1],
			output: arguments[length-1],
		}
	}
	return files{
		inputs: []string{},
	}
}
