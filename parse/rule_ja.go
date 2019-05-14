package parse

import (
	"github.com/ikawaha/kagome/tokenizer"
)

type JapaneseRule struct {
	t                  tokenizer.Tokenizer
	sentenceSeparators []string
}

func NewJapaneseRule() *JapaneseRule {
	return &JapaneseRule{
		t:                  tokenizer.New(),
		sentenceSeparators: []string{"。", ".", "！", "!", "？", "?"},
	}
}

func (rule *JapaneseRule) IsSentenceSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range rule.sentenceSeparators {
		if chr == val {
			return true
		}
	}

	return false
}

func (rule *JapaneseRule) IsWordSeparator(rune rune) bool {
	return true
}

func (rule *JapaneseRule) findSentences(rawText string) Text {
	text := Text{}

	var sentence string

	for _, chr := range rawText {
		if !rule.IsSentenceSeparator(chr) {
			sentence = sentence + string(chr)
		} else if len(sentence) > 0 {
			sentence = sentence + string(chr)

			text.Append(sentence, rule.findWords(sentence))

			sentence = ""
		}
	}

	if len(sentence) > 0 {
		text.Append(sentence, rule.findWords(sentence))
	}

	return text
}

func (rule *JapaneseRule) findWords(rawSentence string) (words []string) {
	words = []string{}

	tokens := rule.t.Tokenize(rawSentence)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		if token.Features()[0] == "名詞" {
			words = append(words, token.Surface)
		}
	}
	return words
}
