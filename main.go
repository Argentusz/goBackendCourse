package main

import (
	"fmt"
	"goBackendCourse/models"
)

func main() {
	ll := models.New[uint64](6)
	ll.Println()
	err := ll.Pop()
	if err != nil {
		fmt.Println(err.Error())
	}
	ll.Add(12)
	ll.Println()
	get, err := ll.At(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(get)
	get, err = ll.At(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(get)
	fmt.Println(ll.Size())
	sz := models.New[uint64](0)
	fmt.Println(sz.Size())

	err = ll.DeleteFrom(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ll.DeleteFrom(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = ll.DeleteFrom(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	ll.Println()

	err = ll.UpdateAt(1, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	ll.Println()

	ls := models.NewFromSlice([]int8{0, 1, 2, 4, 5, 3})
	ls.Println()

	ls.InsertAt(0, 1)
	ls.Println()
	ls.InsertAt(2, 100)
	ls.Println()
	ls.InsertAt(int(ls.Size()), 100)
	ls.Println()
}
