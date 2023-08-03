package result

type Wrappers struct {
	results []Wrapper
}

func (w Wrappers) Get(i int) Wrapper {
	return w.results[i]
}
