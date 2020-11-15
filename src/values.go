package main

import (
	"fmt"
	"strconv"
)

// Value Types ..
const (
	ValueTypeNil  = "nil"
	ValueTypeInt  = "int"
	ValueTypeStr  = "str"
	ValueTypeList = "list"
	ValueTypeMap  = "map"
)

// Value ..
type Value struct {
	Type    string
	Val     string
	ListPtr *List
	MapPtr  *Map
}

func (val Value) String() string {
	if val.Type == ValueTypeList {
		return fmt.Sprintf("%s(%s)", *val.ListPtr, val.Type)
	} else if val.Type == ValueTypeMap {
		return fmt.Sprintf("%s(%s)", *val.MapPtr, val.Type)
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
	case ValueTypeMap:
		return val.MapPtr
	case ValueTypeNil:
		return nil
	default:
		return fmt.Sprintf("Unknow value: %s", val)
	}
}

// Equals ..
func (val Value) Equals(val2 *Value) bool {
	// TODO compare list & map
	return val.Type == val2.Type && val.Val == val2.Val
}

// IsGreaterThan ..
func (val Value) IsGreaterThan(val2 *Value) bool {
	// TODO compare list & map
	return val.Type == val2.Type && val.Val > val2.Val
}

// NewValue ..
func NewValue(val interface{}) *Value {
	if val == nil {
		return &Value{
			Type: ValueTypeNil,
		}
	}

	var dataType string
	var dataVal string
	var dataListPtr *List
	var dataMapPtr *Map

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
	case *Map:
		dataType = ValueTypeMap
		dataMapPtr = val.(*Map)
	default:
		fmt.Printf("Unknown data type: %T\n", val)
	}
	return &Value{
		Type:    dataType,
		Val:     dataVal,
		ListPtr: dataListPtr,
		MapPtr:  dataMapPtr,
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
		return NewValue(nil)
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
		return NewValue(nil)
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

// ConstructList ..
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

// Map ..
type Map struct {
	Data map[string]*Value
}

func (m Map) String() string {
	return fmt.Sprintf("%s", m.Data)
}

// Put adds new values to the map
func (m *Map) Put(key string, val *Value) {
	if m.Data == nil {
		m.Data = make(map[string]*Value)
	}
	m.Data[key] = val
}

// Get returns the value by key
func (m *Map) Get(key string) *Value {
	val, ok := m.Data[key]
	if ok {
		return val
	}
	return NewValue(nil)
}

// Delete removes a key-value from the map
func (m *Map) Delete(key string) {
	_, ok := m.Data[key]
	if ok {
		delete(m.Data, key)
	}
}

// GetKeys returns a list of all keys in the map
func (m *Map) GetKeys() *List {
	keys := List{}
	for k := range m.Data {
		keys.Push(NewValue(k))
	}
	return &keys
}
