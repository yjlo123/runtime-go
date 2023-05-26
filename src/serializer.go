package runtime

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Serialize ..
func Serialize(content *Value) []byte {
	var data []byte
	if content.Type == ValueTypeStr {
		// string
		data = append(data, '"')
		strVal := strings.Replace(content.Val, "\n", "\\n", -1)
		strVal = strings.Replace(strVal, "\\", "\\\\", -1)
		strVal = strings.Replace(strVal, "\"", "\\\"", -1)
		data = append(data, []byte(strVal)...)
		data = append(data, '"')
	} else if content.Type == ValueTypeInt {
		// int
		data = []byte(content.Val)
	} else if content.Type == ValueTypeNil {
		// null
		data = []byte("null")
	} else if content.Type == ValueTypeList {
		// array
		list := content.GetValue().(*List)
		listLen := list.Len().GetValue().(int)
		data = append(data, []byte("[")...)
		data = append(data, Serialize(list.GetByIndex(0))...)
		for i := 1; i < listLen; i++ {
			data = append(data, ',')
			data = append(data, Serialize(list.GetByIndex(i))...)
		}
		data = append(data, ']')
	} else if content.Type == ValueTypeMap {
		// map
		m := content.GetValue().(*Map)
		keys := m.GetKeys()
		keysLen := keys.Len().GetValue().(int)
		data = append(data, '{')
		for i := 0; i < keysLen; i++ {
			key := keys.GetByIndex(i).GetValue().(string)
			data = append(data, '"')
			data = append(data, []byte(key)...)
			data = append(data, '"')
			data = append(data, ':')
			data = append(data, Serialize(m.Get(key))...)
			if i != keysLen-1 {
				data = append(data, ',')
			}
		}
		data = append(data, '}')
	}
	return []byte(data)
}

// Deserialize ..
func Deserialize(str string) *Value {
	if len(str) == 0 {
		return NewValue(nil)
	}

	if str[0] == '[' {
		list := &List{}
		var arr []interface{}
		json.Unmarshal([]byte(str), &arr)
		for _, v := range arr {
			val := parseJSONRec(v)
			list.Push(val)
		}
		return NewValue(list)
	}

	if str[0] == '{' {
		var result map[string]interface{}
		err := json.Unmarshal([]byte(str), &result)
		if err != nil {
			fmt.Println("[Deserialization Error] ", err)
		}
		m := &Map{}
		for k, v := range result {
			m.Put(k, parseJSONRec(v))
		}
		return NewValue(m)
	}

	return NewValue(nil)
}

func parseJSONRec(data interface{}) *Value {
	if data == nil {
		return NewValue(nil)
	}
	kind := reflect.ValueOf(data).Kind()
	if kind == reflect.Map {
		// map
		mm := &Map{}
		for k, v := range data.(map[string]interface{}) {
			mm.Put(k, parseJSONRec(v))
		}
		return NewValue(mm)
	} else if kind == reflect.Slice {
		// array -> list
		lst := &List{}
		for _, v := range data.([]interface{}) {
			lst.Push(parseJSONRec(v))
		}
		return NewValue(lst)
	} else if kind == reflect.String {
		return NewValue(data.(string))
	} else if kind == reflect.Float64 {
		// number -> int
		return NewValue(int(data.(float64)))
	}

	return NewValue(nil)
}
