package gulc

import (
	"testing"
	"reflect"
)


type TestCase struct {
	Input  []interface{}
	Output []interface{}
}

func Assert(t *testing.T, trueVal, outVal interface{}, testCaseName string) {
	if !reflect.DeepEqual(trueVal, outVal) {
		t.Errorf("--- test error:\r\n\tassert function: %s\r\n\texpected value: %+v\r\n\tactual: %+v", testCaseName, trueVal, outVal)
	}
}

func AssertNil(t *testing.T, val interface{}, testCaseName string) {
	if val != nil {
		t.Errorf("--- test error:\r\n\tassert function: %s\r\n\texpected value: nil\r\n\tactual: %+v", testCaseName,val)
	}
}

// UnitTest 全自动化的单元测试
func UnitTest(t *testing.T, testCases []*TestCase, testFunc func(...interface{}) []interface{}) {
	// 基于反射获取函数名字
	funcType := reflect.TypeOf(testFunc)
	funcName := funcType.Name()
	t.Logf("start test %s\n", funcName)
	// 测试每一个测试用例
	for _, testCase := range testCases {
		t.Log("-----------------------------------\n")
		output := testFunc(testCase.Input...)
		if reflect.DeepEqual(output, testCase.Output) {
			t.Errorf("[%s]func output not equal with right output: input=%+v, output=%+v, right output=%+v", funcName, testCase.Input, output, testCase.Output)
		}
	}
}