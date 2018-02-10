package rank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	connectionLeft := make(map[int]int)
	connectionLeft[2] = 1
	connectionLeft[5] = 3

	connectionRight := make(map[int]int)
	connectionRight[65] = 4
	connectionRight[74] = 12

	word := &Word{
		0,
		[]int{1},
		connectionLeft,
		connectionRight,
		"apple",
		1,
		0,
	}

	words := make(map[int]*Word)
	words[0] = word

	wordValIDs := make(map[string]int)
	wordValIDs["apple"] = 0

	sentences := make(map[int]string)
	sentences[1] = "Old apple tree in sunshine."

	rank := Rank{
		0,
		0,
		Relation{
			0,
			0,
			make(map[int]map[int]Score),
		},
		sentences,
		words,
		wordValIDs,
	}

	assert.EqualValues(t, "apple", rank.Words[0].Token)
	assert.EqualValues(t, 1, rank.Words[0].SentenceIDs[0])
	assert.EqualValues(t, 0, rank.WordValID["apple"])
}