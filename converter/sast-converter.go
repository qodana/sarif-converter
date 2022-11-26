package converter

import (
	bytes2 "bytes"
	"encoding/json"
	"github.com/owenrumney/go-sarif/sarif"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v3"
	"os"
)

func ConvertToSast(input []byte) ([]byte, error) {
	s, err := sarif.FromBytes(input)
	if err != nil {
		return nil, err
	}

	sast, err := transformToGLSASTReport(input, err)
	if err != nil {
		return nil, err
	}

	sast.Scan.Scanner.ID = s.Runs[0].Tool.Driver.Name
	sast.Scan.Scanner.Name = s.Runs[0].Tool.Driver.Name
	sast.Scan.Type = "sast"

	bytes, err := json.MarshalIndent(sast, "", "  ")
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func transformToGLSASTReport(input []byte, err error) (*report.Report, error) {
	original := os.Getenv("GITLAB_FEATURES")
	os.Unsetenv("GITLAB_FEATURES")
	sast, err := report.TransformToGLSASTReport(bytes2.NewReader(input), "", "", report.Scanner{})
	os.Setenv("GITLAB_FEATURES", original)

	return sast, err
}
