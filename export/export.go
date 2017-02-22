package export

import (
	"github.com/ionous/mars"
	"github.com/ionous/mars/script"
	"github.com/ionous/mars/script/backend"
)

// Export contains all of mar's built-in commands and primitives.
func Export() *mars.Package {
	if export == nil {
		export = &mars.Package{
			Name: "Export",
			// Scripts:      pkg.Scripts,
			// Tests:        pkg.Tests,
			Dependencies: mars.Dependencies(script.Package()),
			Commands:     (*ExportCommands)(nil),
			Interfaces:   (*ExportInterfaces)(nil),
		}
	}
	return export
}

var export *mars.Package

type ExportInterfaces struct {
	ScriptBundle
	StorySection
	backend.Directive
	StoryFileBootstrap
}

type ExportCommands struct {
	*Story
	*StoryFile
	*Chapter
	*Library
	*LibraryBundle
}

type SectionType int

const (
	UnknownSectionType SectionType = iota
	ChapterSectionType
	LibrarySectionType
)

type Bundle struct {
	Name     string
	Sections []StorySection
}

type ScriptBundle interface {
	GetBundle() string
}

type StoryFile struct {
	Story ScriptBundle
}

// StoryFileBootstrap provides an interface which represents the file containing story data.
// It is an artifact of the fact every command needs to implement an interface.
type StoryFileBootstrap interface {
	StoryFile() *StoryFile
}

func (s *StoryFile) StoryFile() *StoryFile {
	return s
}

// a story contains chapters and libraries
type Story Bundle

func (s Story) GetBundle() string {
	return s.Name
}

// a library packge contains only libraries
type LibraryBundle Bundle

func (p LibraryBundle) GetBundle() string {
	return p.Name
}

// base for chapters and libraries.
type StorySection interface {
	GetSectionType() SectionType
}
