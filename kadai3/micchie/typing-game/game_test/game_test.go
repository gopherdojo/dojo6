package game_test

import (
	"testing"

	"github.com/gopherdojo/dojo6/kadai3/micchie/typing-game/game"
)

func TestGame_Run(t *testing.T) {
}

func TestGame_Judgment(t *testing.T) {
	g := &game.Game{}
	tests := []struct {
		name                string
		word                string
		input               string
		beforeCorrectNumber int
		afterCorrectNumber  int
		result              bool
	}{
		{name: "nomal", word: "bear", input: "bear", beforeCorrectNumber: 1, afterCorrectNumber: 2, result: true},
		{name: "blank", word: "cat", input: "", beforeCorrectNumber: 1, afterCorrectNumber: 1, result: false},
		{name: "incorrect", word: "dog", input: "d", beforeCorrectNumber: 0, afterCorrectNumber: 0, result: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			g.Score.CorrectNumber = test.beforeCorrectNumber
			result := g.Judgment(test.word, test.input)

			if result != test.result {
				t.Errorf(
					"word: %v, input: %v, beforeCorrect: %v, afterCorrect: %v (want: %v, got: %v)",
					test.word,
					test.input,
					test.beforeCorrectNumber,
					test.afterCorrectNumber,
					test.result,
					result,
				)
			}
		})
	}
}
