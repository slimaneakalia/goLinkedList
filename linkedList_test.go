package golinkedlist_test

import (
	"testing"
	"github.com/slimaneakalia/golinkedlist"
)

func TestAddValue(t *testing.T){
	listHead := &golinkedlist.LinkedList{ Value: "head 1" }
	listHead = listHead.AddValue("head 2")

	if listHead.Value != "head 2"{
		t.Errorf("AddValue not working, current value %v", listHead.Value)
	} else {
		t.Log("AddValue is working")
	}

}

func TestFindValue(t *testing.T){
	listHead := &golinkedlist.LinkedList{ Value: "head 1" }
	listHead = listHead.AddValue("head 2")
	listHead = listHead.AddValue("head 3")

	researchedValue := listHead.FindValue("head 3", nil)

	if researchedValue == nil {
		t.Error("FindValue is not working")
	} else {
		t.Log("FindValue is working")
	}
}

func TestPrint(t *testing.T){
	listHead := &golinkedlist.LinkedList{ Value: "head 1" }
	listHead = listHead.AddValue("head 2")
	listHead = listHead.AddValue("head 3")

	str := listHead.Print()

	if str != "head 3 - head 2 - head 1 - "{
		t.Errorf("Print is not working, current output: %v", str)
	}else {
		t.Log("Print is working")
	}
}