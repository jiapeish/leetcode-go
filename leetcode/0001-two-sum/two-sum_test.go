package _001_two_sum

import (
	"fmt"
	"testing"
)

type problem1 struct {
	input1
	output1
}

type input1 struct {
	nums 	[]int
	target 	int
}

type output1 struct {
	solu []int
}

func Test_Problem1(t *testing.T) {
	probs := []problem1{
		{
			input1{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			output1{solu: []int{0, 1}},
		},
		{
			input1{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			output1{solu: []int{1, 2}},
		},
		{
			input1{
				nums:   []int{3, 3},
				target: 6,
			},
			output1{solu: []int{0, 1}},
		},
	}

	for _, pr := range probs {
		in, out := pr.input1, pr.output1
		fmt.Printf("input: %v output: %v answer: %v\n", in, out, twoSum(in.nums, in.target))
	}
}