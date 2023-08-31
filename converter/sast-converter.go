package converter

import (
	bytes2 "bytes"
	"encoding/json"
	"gitlab.com/gitlab-org/security-products/analyzers/report/v4"
	"os"
	"sarif-converter/meta"
	"sarif-converter/now"
	"sarif-converter/sast/sarif"
)

type sastConverter struct {
	time     *now.TimeProvider
	metadata meta.Metadata
}

var Sast = sastConverter{}

func (c sastConverter) Type() string {
	return "sast"
}

func (c sastConverter) Convert(input []byte) ([]byte, error) {
	r, err := sarif.FromBytes(input)
	if err != nil {
		return nil, err
	}

	sast, err := transformToGLSASTReport(input, err)
	if err != nil {
		return nil, err
	}

	if c.time != nil {
		r = r.WithTimeProvider(c.time)
	}

	sast.Scan.Type = "sast"
	sast.Scan = r.Override(sast.Scan, c.metadata)

	bytes, err := json.MarshalIndent(sast, "", "  ")
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c sastConverter) WithTimeProvider(p *now.TimeProvider) Converter {
	if p != nil {
		c.time = p
	}
	return c
}

func (c sastConverter) WithMetadata(metadata *meta.Metadata) sastConverter {
	if metadata != nil {
		c.metadata = *metadata
	}
	return c
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
