package input

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	case1 := testCase{input: "    hello  world   ", expected: []string{"hello", "world"}}
	case2 := testCase{input: "", expected: []string{}}
	case3 := testCase{input: "charizard pikachu      bulbasaur    squirtle", expected: []string{"charizard", "pikachu", "bulbasaur", "squirtle"}}
	case4 := testCase{input: "PIKachU CHARIZARD     ivysaur      GRENinJa", expected: []string{"pikachu", "charizard", "ivysaur", "greninja"}}

	cases := []testCase{case1, case2, case3, case4}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Actual length (%d) doesn't match expected length (%d): ", len(actual), len(c.expected))
		}
		for i, _ := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words at index %d dont match. Expected: %s | Got: %s", i, expectedWord, word)
			}
		}
	}
}
