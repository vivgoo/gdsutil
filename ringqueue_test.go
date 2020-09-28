package gdsutil

import (
	"fmt"
	"testing"
)

func TestInitRingQueue(t *testing.T) {
	_,err := InitRingQueue(0)
	if err == SUCCESS {
		t.Error("fail")
	}

	_,err = InitRingQueue(5)
	if err == FAIL{
		t.Error("fail")
	}
}

type person struct {
	name string
}
func TestRingQueue_EnQueue(t *testing.T) {
	q,_ := InitRingQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	if q.EnQueue(5) == SUCCESS {
		t.Error("队列已满")
	}

}



func TestRingQueue_HeadData(t *testing.T) {
	q,_ := InitRingQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	d,_ := q.HeadData()
	if d != 1 {
		t.Error("应该为1")
	}
}

func TestRingQueue_DeQueue(t *testing.T) {
	q,_ := InitRingQueue(5)
	p1 := person{"James soong"}
	q.EnQueue(p1)
	p2 := person{"Jin Wang"}
	q.EnQueue(p2)
	p3 := person{"Li Jian"}
	q.EnQueue(p3)
	p4 := person{"Xia Wan"}
	q.EnQueue(p4)

	p5 := person{"ChunYan Shen"}
	q.EnQueue(p5)

	for q.IsEmpty() == false {
		t,_ := q.DeQueue()
		fmt.Println(t)
	}
}

func TestRingQueue_IsEmpty(t *testing.T) {
	q,_ := InitRingQueue(5)
	if q.IsEmpty() == false {
		t.Error("应该是空")
	}
}