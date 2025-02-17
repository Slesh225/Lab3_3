package main

import (
	"bufio"
	"os"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList()
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("NewDoublyLinkedList() = %v; want empty list", dll)
	}
}

func TestAddToHead(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToHead("first")
	if dll.Head.Data != "first" || dll.Tail.Data != "first" {
		t.Errorf("AddToHead() = %v; want list with one element 'first'", dll)
	}

	dll.AddToHead("second")
	if dll.Head.Data != "second" || dll.Tail.Data != "first" {
		t.Errorf("AddToHead() = %v; want list with elements ['second', 'first']", dll)
	}
}

func TestAddToTail(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("first")
	if dll.Head.Data != "first" || dll.Tail.Data != "first" {
		t.Errorf("AddToTail() = %v; want list with one element 'first'", dll)
	}

	dll.AddToTail("second")
	if dll.Head.Data != "first" || dll.Tail.Data != "second" {
		t.Errorf("AddToTail() = %v; want list with elements ['first', 'second']", dll)
	}
}

func TestRemoveFromHead(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToHead("first")
	dll.AddToHead("second")

	dll.RemoveFromHead()
	if dll.Head.Data != "first" || dll.Tail.Data != "first" {
		t.Errorf("RemoveFromHead() = %v; want list with one element 'first'", dll)
	}

	dll.RemoveFromHead()
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("RemoveFromHead() = %v; want empty list", dll)
	}

	dll.RemoveFromHead() // Попытка удалить из пустого списка
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("RemoveFromHead() from empty list = %v; want empty list", dll)
	}
}

func TestRemoveFromTail(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("first")
	dll.AddToTail("second")

	dll.RemoveFromTail()
	if dll.Head.Data != "first" || dll.Tail.Data != "first" {
		t.Errorf("RemoveFromTail() = %v; want list with one element 'first'", dll)
	}

	dll.RemoveFromTail()
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("RemoveFromTail() = %v; want empty list", dll)
	}

	dll.RemoveFromTail() // Попытка удалить из пустого списка
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("RemoveFromTail() from empty list = %v; want empty list", dll)
	}
}

func TestRemoveByValue(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("first")
	dll.AddToTail("second")
	dll.AddToTail("third")

	dll.RemoveByValue("second")
	if dll.Head.Data != "first" || dll.Tail.Data != "third" {
		t.Errorf("RemoveByValue() = %v; want list with elements ['first', 'third']", dll)
	}

	dll.RemoveByValue("first")
	if dll.Head.Data != "third" || dll.Tail.Data != "third" {
		t.Errorf("RemoveByValue() = %v; want list with one element 'third'", dll)
	}

	dll.RemoveByValue("third")
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("RemoveByValue() = %v; want empty list", dll)
	}

	dll.RemoveByValue("nonexistent") // Попытка удалить несуществующее значение
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("RemoveByValue() with nonexistent value = %v; want empty list", dll)
	}
}

func TestSearch(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("first")
	dll.AddToTail("second")

	node := dll.Search("second")
	if node == nil || node.Data != "second" {
		t.Errorf("Search() = %v; want node with data 'second'", node)
	}

	node = dll.Search("nonexistent")
	if node != nil {
		t.Errorf("Search() = %v; want nil", node)
	}
}

func TestPrint(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("first")
	dll.AddToTail("second")
	// Здесь можно использовать захват вывода для проверки, но для простоты просто вызовем
	dll.Print()
}

func TestSaveToFile_dll(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("first")
	dll.AddToTail("second")

	filename := "test_dll.txt"
	err := dll.SaveToFile(filename)
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

func TestLoadFromFile_dll(t *testing.T) {
	dll := NewDoublyLinkedList()
	filename := "test_load_dll.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.WriteString("first\nsecond\n")
	file.Close()
	defer os.Remove(filename)

	err = dll.LoadFromFile(filename)
	if err != nil {
		t.Errorf("LoadFromFile() error = %v; want nil", err)
	}

	if dll.Head.Data != "first" || dll.Tail.Data != "second" {
		t.Errorf("LoadFromFile() = %v; want list with elements ['first', 'second']", dll)
	}
}

func TestFindNodeByValue(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("first")
	dll.AddToTail("second")

	node := dll.findNodeByValue("second")
	if node == nil || node.Data != "second" {
		t.Errorf("findNodeByValue() = %v; want node with data 'second'", node)
	}

	node = dll.findNodeByValue("nonexistent")
	if node != nil {
		t.Errorf("findNodeByValue() = %v; want nil", node)
	}
}

// TestSerializeDeserializeText проверяет сериализацию и десериализацию в текстовом формате (JSON)
func TestSerializeDeserializeTextdll(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("первый")
	dll.AddToTail("второй")
	dll.AddToTail("третий")

	// Сериализация в текстовый формат
	serialized, err := dll.SerializeText()
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Десериализация из текстового формата
	newDll := NewDoublyLinkedList()
	err = newDll.DeserializeText(serialized)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Проверка, что данные совпадают
	currentOriginal := dll.Head
	currentNew := newDll.Head
	for currentOriginal != nil && currentNew != nil {
		if currentOriginal.Data != currentNew.Data {
			t.Errorf("Ожидалось %s, получено %s", currentOriginal.Data, currentNew.Data)
		}
		currentOriginal = currentOriginal.Next
		currentNew = currentNew.Next
	}

	if currentOriginal != nil || currentNew != nil {
		t.Error("Длины списков не совпадают после десериализации")
	}
}

// TestSerializeDeserializeBinary проверяет сериализацию и десериализацию в бинарном формате
func TestSerializeDeserializeBinarydll(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.AddToTail("первый")
	dll.AddToTail("второй")
	dll.AddToTail("третий")

	// Сериализация в бинарный формат
	serialized, err := dll.SerializeBinary()
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Десериализация из бинарного формата
	newDll := NewDoublyLinkedList()
	err = newDll.DeserializeBinary(serialized)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Проверка, что данные совпадают
	currentOriginal := dll.Head
	currentNew := newDll.Head
	for currentOriginal != nil && currentNew != nil {
		if currentOriginal.Data != currentNew.Data {
			t.Errorf("Ожидалось %s, получено %s", currentOriginal.Data, currentNew.Data)
		}
		currentOriginal = currentOriginal.Next
		currentNew = currentNew.Next
	}

	if currentOriginal != nil || currentNew != nil {
		t.Error("Длины списков не совпадают после десериализации")
	}
}
