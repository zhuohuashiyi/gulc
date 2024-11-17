package gulc

import (
	"testing"
	"reflect"
)

var reflectInts = []reflect.Kind{reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64}
var reflectUInts = []reflect.Kind{reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64}
var reflectFloats = []reflect.Kind{reflect.Float32, reflect.Float64}

type TestCase struct {
	Input  []interface{}
	Output []interface{}
}


// DeepEqual 针对任意两数的深度比较，注意，有符号数3不等于无符号数3
func DeepEqual(v1, v2 interface{}) bool {
	k1 := reflect.TypeOf(v1).Kind()
	k2 := reflect.TypeOf(v2).Kind()

	if k1 == reflect.Invalid || k2 == reflect.Invalid {
		return false
	}

	
	// 首先比较有符号整数类型
	if IndexOf(reflectInts, k1) != -1 {
		if IndexOf(reflectInts, k2) == -1 {
			return false
		}
		return reflect.ValueOf(v1).Int() == reflect.ValueOf(v2).Int()
	}
	// 无符号数
	if IndexOf(reflectUInts, k1) != -1 {
		if IndexOf(reflectUInts, k2) == -1 {
			return false
		}
		return reflect.ValueOf(v1).Uint() == reflect.ValueOf(v2).Uint()
	}
	// 浮点数
	if IndexOf(reflectFloats, k1) != -1 {
		if IndexOf(reflectFloats, k2) == -1 {
			return false
		}
		return reflect.ValueOf(v1).Float() == reflect.ValueOf(v2).Float()
	}
	// 字符串
	if k1 == reflect.String {
		if k2 != reflect.String {
			return false
		}
		return reflect.ValueOf(v1).String() == reflect.ValueOf(v2).String()
	}

	// 布尔值
	if k1 == reflect.Bool {
		if k2 != reflect.Bool {
			return false
		}
		return reflect.ValueOf(v1).Bool() == reflect.ValueOf(v2).Bool()
	}
	// 数组
	// TODO: 复杂类型比较实现
	if k1 == reflect.Array {
		if k2 != reflect.Array {
			return false
		}
	}
	return false
}

func Assert(t *testing.T, trueVal, outVal interface{}, testCaseName string) {
	if !DeepEqual(trueVal, outVal) {
		t.Errorf("--- test error:\r\n\tassert function: %s\r\n\texpected value: %+v\r\n\tactual: %+v", testCaseName, trueVal, outVal)
	}
}

func AssertNil(t *testing.T, val interface{}, testCaseName string) {
	if val != nil {
		t.Errorf("--- test error:\r\n\tassert function: %s\r\n\texpected value: nil\r\n\tactual: %+v", testCaseName,val)
	}
}