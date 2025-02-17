package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

// DoublyNode представляет узел в двусвязном списке
type DoublyNode struct {
	Data string
	Next *DoublyNode
	Prev *DoublyNode
}

// DoublyLinkedList представляет структуру двусвязного списка
type DoublyLinkedList struct {
	Head *DoublyNode // Public
	Tail *DoublyNode // Public
}

// NewDoublyLinkedList создает новый двусвязный список (Public)
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// AddToHead добавляет новый узел в начало списка (Public)
func (dll *DoublyLinkedList) AddToHead(value string) {
	newNode := &DoublyNode{Data: value, Next: dll.Head}
	if dll.Head != nil {
		dll.Head.Prev = newNode
	} else {
		dll.Tail = newNode
	}
	dll.Head = newNode
}

// AddToTail добавляет новый узел в конец списка (Public)
func (dll *DoublyLinkedList) AddToTail(value string) {
	newNode := &DoublyNode{Data: value, Prev: dll.Tail}
	if dll.Tail != nil {
		dll.Tail.Next = newNode
	} else {
		dll.Head = newNode
	}
	dll.Tail = newNode
}

// RemoveFromHead удаляет узел из начала списка (Public)
func (dll *DoublyLinkedList) RemoveFromHead() {
	if dll.Head == nil {
		return
	}
	temp := dll.Head
	dll.Head = dll.Head.Next
	if dll.Head != nil {
		dll.Head.Prev = nil
	} else {
		dll.Tail = nil
	}
	temp.Next = nil
}

// RemoveFromTail удаляет узел из конца списка (Public)
func (dll *DoublyLinkedList) RemoveFromTail() {
	if dll.Tail == nil {
		return
	}
	temp := dll.Tail
	dll.Tail = dll.Tail.Prev
	if dll.Tail != nil {
		dll.Tail.Next = nil
	} else {
		dll.Head = nil
	}
	temp.Prev = nil
}

// RemoveByValue удаляет первый узел с указанным значением из списка (Public)
func (dll *DoublyLinkedList) RemoveByValue(value string) {
	current := dll.Head
	for current != nil {
		if current.Data == value {
			if current == dll.Head {
				dll.RemoveFromHead()
			} else if current == dll.Tail {
				dll.RemoveFromTail()
			} else {
				current.Prev.Next = current.Next
				current.Next.Prev = current.Prev
			}
			return
		}
		current = current.Next
	}
}

// Search ищет узел с указанным значением в списке (Public)
func (dll *DoublyLinkedList) Search(value string) *DoublyNode {
	current := dll.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Print выводит элементы списка (Public)
func (dll *DoublyLinkedList) Print() {
	current := dll.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}

// SaveToFile сохраняет список в файл (Public)
func (dll *DoublyLinkedList) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	current := dll.Head
	for current != nil {
		_, err := file.WriteString(current.Data + "\n")
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}
		current = current.Next
	}
	return nil
}

// LoadFromFile загружает список из файла (Public)
func (dll *DoublyLinkedList) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dll.AddToTail(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// findNodeByValue вспомогательная функция для поиска узла по значению (Private)
func (dll *DoublyLinkedList) findNodeByValue(value string) *DoublyNode {
	current := dll.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// SerializeText сериализует двусвязный список в текстовый формат (JSON)
func (dll *DoublyLinkedList) SerializeText() (string, error) {
	var data []string
	current := dll.Head
	for current != nil {
		data = append(data, current.Data)
		current = current.Next
	}
	result, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации в текстовый формат: %v", err)
	}
	return string(result), nil
}

// DeserializeText десериализует двусвязный список из текстового формата (JSON)
func (dll *DoublyLinkedList) DeserializeText(data string) error {
	var temp []string
	err := json.Unmarshal([]byte(data), &temp)
	if err != nil {
		return fmt.Errorf("ошибка десериализации из текстового формата: %v", err)
	}
	dll.Head = nil
	dll.Tail = nil
	for _, value := range temp {
		dll.AddToTail(value)
	}
	return nil
}

// SerializeBinary сериализует двусвязный список в бинарный формат
func (dll *DoublyLinkedList) SerializeBinary() ([]byte, error) {
	var result []byte
	current := dll.Head
	for current != nil {
		// Записываем длину строки
		length := uint32(len(current.Data))
		lengthBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBytes, length)
		result = append(result, lengthBytes...)

		// Записываем строку
		result = append(result, []byte(current.Data)...)
		current = current.Next
	}
	return result, nil
}

// DeserializeBinary десериализует двусвязный список из бинарного формата
func (dll *DoublyLinkedList) DeserializeBinary(data []byte) error {
	dll.Head = nil
	dll.Tail = nil
	offset := 0
	for offset < len(data) {
		// Читаем длину строки
		length := binary.LittleEndian.Uint32(data[offset : offset+4])
		offset += 4

		// Читаем строку
		strData := string(data[offset : offset+int(length)])
		offset += int(length)
		dll.AddToTail(strData)
	}
	return nil
}
