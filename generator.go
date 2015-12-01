package projectnamegenerator

import (
	"math/rand"
	"strings"
	"time"
)

type ProjectNameGenerator struct {
	Hyphenate    bool
	Lowercase    bool
	HasAdverb    bool
	HasVerb      bool
	HasAdjective bool
	HasNoun      bool

	Adverbs    [][]string
	Verbs      [][]string
	Adjectives [][]string
	Nouns      [][]string
}

func NewGenerator() *ProjectNameGenerator {

	adjectives := [][]string{
		Appearance_Adjectives,
		Color_Adjectives,
		Condition_Adjectives,
		Negative_Feelings_Adjectives,
		Good_Feelings_Adjectives,
		Shape_Adjectives,
		Size_Adjectives,
		Sound_Adjectives,
		Time_Adjectives,
		TasteTouch_Adjectives,
		Touch_Adjectives,
		Quantity_Adjectives,
	}

	nouns := [][]string{
		Nouns,
	}

	adverbs := [][]string{
		Adverbs,
	}

	verbs := [][]string{
		Verbs,
	}

	png := ProjectNameGenerator{
		Hyphenate:    false,
		Lowercase:    false,
		HasAdverb:    true,
		HasVerb:      true,
		HasAdjective: true,
		HasNoun:      true,

		Adverbs:    adverbs,
		Verbs:      verbs,
		Adjectives: adjectives,
		Nouns:      nouns,
	}

	return &png
}

func (png *ProjectNameGenerator) NewName() string {

	rand.Seed(time.Now().UTC().UnixNano())
	adjSet := rand.Intn(len(png.Adjectives))
	nounSet := rand.Intn(len(png.Nouns))
	adverbSet := rand.Intn(len(png.Adverbs))
	verbSet := rand.Intn(len(png.Verbs))

	adjItem := rand.Intn(len(png.Adjectives[adjSet]))
	nounItem := rand.Intn(len(png.Nouns[nounSet]))
	adverbItem := rand.Intn(len(png.Adverbs[adverbSet]))
	verbItem := rand.Intn(len(png.Verbs[verbSet]))

	var parts []string
	if png.HasAdverb {
		parts = append(parts, png.Adverbs[adverbSet][adverbItem])
	}
	if png.HasVerb {
		parts = append(parts, png.Verbs[verbSet][verbItem])
	}
	if png.HasAdjective {
		parts = append(parts, png.Adjectives[adjSet][adjItem])
	}
	if png.HasNoun {
		parts = append(parts, png.Nouns[nounSet][nounItem])
	}

	sep := " "
	if png.Hyphenate {
		sep = "-"
	}

	s := strings.Join(parts, sep)
	if png.Lowercase {
		s = strings.ToLower(s)
	}
	return s
}

func NewName() string {
	png := NewGenerator()
	return png.NewName()
}
