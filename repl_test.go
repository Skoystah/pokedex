package main

import "testing"

func TestCleanInput(t *testing.T)  {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "   hello   world   ",
            expected: []string{"hello","world"},
        },
        {
            input: "   Hello   World   ",
            expected: []string{"hello","world"},
        },
        {
            input: "ni   hao",
            expected: []string{"ni","hao"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        
        if len(actual) != len(c.expected) {
            t.Errorf("error - number of words not matching: actual = %v | expected = %v", len(actual), len(c.expected))
            t.Fail()
        }

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]

            if word != expectedWord {
                t.Errorf("error - words not matching: actual = %s | expected = %s", word, expectedWord)
                t.Fail()
            }
        }
    }
}
