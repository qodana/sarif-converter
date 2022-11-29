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
	gf := newGitLabFeatures()

	gf.unset()
	sast, err := report.TransformToGLSASTReport(bytes2.NewReader(input), "", "", report.Scanner{})
	gf.restore()

	return sast, err
}

type gitlabFeatures struct {
	original string
}

func newGitLabFeatures() gitlabFeatures {
	return gitlabFeatures{}
}

func (f gitlabFeatures) unset() {
	f.original = os.Getenv("GITLAB_FEATURES")
	_ = os.Unsetenv("GITLAB_FEATURES")
}

func (f gitlabFeatures) restore() {
	_ = os.Setenv("GITLAB_FEATURES", f.original)
}