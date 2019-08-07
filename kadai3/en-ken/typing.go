package typing

import (
	"math/rand"
	"time"
)

// Typing is the class to judge input
type Typing struct {
	dict     []string
	nextText string
}

// NewTyping is a constructor
func NewTyping(textDict []string) *Typing {
	rand.Seed(time.Now().UnixNano())

	return &Typing{
		dict: textDict,
	}
}

// GetNextText returns next text
func (t *Typing) GetNextText() string {
	i := rand.Int() % len(t.dict)
	t.nextText = t.dict[i]
	return t.nextText
}

// IsCorrect judges if input is correct or not.
func (t *Typing) IsCorrect(inputText string) bool {
	return t.nextText == inputText
}
