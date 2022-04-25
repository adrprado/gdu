package common

import (
	iofs "io/fs"

	"github.com/dundee/gdu/v5/pkg/fs"
)

// CurrentProgress struct
type CurrentProgress struct {
	CurrentItemName string
	ItemCount       int
	TotalSize       int64
}

// ShouldDirBeIgnored whether path should be ignored
type ShouldDirBeIgnored func(name, path string) bool

type ShouldFileBeIgnored func(info iofs.FileInfo) bool

// Analyzer is type for dir analyzing function
type Analyzer interface {
	AnalyzeDir(path string, ignore ShouldDirBeIgnored, ignoreBefore ShouldFileBeIgnored, constGC bool) fs.Item
	GetProgressChan() chan CurrentProgress
	GetDone() SignalGroup
	ResetProgress()
}
