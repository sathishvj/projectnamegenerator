package projectnamegenerator

import (
	"fmt"
	"strings"
	"testing"
)

func Test_NewName(t *testing.T) {
	n := NewName()
	c := strings.Count(n, " ")
	if c != 3 {
		t.Errorf("Did not have all 4 parts: %s", n)
	}

	png := NewGenerator()
	png.Hyphenate = true
	n = png.NewName()
	c = strings.Count(n, "-")
	if c < 3 { // some words might have hyphens
		t.Errorf("Did not have all 4 parts: %s. Hyphen count: %d", n, c)
	}

	png = NewGenerator()
	png.Hyphenate = true
	png.Lowercase = true
	n = png.NewName()
	c = strings.Count(n, "-")
	if strings.ToLower(n) != n { // some words might have hyphens
		t.Errorf("Expected lower case but wasn't: %s", n)
	}

	png = NewGenerator()
	png.Lowercase = true
	png.HasAdverb = false
	n = png.NewName()
	c = strings.Count(n, " ")
	if c != 2 { // some words might have hyphens
		t.Errorf("Removed adverb, so was expecting only 3 parts: %s. ", n)
	}

	png = NewGenerator()
	png.Lowercase = true
	png.HasAdverb = false
	png.HasNoun = false
	n = png.NewName()
	c = strings.Count(n, " ")
	if c != 1 { // some words might have hyphens
		t.Errorf("Removed adverb and noun, so was expecting only 2 parts: %s. ", n)
	}

	png = NewGenerator()
	png.Lowercase = true
	png.HasAdverb = false
	png.HasNoun = false
	png.HasVerb = false
	n = png.NewName()
	c = strings.Count(n, " ")
	if c != 0 { // some words might have hyphens
		t.Errorf("Removed adverb, verb, and noun, so was expecting only 1 part: %s. ", n)
	}

	png = NewGenerator()
	png.Lowercase = true
	png.HasAdverb = false
	png.HasNoun = false
	png.HasVerb = false
	png.HasAdjective = false
	n = png.NewName()
	if n != "" { // some words might have hyphens
		t.Errorf("Removed all parts, so was expecting an empty string: %s. ", n)
	}

	png = NewGenerator()
	for i := 0; i < 90; i++ {
		fmt.Println(png.NewName())
	}

}
