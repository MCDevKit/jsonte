package test

import (
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"go.uber.org/zap"
	"testing"
)

var CacheFS = safeio.CreateFakeFS(map[string]interface{}{}, true)

type TestCall struct {
	Name string
	Args []types.JsonType
}

var TestCalls = map[string][]TestCall{}

func TestMain(m *testing.M) {
	utils.InitLogging(zap.DebugLevel)
	burrito.PrintStackTrace = false
	types.Init()
	functions.Init()
	registerHelperFunctions()
	m.Run()
}

func registerHelperFunctions() {
	const group = "test"
	functions.RegisterGroup(functions.Group{
		Name: group,
	})
	functions.RegisterFunction(functions.JsonFunction{
		Group: group,
		Name:  "testCall",
		Body:  callTestFunction,
	})
}

func callTestFunction(name *types.JsonString, args ...types.JsonType) (*types.JsonNull, error) {
	calls, ok := TestCalls[name.StringValue()]
	if !ok {
		calls = []TestCall{}
	}
	calls = append(calls, TestCall{
		Name: name.StringValue(),
		Args: args,
	})
	TestCalls[name.StringValue()] = calls
	return types.Null, nil
}

func ClearTestCalls() {
	TestCalls = map[string][]TestCall{}
}

func AssertTestFunctionCalledNTimes(t *testing.T, name string, n int) {
	if n < 0 {
		t.Fatalf("Invalid number of calls: %d", n)
	}
	calls, ok := TestCalls[name]
	if !ok {
		if n == 0 {
			return
		}
		t.Fatalf("Function %s not called", name)
	}
	if len(calls) != n {
		t.Fatalf("Function %s called %d times, expected %d", name, len(calls), n)
	}
}

func AssertTestFunctionCalled(t *testing.T, name string) {
	_, ok := TestCalls[name]
	if !ok {
		t.Fatalf("Function %s not called", name)
	}
}

func AssertTestFunctionNotCalled(t *testing.T, name string) {
	AssertTestFunctionCalledNTimes(t, name, 0)
}

func AssertTestFunctionCalledWith(t *testing.T, name string, index int, args ...types.JsonType) {
	if index < 0 {
		t.Fatalf("Invalid index: %d", index)
	}
	calls, ok := TestCalls[name]
	if !ok {
		t.Fatalf("Function %s not called", name)
	}
	if index >= len(calls) {
		t.Fatalf("Function %s called %d times, but index %d requested", name, len(calls), index)
	}
	call := calls[index]
	if len(call.Args) != len(args) {
		t.Fatalf("Function %s called with %d arguments, expected %d", name, len(call.Args), len(args))
	}
	for i, arg := range args {
		if !arg.Equals(call.Args[i]) {
			t.Fatalf("Function %s called with incorrect arguments", name)
		}
	}
}
