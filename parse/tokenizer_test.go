package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeText(t *testing.T) {
	rule := NewRule()

	text := TokenizeText(
		"This is the right sentence. This sentence without end mark",
		rule,
	)

	assert.Equal(t,
		" This sentence without end mark",
		text.parsedSentences[1].original,
	)
}

func TestTokenizeTextJa(t *testing.T) {
	rule := NewJapaneseRule()

	text := TokenizeText(
		"メロスは激怒した。必ず、かの邪知暴虐の王を除かなければならぬと決意した。メロスには政治がわからぬ。メロスは、村の牧人である。笛を吹き、羊と遊んで暮らしてきた。けれども邪悪に対しては、人一倍に敏感であった。",
		rule,
	)

	assert.Equal(t,
		"必ず、かの邪知暴虐の王を除かなければならぬと決意した。",
		text.parsedSentences[1].original,
	)
}
