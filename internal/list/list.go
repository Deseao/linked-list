package list

//List represents a doubly linked list. The zero value for List is an empty list ready to use.
type List struct {
	head *Element
	tail *Element
}

//New returns an initialized list.
func New() *List {
	return new(List)
}

//Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {
	return l.tail
}

//Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {
	return l.head
}

//Init initializes or clears list l.
func (l *List) Init() *List {
	if l == nil {
		return New()
	}

	l.head = nil
	l.tail = nil
	return l
}

//InsertAfter inserts a new element e with value v immediately after mark and returns e. If mark is not an element of l, the list is not modified. The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	e := &Element{Value: v}
	for i := l.head; i != nil; i = i.Next() {
		if i == mark {
			if i.next == nil {
				//tail scenario, insert and update list tail
				l.tail = e
				i.next = e
				e.prev = i
				return e
			}
			temp := i.next
			i.next = e
			e.prev = i
			e.next = temp
			temp.prev = e
			return e
		}
	}
	return nil
}

func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	e := &Element{Value: v}
	for i := l.head; i != nil; i = i.Next() {
		if i == mark {
			if i.prev == nil {
				//head scenario, insert at front
				l.head = e
				i.prev = e
				e.next = i
				return e
			}
			temp := i.prev
			i.prev = e
			e.next = i
			e.prev = temp
			temp.next = e
			return e
		}
	}
	return nil
}

func (l *List) Len() int {
	count := 0
	for e := l.head; e != nil; e.Next() {
		count++
	}
	return count
}

func (l *List) MoveAfter(e, mark *Element) {
	for i := l.head; i != nil; i = i.Next() {
		if i == e {
			prev := i.prev
			next := i.next
			prev.next = next
			next.prev = prev
			l.InsertAfter(i.Value, mark)
		}
	}
}

func (l *List) MoveBefore(e, mark *Element) {
	for i := l.head; i != nil; i = i.Next() {
		if i == e {
			prev := i.prev
			next := i.next
			prev.next = next
			next.prev = prev
			l.InsertBefore(i.Value, mark)
		}
	}
}

func (l *List) MoveToBack(e *Element) {
	for i := l.head; i != nil; i = i.Next() {
		if i == e {
			prev := i.prev
			next := i.next
			prev.next = next
			next.prev = prev
			l.PushBack(i.Value)
		}
	}
}

func (l *List) MoveToFront(e *Element) {
	for i := l.head; i != nil; i = i.Next() {
		if i == e {
			prev := i.prev
			next := i.next
			prev.next = next
			next.prev = prev
			l.PushFront(i.Value)
		}
	}
}

func (l *List) PushBack(v interface{}) *Element {
	e := &Element{Value: v, prev: l.tail}
	if l.tail != nil {
		l.tail.next = e
	}
	if l.head == nil {
		l.head = e
	}
	l.tail = e
	return e
}

func (l *List) PushBackList(other *List) {
	for i := other.head; i != nil; i = i.Next() {
		l.PushBack(i.Value)
	}
}

func (l *List) PushFront(v interface{}) *Element {
	e := &Element{Value: v, next: l.head}
	if l.head != nil {
		l.head.prev = e
	}
	if l.tail == nil {
		l.tail = e
	}
	l.head = e
	return e
}

func (l *List) PushFrontList(other *List) {
	for i := other.head; i != nil; i = i.Next() {
		l.PushFront(i.Value)
	}
}

func (l *List) Remove(e *Element) interface{} {
	for i := l.head; i != nil; i = i.Next() {
		if i == e {
			prev := i.prev
			next := i.next
			prev.next = next
			next.prev = prev
			return i.Value
		}
	}
	return nil
}
