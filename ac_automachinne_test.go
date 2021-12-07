package lad

import (
	"testing"
)

func TestAcMachine_Add(t *testing.T) {
	machine := New()
	machine.Add("我们")
	machine.Add("历史")
	machine.Add("listen to me")
	p := machine.root

	p.view()
}

func TestAcMachine_Load(t *testing.T) {
	machine := New()
	if err := machine.Load("./test.data"); err != nil {
		t.Error(err)
	}

	machine.root.view()
}

func TestAcMachine_Build(t *testing.T) {
	machine := New()
	if err := machine.Load("./test.data"); err != nil {
		t.Error(err)
	}
	machine.Build()
	machine.root.view()
}

func TestAcMachine_Find(t *testing.T) {
	machine := New()
	if err := machine.Load("./test.data"); err != nil {
		t.Error(err)
	}
	machine.Build()

	t.Log(machine.Find("ab       cdadfadfadfadf"))
}

func TestAcMachine_Match(t *testing.T) {
	machine := New()
	if err := machine.Load("./test.data"); err != nil {
		t.Error(err)
	}
	machine.Build()

	t.Log(machine.Match("abx       cdadfadfadfadf"))
}

func TestAcMachine_Replace(t *testing.T) {
	machine := New()
	if err := machine.Load("./test.data"); err != nil {
		t.Error(err)
	}
	machine.Build()

	t.Log(machine.Replace("ab       cdadfadfadfadf", "****"))
}
