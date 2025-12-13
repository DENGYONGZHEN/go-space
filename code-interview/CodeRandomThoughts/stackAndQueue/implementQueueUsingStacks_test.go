package stackandqueue

import (
	"reflect"
	"testing"
)

// Example 1:
// Input
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// Output
// [nil, null, null, 1, 1, false]
// Explanation
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.peek(); // return 1
// myQueue.pop(); // return 1, queue is [2]
// myQueue.empty(); // return false

func TestImplementQueueUsingStack(t *testing.T) {

	testCases := []struct {
		input []string
		want  []any
	}{
		{input: []string{"MyQueue", "push", "push", "peek", "pop", "empty"}, want: []any{nil, nil, nil, 1, 1, false}},
	}

	for _, tc := range testCases {
		var myQueue MyQueue
		result := []any{}
		count := 1
		for _, operate := range tc.input {
			switch operate {
			case "MyQueue":
				myQueue = QueueConstructor()
				result = append(result, nil)
			case "push":
				myQueue.Push(count)
				count++
				result = append(result, nil)
			case "peek":
				result = append(result, myQueue.Peek())
			case "pop":
				result = append(result, myQueue.Pop())
			case "empty":
				result = append(result, myQueue.Empty())
			}
		}

		if !reflect.DeepEqual(tc.want, result) {
			t.Logf("TestImplementQueueUsingStack want %v, except %v", tc.want, result)
		}
	}

}
