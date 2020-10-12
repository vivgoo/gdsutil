package gdsutil

import "testing"

func TestNewLinkList(t *testing.T) {
	l := NewLinkList()
	l.Add(1)
	l.Insert(2, 0)
	l.Insert(100, 1)
	l.Insert(100, 3)
	l.Insert(100, 4)
	l.Delete(100)
	l.PrintList()
}
