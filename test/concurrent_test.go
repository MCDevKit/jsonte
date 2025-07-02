package test

import (
	"sync"
	"testing"

	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/types"
)

func TestConcurrentProcess(t *testing.T) {
	template := `{"$template":{"value":"{{=1+1}}"}}`
	expected, err := types.ParseJsonObject([]byte(`{"value":2}`))
	if err != nil {
		t.Fatal(err)
	}
	modules := map[string]jsonte.JsonModule{}
	scope := types.NewJsonObject()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := jsonte.Process("test", template, scope, modules, -1)
			if err != nil {
				t.Error(err)
				return
			}
			obj := res.Get("test").(*types.JsonObject)
			compareJsonObject(t, expected, obj, "", true)
		}()
	}
	wg.Wait()
}

func TestConcurrentMCFunction(t *testing.T) {
	scope := types.NewJsonObject()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out, err := jsonte.ProcessMCFunction("say #{1+1}", scope)
			if err != nil {
				t.Error(err)
				return
			}
			if out != "say 2" {
				t.Errorf("expected 'say 2', got %s", out)
			}
		}()
	}
	wg.Wait()
}

func TestConcurrentLang(t *testing.T) {
	scope := types.NewJsonObject()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out, err := jsonte.ProcessLangFile("foo=##{1+1}", scope)
			if err != nil {
				t.Error(err)
				return
			}
			if out != "foo=2" {
				t.Errorf("expected 'foo=2', got %s", out)
			}
		}()
	}
	wg.Wait()
}
