package file

import (
	"codequality-converter/file/reader"
	"encoding/json"
	"github.com/owenrumney/go-sarif/v2/sarif"
)

type Multiple struct {
	paths []string
}

func (m Multiple) IsEmpty() bool {
	return len(m.paths) > 0
}

func (m Multiple) Read(reader reader.Reader) ([]byte, error) {
	reports, err := m.reports(reader)
	if err != nil {
		return nil, err
	}

	return json.Marshal(merge(reports))
}

func (m Multiple) Paths() []string {
	return m.paths
}

func (m Multiple) reports(reader reader.Reader) ([]*sarif.Report, error) {
	//goland:noinspection GoPreferNilSlice
	result := []*sarif.Report{}

	for _, path := range m.paths {
		r, err := report(path, reader)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}

func merge(reports []*sarif.Report) *sarif.Report {
	first := reports[0]
	rest := reports[1:]

	for _, report := range rest {
		for _, run := range report.Runs {
			first.Runs = append(first.Runs, run)
		}
	}

	return first
}

func report(path string, reader reader.Reader) (*sarif.Report, error) {
	b, err := reader.Read(path)
	if err != nil {
		return nil, err
	}
	return sarif.FromBytes(b)
}

func newMultipleFiles(paths []string) Multiple {
	return Multiple{paths: paths}
}
