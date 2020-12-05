package golinkedlist

import "fmt"

type LinkedList struct{
	Value interface{}
	Next *LinkedList
}

type EqualityComparatorFunc func(interface{}, interface{}) bool

func (head *LinkedList) AddValue(v interface{}) *LinkedList{
	return &LinkedList{
		Value: v,
		Next: head,
	}
}

func defaultEqualityComparator(a interface{}, b interface{}) bool {
	return a == b
}


func (head *LinkedList) FindValue(value interface{},
	equalityComparator EqualityComparatorFunc) *LinkedList{
	tmp := head
	eqCmp := equalityComparator
	if equalityComparator == nil {
		eqCmp = defaultEqualityComparator
	}

	for tmp != nil {
		if eqCmp(tmp.Value, value) {
			return tmp
		}

		tmp = tmp.Next
	}

	return nil
}

func (head *LinkedList) Print() string{
	str := ""
	tmp := head
	for tmp != nil {
		str += fmt.Sprintf("%v - ", tmp.Value)
		tmp = tmp.Next
	}

	return str
}