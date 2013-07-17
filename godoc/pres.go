// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package godoc

import (
	"text/template"
)

// Presentation generates output from a corpus.
type Presentation struct {
	Corpus *Corpus

	// TabWidth optionally specifies the tab width.
	TabWidth int

	ShowTimestamps bool
	ShowPlayground bool
	ShowExamples   bool
	DeclLinks      bool
}

// NewPresentation returns a new Presentation from a corpus.
func NewPresentation(c *Corpus) *Presentation {
	if c == nil {
		panic("nil Corpus")
	}
	return &Presentation{
		Corpus:       c,
		TabWidth:     4,
		ShowExamples: true,
		DeclLinks:    true,
	}
}

func (p *Presentation) FuncMap() template.FuncMap {
	panic("")
}