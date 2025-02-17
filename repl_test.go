package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Pichu Pikachu Raichu",
			expected: []string{"pichu", "pikachu", "raichu"},
		},
		{
			input:    "BULBASAUR IVYSAUR VENUSAUR",
			expected: []string{"bulbasaur", "ivysaur", "venusaur"},
		},
		{
			input:    " charmander charmeleon charizard ",
			expected: []string{"charmander", "charmeleon", "charizard"},
		},
		{
			input:    "  Squirtle  Wartortle  Blastoise  ",
			expected: []string{"squirtle", "wartortle", "blastoise"},
		},
		// more cases
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Expected word %q, got %q", expectedWord, word)
			}
		}
	}
}
