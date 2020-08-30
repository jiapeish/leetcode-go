package _003_longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

type problem3 struct {
	input3
	output3
}

type input3 struct {
	s string
}

type output3 struct {
	solu int
}

func Test_Problem3(t *testing.T) {
	probs := []problem3{
		{
			input3{s: "abcabcbb"},
			output3{solu: 3},
		},
		{
			input3{s: "bbbbb"},
			output3{solu: 1},
		},
		{
			input3{s: "pwwkew"},
			output3{solu: 3},
		},
	}

	for _, pr := range probs {
		in, out := pr.input3, pr.output3
		fmt.Printf("input: %v\toutput: %v\tanswer: %v\n", in, out, lengthOfLongestSubstring(in.s))
	}


}
