package gdsutil

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {

}
func TestQueue_EnQueue(t *testing.T) {
	q := NewQueue()
	p1 := person{"song"}
	q.EnQueue(p1)
	p2 := person{"wang"}
	q.EnQueue(p2)
	p3 := person{"Yii"}
	q.EnQueue(p3)

	//fmt.Println(q.HeadNode())

	for !q.IsEmpty() {
		node := q.DeQueue()
		fmt.Println(node)
	}
}
