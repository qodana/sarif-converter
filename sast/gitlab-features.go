package sast

import "os"

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
