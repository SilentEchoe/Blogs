package Cache

type list struct {
	value interface{}
	next  *list
}

func NewList(value interface{}) *list {
	var firstNode = &list{
		value: value,
	}
	firstNode.next = firstNode
	return firstNode
}

func (l *list) Add(value interface{}) {
	var next = &list{
		value: value,
	}
	l.next = next
}
