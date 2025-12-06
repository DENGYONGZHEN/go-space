package list

// Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
// A node in a singly linked list should have two attributes: val and next.
//  val is the value of the current node, and next is a pointer/reference to the next node.
// If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list.
// Assume all nodes in the linked list are 0-indexed.

// Implement the MyLinkedList class:

// MyLinkedList() Initializes the MyLinkedList object.
// int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
// void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
// void addAtTail(int val) Append a node of value val as the last element of the linked list.
// void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list,
// the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
// void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.

// Example 1:

// Input
// ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
// [[], [1], [3], [1, 2], [1], [1], [1]]
// Output
// [null, null, null, null, 2, null, 3]

// Explanation
// MyLinkedList myLinkedList = new MyLinkedList();
// myLinkedList.addAtHead(1);
// myLinkedList.addAtTail(3);
// myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
// myLinkedList.get(1);              // return 2
// myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
// myLinkedList.get(1);              // return 3

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

// 构造函数
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// 获取 index 位置的值
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.val
}

// 头部添加节点
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{val: val, next: this.head}
	this.head = newNode
	this.size++
}

// 尾部添加节点
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{val: val}
	if this.head == nil {
		this.head = newNode
	} else {
		current := this.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	this.size++
}

// 在 index 处插入节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size { // 超出范围，直接返回
		return
	}
	if index == 0 { // 等同于 AddAtHead
		this.AddAtHead(val)
		return
	}
	current := this.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	newNode := &Node{val: val, next: current.next}
	current.next = newNode
	this.size++
}

// 删除 index 处的节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		this.head = this.head.next
	} else {
		current := this.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}
	this.size--
}
