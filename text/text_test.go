package text

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type test struct {
		name  string
		input string
		want  []string
	}

	tests := []test{
		{name: "simple", input: "hello world", want: []string{"hello", "world"}},
		{name: "all_caps", input: "PIKACHU CHARMANDER BULBASAUR TRECKO", want: []string{"pikachu", "charmander", "bulbasaur", "trecko"}},
		{name: "much_whitespace", input: "       pikachu         bulbaSAUR    jiGGLYPUFF  ", want: []string{"pikachu", "bulbasaur", "jigglypuff"}},
		{name: "none", input: "", want: []string{}},
	}

	for _, tc := range tests {
		got := cleanInput(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
