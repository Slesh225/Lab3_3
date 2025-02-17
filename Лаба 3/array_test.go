package main

import (
	"bufio"
	"os"
	"testing"
)

func TestNewArray(t *testing.T) {
	capacity := 5
	arr := NewArray(capacity)
	if arr.maxCapacity != capacity || arr.size != 0 || len(arr.data) != capacity {
		t.Errorf("NewArray() = %v; want array with capacity %d and size 0", arr, capacity)
	}
}

func TestArrayAdd(t *testing.T) {
	arr := NewArray(5)
	arr.Add(0, "first")
	if arr.size != 1 || arr.data[0] != "first" {
		t.Errorf("Add() = %v; want array with one element 'first'", arr)
	}

	arr.Add(1, "second")
	if arr.size != 2 || arr.data[1] != "second" {
		t.Errorf("Add() = %v; want array with two elements ['first', 'second']", arr)
	}

	arr.Add(1, "middle")
	if arr.size != 3 || arr.data[1] != "middle" || arr.data[2] != "second" {
		t.Errorf("Add() = %v; want array with three elements ['first', 'middle', 'second']", arr)
	}

	// Попытка добавить элемент за пределами допустимого индекса
	arr.Add(5, "invalid")
	if arr.size != 3 {
		t.Errorf("Add() with invalid index = %v; want array size 3", arr)
	}

	// Попытка добавить элемент в заполненный массив
	arr.Add(3, "third")
	arr.Add(4, "fourth")
	arr.Add(5, "fifth")
	if arr.size != 5 {
		t.Errorf("Add() to full array = %v; want array size 5", arr)
	}
}

func TestArrayAddToTheEnd(t *testing.T) {
	arr := NewArray(3)
	arr.AddToTheEnd("first")
	if arr.size != 1 || arr.data[0] != "first" {
		t.Errorf("AddToTheEnd() = %v; want array with one element 'first'", arr)
	}

	arr.AddToTheEnd("second")
	if arr.size != 2 || arr.data[1] != "second" {
		t.Errorf("AddToTheEnd() = %v; want array with two elements ['first', 'second']", arr)
	}

	// Попытка добавить элемент в заполненный массив
	arr.AddToTheEnd("third")
	arr.AddToTheEnd("fourth")
	if arr.size != 3 {
		t.Errorf("AddToTheEnd() to full array = %v; want array size 3", arr)
	}
}

func TestArrayRemove(t *testing.T) {
	arr := NewArray(5)
	arr.AddToTheEnd("first")
	arr.AddToTheEnd("second")
	arr.AddToTheEnd("third")

	arr.Remove(1)
	if arr.size != 2 || arr.data[0] != "first" || arr.data[1] != "third" {
		t.Errorf("Remove() = %v; want array with two elements ['first', 'third']", arr)
	}

	// Попытка удалить элемент за пределами допустимого индекса
	arr.Remove(2)
	if arr.size != 2 {
		t.Errorf("Remove() with invalid index = %v; want array size 2", arr)
	}

	arr.Remove(0)
	arr.Remove(0)
	if arr.size != 0 {
		t.Errorf("Remove() all elements = %v; want empty array", arr)
	}

	// Попытка удалить элемент из пустого массива
	arr.Remove(0)
	if arr.size != 0 {
		t.Errorf("Remove() from empty array = %v; want empty array", arr)
	}
}

func TestArrayReplace(t *testing.T) {
	arr := NewArray(3)
	arr.AddToTheEnd("first")
	arr.AddToTheEnd("second")

	arr.Replace(1, "new_second")
	if arr.data[1] != "new_second" {
		t.Errorf("Replace() = %v; want array with elements ['first', 'new_second']", arr)
	}

	// Попытка заменить элемент за пределами допустимого индекса
	arr.Replace(2, "invalid")
	if arr.size != 2 {
		t.Errorf("Replace() with invalid index = %v; want array size 2", arr)
	}
}

func TestArrayPrint(t *testing.T) {
	arr := NewArray(3)
	arr.AddToTheEnd("first")
	arr.AddToTheEnd("second")
	// Здесь можно использовать захват вывода для проверки, но для простоты просто вызовем
	arr.Print()
}

func TestArrayLength(t *testing.T) {
	arr := NewArray(5)
	if arr.Length() != 0 {
		t.Errorf("Length() = %v; want 0", arr.Length())
	}

	arr.AddToTheEnd("first")
	if arr.Length() != 1 {
		t.Errorf("Length() = %v; want 1", arr.Length())
	}
}

func TestArraySaveToFile(t *testing.T) {
	arr := NewArray(3)
	arr.AddToTheEnd("first")
	arr.AddToTheEnd("second")

	filename := "test_array.txt"
	err := arr.SaveToFile(filename)
	if err != nil {
		t.Errorf("SaveToFile() error = %v; want nil", err)
	}
	defer os.Remove(filename) // Очистка после теста

	// Проверка содержимого файла
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 2 || lines[0] != "first" || lines[1] != "second" {
		t.Errorf("File content = %v; want ['first', 'second']", lines)
	}
}

func TestArrayLoadFromFile(t *testing.T) {
	arr := NewArray(3)
	filename := "test_load_array.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.WriteString("first\nsecond\n")
	file.Close()
	defer os.Remove(filename)

	err = arr.LoadFromFile(filename)
	if err != nil {
		t.Errorf("LoadFromFile() error = %v; want nil", err)
	}

	if arr.size != 2 || arr.data[0] != "first" || arr.data[1] != "second" {
		t.Errorf("LoadFromFile() = %v; want array with two elements ['first', 'second']", arr)
	}
}

func TestArrayGet(t *testing.T) {
	arr := NewArray(3)
	arr.AddToTheEnd("first")
	arr.AddToTheEnd("second")

	if arr.Get(0) != "first" {
		t.Errorf("Get() = %v; want 'first'", arr.Get(0))
	}

	if arr.Get(1) != "second" {
		t.Errorf("Get() = %v; want 'second'", arr.Get(1))
	}

	// Попытка получить элемент за пределами допустимого индекса
	if arr.Get(2) != "" {
		t.Errorf("Get() with invalid index = %v; want empty string", arr.Get(2))
	}
}

func TestArrayEquals(t *testing.T) {
	arr1 := NewArray(3)
	arr1.AddToTheEnd("first")
	arr1.AddToTheEnd("second")

	arr2 := NewArray(3)
	arr2.AddToTheEnd("first")
	arr2.AddToTheEnd("second")

	if !arr1.Equals(arr2) {
		t.Errorf("Equals() = %v; want true", arr1.Equals(arr2))
	}

	arr2.Replace(1, "new_second")
	if arr1.Equals(arr2) {
		t.Errorf("Equals() = %v; want false", arr1.Equals(arr2))
	}

	// Сравнение с nil
	if arr1.Equals(nil) {
		t.Errorf("Equals() with nil = %v; want false", arr1.Equals(nil))
	}
}

func TestArrayDeserializeTextarray(t *testing.T) {
	arr := NewArray(3)
	data := `["first","second"]`

	err := arr.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if arr.size != 2 || arr.data[0] != "first" || arr.data[1] != "second" {
		t.Errorf("DeserializeText() = %v; want array with two elements ['first', 'second']", arr)
	}
}

func TestArrayDeserializeBinary(t *testing.T) {
	arr := NewArray(3)
	data := []byte{
		5, 0, 0, 0, // длина "first"
		'f', 'i', 'r', 's', 't', // строка "first"
		6, 0, 0, 0, // длина "second"
		's', 'e', 'c', 'o', 'n', 'd', // строка "second"
	}

	err := arr.DeserializeBinary(data)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if arr.size != 2 || arr.data[0] != "first" || arr.data[1] != "second" {
		t.Errorf("DeserializeBinary() = %v; want array with two elements ['first', 'second']", arr)
	}
}

func TestArrayDeserializeBinaryError(t *testing.T) {
	arr := NewArray(3)
	data := []byte{
		5, 0, 0, 0, // длина "first"
		'f', 'i', 'r', 's', // не хватает одного байта для строки "first"
	}

	err := arr.DeserializeBinary(data)
	if err == nil {
		t.Errorf("DeserializeBinary() error = nil; want error")
	}
}

func TestArraySerializeTextEmpty(t *testing.T) {
	arr := NewArray(3)

	serialized, err := arr.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected := `[]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}
}

func TestArrayDeserializeTextEmpty(t *testing.T) {
	arr := NewArray(3)
	data := `[]`

	err := arr.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if arr.size != 0 {
		t.Errorf("DeserializeText() = %v; want empty array", arr)
	}
}
func TestArraySerializeText(t *testing.T) {
	arr := NewArray(3)
	arr.AddToTheEnd("first")
	arr.AddToTheEnd("second")

	serialized, err := arr.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected := `["first","second"]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}

	// Десериализация
	newArr := NewArray(3)
	err = newArr.DeserializeText(serialized)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if newArr.size != 2 || newArr.data[0] != "first" || newArr.data[1] != "second" {
		t.Errorf("DeserializeText() = %v; want array with two elements ['first', 'second']", newArr)
	}
}
func TestArraySerializeBinary(t *testing.T) {
	arr := NewArray(3)
	arr.AddToTheEnd("first")
	arr.AddToTheEnd("second")

	serialized, err := arr.SerializeBinary()
	if err != nil {
		t.Errorf("SerializeBinary() error = %v; want nil", err)
	}

	// Ожидаемый результат: длина строки "first" (5) + строка "first" + длина строки "second" (6) + строка "second"
	expected := []byte{
		5, 0, 0, 0, // длина "first"
		'f', 'i', 'r', 's', 't', // строка "first"
		6, 0, 0, 0, // длина "second"
		's', 'e', 'c', 'o', 'n', 'd', // строка "second"
	}

	if len(serialized) != len(expected) {
		t.Errorf("SerializeBinary() length = %v; want %v", len(serialized), len(expected))
	}

	for i := range expected {
		if serialized[i] != expected[i] {
			t.Errorf("SerializeBinary() = %v; want %v", serialized, expected)
			break
		}
	}

	// Десериализация
	newArr := NewArray(3)
	err = newArr.DeserializeBinary(serialized)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if newArr.size != 2 || newArr.data[0] != "first" || newArr.data[1] != "second" {
		t.Errorf("DeserializeBinary() = %v; want array with two elements ['first', 'second']", newArr)
	}
}
