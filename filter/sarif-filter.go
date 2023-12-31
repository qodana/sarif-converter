package filter

import (
	"encoding/json"
	"github.com/owenrumney/go-sarif/v2/sarif"
	"os"
	"sarif-converter/main/argument"
)

func AllSarifFilter(report []byte, a *argument.Arguments) ([]byte, error) {
	s, err := sarif.FromBytes(report)
	if err != nil {
		return nil, err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	srcRoot := a.SrcRoot(pwd)
	if srcRoot != nil {
		f := NewRelativePathFilter(*srcRoot)
		f.Run(s)
	}

	return json.Marshal(s)
}
