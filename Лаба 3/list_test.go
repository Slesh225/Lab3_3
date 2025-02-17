package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	list := NewList()
	assert.NotNil(t, list, "Новый список не должен быть nil")
	assert.Equal(t, 0, len(list.elements), "Новый список должен быть пустым")
}

func TestList_Push(t *testing.T) {
	list := NewList()
	list.Push("test1")
	assert.Equal(t, 1, len(list.elements), "Длина списка должна быть 1 после добавления элемента")
	assert.Equal(t, "test1", list.elements[0], "Элемент должен быть 'test1'")

	list.Push("test2")
	assert.Equal(t, 2, len(list.elements), "Длина списка должна быть 2 после добавления второго элемента")
	assert.Equal(t, "test2", list.elements[1], "Элемент должен быть 'test2'")
}

func TestList_Delete(t *testing.T) {
	list := NewList()
	list.Push("test1")
	list.Push("test2")

	list.Delete()
	assert.Equal(t, 1, len(list.elements), "Длина списка должна быть 1 после удаления элемента")
	assert.Equal(t, "test2", list.elements[0], "Элемент должен быть 'test2'")

	list.Delete()
	assert.Equal(t, 0, len(list.elements), "Длина списка должна быть 0 после удаления всех элементов")

	// Удаление из пустого списка
	list.Delete()
	assert.Equal(t, 0, len(list.elements), "Длина списка должна остаться 0 при удалении из пустого списка")
}

func TestList_Get(t *testing.T) {
	list := NewList()

	// Получение из пустого списка
	value, err := list.Get()
	assert.Equal(t, "", value, "Значение должно быть пустой строкой для пустого списка")
	assert.Equal(t, errors.New("список пуст"), err, "Ошибка должна быть 'список пуст'")

	list.Push("test1")
	value, err = list.Get()
	assert.Equal(t, "test1", value, "Значение должно быть 'test1'")
	assert.Nil(t, err, "Ошибка должна быть nil")

	list.Push("test2")
	value, err = list.Get()
	assert.Equal(t, "test1", value, "Значение должно быть 'test1' (первый элемент)")
	assert.Nil(t, err, "Ошибка должна быть nil")
}
