package runtime

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

func (val *Value) String() string {
	if val.Type == ValueTypeList {
		return fmt.Sprintf("%s", *val.ListPtr)
	} else if val.Type == ValueTypeMap {
		return fmt.Sprintf("%s", *val.MapPtr)
	} else if val.Type == ValueTypeStr {
		return fmt.Sprintf("'%s'", val.Val)
	} else if val.Type == ValueTypeNil {
		return "nil"
	}
	return fmt.Sprintf("%s", val.Val)
}

// StringWithType ...
func (val *Value) StringWithType() string {
	if val.Type == ValueTypeList {
		return fmt.Sprintf("%s(%s)", *val.ListPtr, val.Type)
	} else if val.Type == ValueTypeMap {
		return fmt.Sprintf("%s(%s)", *val.MapPtr, val.Type)
	}
	return fmt.Sprintf("%s(%s)", val.Val, val.Type)
}

// GetValue returns the real value of the Value
func (val *Value) GetValue() interface{} {
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
func (val *Value) Equals(val2 *Value) bool {
	if val.Type == ValueTypeInt && val2.Type == ValueTypeInt {
		// compare ints
		return val.GetValue().(int) == val2.GetValue().(int)
	} else if val.Type == ValueTypeList && val2.Type == ValueTypeList {
		// compare lists
		v1 := val.GetValue().(*List)
		v2 := val2.GetValue().(*List)
		if v1.Len() != v2.Len() {
			return false
		}
		for i := 0; i < v1.Len().GetValue().(int); i++ {
			if !v1.GetByIndex(i).Equals(v2.GetByIndex(i)) {
				return false
			}
		}
		return true
	}
	// TODO compare map
	return val.Type == val2.Type && val.Val == val2.Val
}

// IsGreaterThan ..
func (val *Value) IsGreaterThan(val2 *Value) bool {
	if val.Type == ValueTypeInt && val2.Type == ValueTypeInt {
		return val.GetValue().(int) > val2.GetValue().(int)
	}
	// TODO compare list & map
	return val.Type == val2.Type && val.Val > val2.Val
}

// MakeCopy ..
// only available for str, list, nil
func (val *Value) MakeCopy() *Value {
	switch val.Type {
	case ValueTypeInt:
		return NewValue(val.GetValue().(int))
	case ValueTypeStr:
		return NewValue(val.GetValue().(string))
	case ValueTypeNil:
		return NewValue(nil)
	default:
		panic("Invalid arg data type")
	}
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
	Head    *ListNode
	Tail    *ListNode
	Length  int
	HeadIdx int
	IdxMap  map[int]*ListNode
}

// Len returns the length of the list
func (list *List) Len() *Value {
	return NewValue(list.Length)
}

// GetByIndex ..
func (list *List) GetByIndex(idx int) *Value {
	if idx < 0 || idx >= list.Length {
		return NewValue(nil)
	}
	return list.IdxMap[list.HeadIdx+idx].Data
}

// SetByIndex ..
func (list *List) SetByIndex(idx int, val *Value) {
	if idx < 0 {
		return
	}
	if idx >= list.Length {
		for idx > list.Length {
			list.Push(NewValue(nil))
		}
		list.Push(val)
		return
	}
	// replace target by the new node
	target := list.IdxMap[idx]
	newNode := &ListNode{}
	newNode.Data = val
	if target.Prev != nil {
		newNode.Prev = target.Prev
		target.Prev.Next = newNode
	} else {
		list.Head = newNode
	}
	newNode.Next = target.Next
	if target.Next != nil {
		target.Next.Prev = newNode
	} else {
		list.Tail = newNode
	}
	list.IdxMap[idx] = newNode
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
		list.Length = 0
	} else {
		list.Head.Next.Prev = nil
		list.Head = list.Head.Next
		list.Length--
	}
	delete(list.IdxMap, list.HeadIdx)
	list.HeadIdx++
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
		list.Length = 0
	} else {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Length--
	}
	delete(list.IdxMap, list.HeadIdx+list.Length)
	return data
}

// Push adds an element to the right side of the list
func (list *List) Push(val *Value) {
	newNode := &ListNode{}
	newNode.Data = val
	if list.Head == nil {
		list.Tail = newNode
		list.Head = newNode
		list.Length = 1
		list.IdxMap = make(map[int]*ListNode)
	} else {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
		list.Length++
	}
	list.IdxMap[list.HeadIdx+list.Length-1] = newNode
}

func (list List) String() string {
	var str string = "["
	head := list.Head
	if head != nil {
		for head != nil {
			str += (head.Data.String() + ",")
			head = head.Next
		}
	}
	if len(str) > 1 {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}

// ToValueArray ..
func (list *List) ToValueArray() []*Value {
	var arr []*Value
	head := list.Head
	if head != nil {
		for head != nil {
			arr = append(arr, head.Data)
			head = head.Next
		}
	}
	return arr
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
	return &List{
		Head:   head,
		Tail:   node,
		Length: len(nums),
	}
}

// Map ..
type Map struct {
	Data map[string]*Value
	Keys []string
}

func (m Map) String() string {
	var str string = "{"
	for _, k := range m.Keys {
		str += fmt.Sprintf("%s:%v", k, m.Data[k])
		str += ","
	}
	if len(str) > 1 {
		str = str[:len(str)-1]
	}
	str += "}"
	return str
	// return fmt.Sprintf("%s", m.Data)
}

// Put adds new values to the map
func (m *Map) Put(key string, val *Value) {
	if m.Data == nil {
		m.Data = make(map[string]*Value)
	}
	m.Data[key] = val
	s := indexOf(key, m.Keys)
	if s == -1 {
		m.Keys = append(m.Keys, key)
	}
}

// Get returns the value by key
func (m *Map) Get(key string) *Value {
	val, ok := m.Data[key]
	if ok {
		return val
	}
	return NewValue(nil)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// Delete removes a key-value from the map
func (m *Map) Delete(key string) {
	_, ok := m.Data[key]
	if ok {
		delete(m.Data, key)
	}
	s := indexOf(key, m.Keys)
	m.Keys = append(m.Keys[:s], m.Keys[s+1:]...)
}

// GetKeys returns a list of all keys in the map
func (m *Map) GetKeys() *List {
	keys := List{}
	for _, k := range m.Keys {
		keys.Push(NewValue(k))
	}
	return &keys
}
