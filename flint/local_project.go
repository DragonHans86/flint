package flint

import (
	"path/filepath"
)

type LocalProject struct {
	Path string
}

func (l *LocalProject) searchPath(pattern string) bool {
	search := filepath.Join(l.Path, pattern)
	matches, _ := filepath.Glob(search)
	return len(matches) > 0
}

func (l *LocalProject) CheckReadme() bool {
	return l.searchPath("README*")
}

func (l *LocalProject) CheckContributing() bool {
	return l.searchPath("CONTRIBUTING*")
}

func (l *LocalProject) CheckLicense() bool {
	return l.searchPath("LICENSE*") || l.searchPath("COPYING*")
}

func (l *LocalProject) CheckChangelog() bool {
	return l.searchPath("CHANGELOG*")
}

func (l *LocalProject) CheckBootstrap() bool {
	return l.searchPath("script/bootstrap")
}

func (l *LocalProject) CheckTestScript() bool {
	return l.searchPath("script/test")
}

func (l *LocalProject) CheckCodeOfConduct() bool {
	return l.searchPath("CODE_OF_CONDUCT*")
}
