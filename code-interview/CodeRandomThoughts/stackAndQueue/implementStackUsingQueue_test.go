package stackandqueue

import (
	"reflect"
	"testing"
)

// Example 1:

// Input
// ["MyStack", "push", "push", "top", "pop", "empty"]
// [[], [1], [2], [], [], []]
// Output
// [null, null, null, 2, 2, false]

// Explanation
// MyStack myStack = new MyStack();
// myStack.push(1);
// myStack.push(2);
// myStack.top(); // return 2
// myStack.pop(); // return 2
// myStack.empty(); // return False

func TestImplementStackUsingQueue(t *testing.T) {

	testCase := []struct {
		input []string
		want  []any
	}{
		{
			input: []string{"MyStack", "push", "push", "top", "pop", "empty"},
			want:  []any{nil, nil, nil, 2, 2, false},
		},
	}
	for _, tc := range testCase {
		var result []any
		number := 1
		var myStack MyStack

		for _, operation := range tc.input {
			switch operation {
			case "MyStack":
				myStack = StackConstructor()
				result = append(result, nil)
			case "push":
				myStack.Push(number)
				number++
				result = append(result, nil)
			case "top":
				result = append(result, myStack.Top())
			case "pop":
				result = append(result, myStack.Pop())
			case "empty":
				result = append(result, myStack.Empty())
			}
		}

		if !reflect.DeepEqual(result, tc.want) {
			t.Fatalf("TestImplementStackUsingQueue want %v, instead %v", tc.want, result)
		}
	}
}
