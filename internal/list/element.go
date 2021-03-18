package list

// Element is an element of a linked list.
type Element struct {
	Value interface{}
	next  *Element
	prev  *Element
}

//Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}

//Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
	return e.prev
}
