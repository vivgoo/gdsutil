package gdsutil

// 单向链表中的结点信息
type LinkNode struct {
	data interface{}
	next *LinkNode
}

// 单向链表
type LinkList struct {
	head  *LinkNode
	count int
}

// 创建链表
func NewLinkList() *LinkList {
	return &LinkList{
		nil,
		0,
	}
}

// 判断是否为空
func (l *LinkList) IsEmpty() bool {
	return 0 == l.count
}

// 获取链表的长度
func (l *LinkList) Length() int {
	return l.count
}

// 在头部添加数据,把原来的头部替换掉,还需要把现在的元素指向原来的头部
func (l *LinkList) Add(data interface{}) {
	node := &LinkNode{
		data,
		nil,
	}
	node.next = l.head
	l.head = node
	l.count++
}

// 查找指定序号的结点，如果找到返回其指针
func (l *LinkList) findNode(index int) *LinkNode {
	p := l.head
	i := 0
	for p != nil && i < index {
		p = p.next
		i++
	}
	if p != nil && i == index {
		return p
	} else {
		return nil
	}
}

// 在指定位置插入结点
func (l *LinkList) Insert(data interface{}, index int) bool {
	if nil == l.head {
		return false
	}

	q := &LinkNode{
		data, nil,
	}

	if 0 == index {
		q.next = l.head
		l.head = q
		l.count++
		return true
	}

	var p *LinkNode
	p = l.findNode(index - 1)
	if nil == p {
		return false
	}
	q.next = p.next
	p.next = q
	l.count++
	return true
}

func (l *LinkList) Delete(data interface{}) {
	if nil == l.head {
		return
	}
	node := l.head
	if node.data == data {
		l.head = node.next
		l.count--
		return
	}

	for node.next != nil {
		if node.next.data == data {
			node.next = node.next.next
			l.count--
		} else {
			node = node.next
		}
	}
}

func (l *LinkList) PrintList() {
	h := l.head
	for h != nil {
		println("data:", h.data)
		h = h.next
	}
}
