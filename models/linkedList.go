package models

import (
	"fmt"
	"strconv"
)

type outOfRange struct {
	len uint64
	idx int
}

func (err outOfRange) Error() string {
	return fmt.Sprint("index ", strconv.Itoa(err.idx), " is out of range of a list with len ", strconv.Itoa(int(err.len)))
}

// LinkedList Структура двухсвязного списка для вызова методов и.т.п.
type LinkedList[T comparable] struct {
	head *node[T]
	tail *node[T]
}

// node Структура узла двухсвязного списка (не доступна для вызова вне пакета).
type node[T comparable] struct {
	data T
	next *node[T]
	prev *node[T]
}

// New Создание списка данного типа с данной длиной (заполняется стандартными значениями типа)
func New[T comparable](len uint64) LinkedList[T] {
	var res LinkedList[T]
	for i := uint64(0); i < len; i++ {
		var defaultVal T
		res.Add(defaultVal)
	}
	return res
}

// NewFromSlice Создание списка из слайса того же типа
func NewFromSlice[T comparable](slice []T) LinkedList[T] {
	var res LinkedList[T]
	for i := 0; i < len(slice); i++ {
		res.Add(slice[i])
	}
	return res
}

// Add Добавление нового узла с данным значением в конец списка
func (l *LinkedList[T]) Add(val T) {
	newNode := node[T]{data: val, next: nil}
	if l.tail != nil {
		newNode.prev = l.tail
		l.tail.next = &newNode
		l.tail = &newNode
	} else {
		l.head = &newNode
		l.tail = &newNode
	}
}

// Pop Удаление узла из конца списка
func (l *LinkedList[T]) Pop() error {
	if l.tail == nil {
		return outOfRange{len: 0, idx: 0}
	}
	newTail := l.tail.prev
	newTail.next = nil
	l.tail = newTail
	return nil
}

// At Получение значение из узла с данным индексом
func (l *LinkedList[T]) At(idx int) (T, error) {
	ptr := l.head
	for i := 0; i < idx; i++ {
		if ptr.next != nil {
			ptr = ptr.next
		} else {
			var defaultValue T
			return defaultValue, outOfRange{len: l.Size(), idx: idx}
		}
	}
	return ptr.data, nil
}

// Size Получение размера списка
func (l *LinkedList[T]) Size() uint64 {
	var cnt uint64
	for ptr := l.head; ptr != nil; ptr = ptr.next {
		cnt++
	}
	return cnt
}

// DeleteFrom Удаление узла данного индекса из списка
func (l *LinkedList[T]) DeleteFrom(idx int) error {
	ptr := l.head
	for i := 0; i < idx; i++ {
		ptr = ptr.next
		if ptr == nil {
			return outOfRange{len: l.Size(), idx: idx}
		}
	}
	if ptr.prev != nil {
		ptr.prev.next = ptr.next
	} else {
		l.head = ptr.next
	}
	if ptr.next != nil {
		ptr.next.prev = ptr.prev
	} else {
		l.tail = ptr.prev
	}
	return nil
}

// UpdateAt изменение значения узла с данным индексом
func (l *LinkedList[T]) UpdateAt(idx int, val T) error {
	ptr := l.head
	for i := 0; i < idx; i++ {
		ptr = ptr.next
		if ptr == nil {
			return outOfRange{len: l.Size(), idx: idx}
		}
	}
	ptr.data = val
	return nil
}

// InsertAt Вставить узел с данным значением по данному индексу (т.е. он будет доступен по данном индексу)
func (l *LinkedList[T]) InsertAt(idx int, val T) error {
	ptr := l.head
	newNode := node[T]{data: val}
	if idx == 0 {
		newNode.next = l.head
		l.head = &newNode
		return nil
	}
	for i := 1; i < idx; i++ {
		if ptr.next != nil {
			ptr = ptr.next
		} else {
			return outOfRange{len: l.Size(), idx: idx}
		}
	}
	newNode.next = ptr.next
	newNode.prev = ptr
	if ptr.next != nil {
		ptr.next.prev = &newNode
		ptr.next = &newNode
		return nil
	}
	l.tail.next = &newNode
	l.tail = &newNode
	return nil
}

// Println вывод связанного списка в терминал
func (l *LinkedList[T]) Println() {
	ptr := l.head
	fmt.Print("[")
	if ptr == nil {
		fmt.Print("]\n")
		return
	}
	for ptr.next != nil {
		fmt.Print(ptr.data, " -> ")
		ptr = ptr.next
	}
	fmt.Print(ptr.data, "]\n")
}
