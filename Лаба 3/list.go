package main

import "errors"

type List struct {
	elements []string
}

func NewList() *List {
	return &List{}
}

func (l *List) Push(value string) {
	l.elements = append(l.elements, value)
}

func (l *List) Delete() {
	if len(l.elements) > 0 {
		l.elements = l.elements[1:]
	}
}

func (l *List) Get() (string, error) {
	if len(l.elements) == 0 {
		return "", errors.New("список пуст")
	}
	return l.elements[0], nil
}
