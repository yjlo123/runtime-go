package main

import (
	"fmt"
	"strconv"
)

const (
	ValueTypeInt  = "int"
	ValueTypeStr  = "str"
	ValueTypeList = "list"
)

// Value ..
type Value struct {
	Type    string
	Val     string
	ListPtr *List
}

func (val Value) String() string {
	if val.Type == ValueTypeList {
		return fmt.Sprintf("%s(%s)", *val.ListPtr, val.Type)
	}
	return fmt.Sprintf("%s(%s)", val.Val, val.Type)
}

// GetValue returns the real value of the Value
func (val Value) GetValue() interface{} {
	switch val.Type {
	case ValueTypeInt:
		intVal, _ := strconv.Atoi(val.Val)
		return intVal
	case ValueTypeStr:
		return val.Val
	case ValueTypeList:
		return val.ListPtr
	default:
		return nil
	}
}

// Equals ..
func (val Value) Equals(val2 *Value) bool {
	// TODO compare list & map
	return val.Type == val2.Type && val.Val == val2.Val
}

// NewValue ..
func NewValue(val interface{}) *Value {
	var dataType string
	var dataVal string
	var dataListPtr *List
	switch val.(type) {
	case int:
		dataType = ValueTypeInt
		dataVal = strconv.Itoa(val.(int))
	case string:
		dataType = ValueTypeStr
		dataVal = val.(string)
	case *List:
		dataType = ValueTypeList
		dataListPtr = val.(*List)
	default:
		fmt.Printf("Unknown data type: %T\n", val)
	}
	return &Value{
		Type:    dataType,
		Val:     dataVal,
		ListPtr: dataListPtr,
	}
}

// ListNode ..
type ListNode struct {
	Data *Value
	Prev *ListNode
	Next *ListNode
}

// List ..
type List struct {
	Head *ListNode
	Tail *ListNode
}

// Poll extrats an element from the left side of the list
func (list *List) Poll() *Value {
	if list.Head == nil {
		return nil
	}
	data := list.Head.Data
	if list.Head.Next == nil {
		list.Tail = nil
		list.Head = nil
	} else {
		list.Head.Next.Prev = nil
		list.Head = list.Head.Next
	}
	return data
}

// Pop extrats an element from the right side of the list
func (list *List) Pop() *Value {
	if list.Tail == nil {
		return nil
	}
	data := list.Tail.Data
	if list.Tail.Prev == nil {
		list.Tail = nil
		list.Head = nil
	} else {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
	}
	return data
}

// Push adds an element to the right side of the list
func (list *List) Push(val *Value) {
	newNode := &ListNode{}
	newNode.Data = val
	if list.Head == nil {
		list.Tail = newNode
		list.Head = newNode
	} else {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	}
}

func (list List) String() string {
	var str string = "["
	head := list.Head
	if head != nil {
		for head != nil {
			str += (head.Data.Val + ",")
			head = head.Next
		}
	}
	if len(str) > 1 {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}

func ConstructList(nums []interface{}) *List {
	node := &ListNode{}
	head := node
	for i, v := range nums {
		node.Data = NewValue(v)
		if i == len(nums)-1 {
			break
		}
		newNode := &ListNode{}
		node.Next = newNode
		newNode.Prev = node
		node = node.Next
	}
	return &List{head, node}
}

func testValues() {
	list := ConstructList([]interface{}{1, 2, "a", 4})

	fmt.Println(list.Poll())
	fmt.Println(list.Poll())
	fmt.Println(list.Poll())
	//fmt.Println(list)
	list.Push(NewValue(77))
	list.Push(NewValue("abc"))
	//fmt.Println(list)
	fmt.Println(list.Pop())
	fmt.Println(list.Pop())
	//fmt.Println(list)
	list2 := ConstructList([]interface{}{1, 2})
	list.Push(NewValue(list2))
	listFromList := list.Pop()
	fmt.Println(listFromList)

	v := listFromList.GetValue().(*List).Pop()
	fmt.Println(v)
}
