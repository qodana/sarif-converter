package converter

import (
	bytes2 "bytes"
	"encoding/json"
	"github.com/owenrumney/go-sarif/v2/sarif"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"os"
)

type sastConverter struct {
}

var Sast = sastConverter{}

func (c sastConverter) Type() string {
	return "sast"
}

func (c sastConverter) Convert(input []byte) ([]byte, error) {
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
