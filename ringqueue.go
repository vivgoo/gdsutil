package gdsutil

type STATUS int

const (
	SUCCESS      STATUS = 1
	FAIL      STATUS = 0
)

type RingQueue struct {
	queueData []interface{}
	maxSize int
	head int
	tail int
}

func InitRingQueue(maxSize int) (rq RingQueue, status STATUS) {
	if maxSize <= 0  {return RingQueue{},FAIL}
	return RingQueue{
		queueData: make([]interface{},maxSize),
		maxSize: maxSize,
		head: maxSize - 1,
		tail: maxSize - 1,
	},SUCCESS

	return RingQueue{},FAIL
}

func (rq *RingQueue) IsEmpty() bool  {
	return rq.head == rq.tail
}

func (rq *RingQueue) IsFull() bool  {
	return (rq.tail + 1) % rq.maxSize == rq.head
}

func (rq *RingQueue) EnQueue(data interface{})  STATUS {
	if (rq.tail + 1) % rq.maxSize == rq.head {
		return FAIL
	}
	rq.tail = (rq.tail + 1) % rq.maxSize
	rq.queueData[rq.tail] = data

	return SUCCESS
}

func (rq *RingQueue) HeadData() (data interface{},status STATUS)  {
	var tempData interface{}
	if rq.head == rq.tail {
		return tempData,FAIL
	}
	return rq.queueData[(rq.head+1) % rq.maxSize],SUCCESS
}

func (rq *RingQueue) DeQueue() (data interface{},status STATUS) {
	var tempData interface{}
	if rq.head == rq.tail {
		return tempData,FAIL
	}

	rq.head = (rq.head + 1) % rq.maxSize
	tempData = rq.queueData[rq.head]
	return tempData,SUCCESS
}