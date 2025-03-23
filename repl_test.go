package main

import "testing"

func TestCleanInput(t *testing.T) {
	// ...
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    "  h,e,l,l,o  world  ",
			expected: []string{"h,e,l,l,o", "world"},
		},
		{
			input:    "  On top of the world,                    looking down on creation  ",
			expected: []string{"on", "top", "of", "the", "world,", "looking", "down", "on", "creation"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  H,e,L,l,O  world  ",
			expected: []string{"h,e,l,l,o", "world"},
		},
		{
			input:    "  h,e,l,l,o\t  world\n  ",
			expected: []string{"h,e,l,l,o", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("got %d words, expected %d for input '%s'", len(actual), len(c.expected), c.input)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("words not as expected got %v, expected %v", word, expectedWord)
			}
		}
	}
}
