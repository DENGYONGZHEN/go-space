package stack

import (
	"strconv"
	"testing"
)

func TestMinStack(t *testing.T) {
	testCases := []struct {
		operations []string
		input      []string
		want       []string
	}{
		{
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			input:      []string{"", "-2", "0", "-3", "", "", "", ""},
			want:       []string{"", "", "", "", "-3", "", "0", "-2"},
		},
	}

	var stack MinStack
	for _, tc := range testCases {
		for i, operation := range tc.operations {
			switch operation {
			case "MinStack":
				stack = Constructor()
			case "push":
				if val, err := strconv.Atoi(tc.input[i]); err == nil {
					stack.Push(val)
				}
			case "getMin":
				want, err := strconv.Atoi(tc.want[i])
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if want != stack.GetMin() {
					t.Errorf("GetMin(),should get %v,but get %v", want, stack.GetMin())
				}
			case "pop":
				stack.Pop()
			case "top":
				top, err := strconv.Atoi(tc.want[i])
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if top != stack.Top() {
					t.Errorf("Top(),should get %v,but get %v", top, stack.Top())
				}
			}
		}
	}
}
