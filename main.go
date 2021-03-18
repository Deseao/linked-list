package main

import (
	//"container/list"
	"fmt"

	"github.com/Deseao/linked-list/internal/list"
)

func main() {
	l := list.New()
	l.PushFront(1)
	two := l.PushFront(2)
	l.PushFront("three")
	val := l.Remove(two)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Printf("two val: %v\n", val)
	fmt.Printf("List: %#v\n", l)
}
