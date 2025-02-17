package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

// Node представляет узел в односвязном списке
type Node struct {
	Data string
	Next *Node
}

// SinglyLinkedList представляет структуру однонаправленного списка
type SinglyLinkedList struct {
	Head *Node // Public
	Size int   // Public
}

// NewSinglyLinkedList создает новый однонаправленный список (Public)
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// AddToHead добавляет новый узел в начало списка (Public)
func (sll *SinglyLinkedList) AddToHead(value string) {
	newNode := &Node{Data: value, Next: sll.Head}
	sll.Head = newNode
	sll.Size++
}

// AddToTail добавляет новый узел в конец списка (Public)
func (sll *SinglyLinkedList) AddToTail(value string) {
	newNode := &Node{Data: value}
	if sll.Head == nil {
		sll.Head = newNode
	} else {
		current := sll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	sll.Size++
}

// RemoveHead удаляет узел из начала списка (Public)
func (sll *SinglyLinkedList) RemoveHead() {
	if sll.Head == nil {
		return
	}
	temp := sll.Head
	sll.Head = sll.Head.Next
	temp.Next = nil
	sll.Size--
}

// RemoveTail удаляет узел из конца списка (Public)
func (sll *SinglyLinkedList) RemoveTail() {
	if sll.Head == nil {
		return
	}
	if sll.Head.Next == nil {
		sll.Head = nil
	} else {
		current := sll.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		current.Next = nil
	}
	sll.Size--
}

// RemoveByValue удаляет первый узел с указанным значением из списка (Public)
func (sll *SinglyLinkedList) RemoveByValue(value string) {
	if sll.Head == nil {
		return
	}
	if sll.Head.Data == value {
		sll.RemoveHead()
		return
	}
	current := sll.Head
	for current.Next != nil {
		if current.Next.Data == value {
			temp := current.Next
			current.Next = temp.Next
			temp.Next = nil
			sll.Size--
			return
		}
		current = current.Next
	}
}

// Search ищет узел с указанным значением в списке (Public)
func (sll *SinglyLinkedList) Search(value string) *Node {
	current := sll.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Print выводит элементы списка (Public)
func (sll *SinglyLinkedList) Print() {
	current := sll.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}

// SaveToFile сохраняет список в файл (Public)
func (sll *SinglyLinkedList) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	current := sll.Head
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
func (sll *SinglyLinkedList) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sll.AddToTail(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// findNodeByValue вспомогательная функция для поиска узла по значению (Private)
func (sll *SinglyLinkedList) findNodeByValue(value string) *Node {
	current := sll.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// SerializeText сериализует односвязный список в текстовый формат (JSON)
func (sll *SinglyLinkedList) SerializeText() (string, error) {
	data := []string{} // Инициализация пустого слайса
	current := sll.Head
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

// DeserializeText десериализует односвязный список из текстового формата (JSON)
func (sll *SinglyLinkedList) DeserializeText(data string) error {
	var temp []string
	err := json.Unmarshal([]byte(data), &temp)
	if err != nil {
		return fmt.Errorf("ошибка десериализации из текстового формата: %v", err)
	}
	sll.Head = nil
	sll.Size = 0
	for _, value := range temp {
		sll.AddToTail(value)
	}
	return nil
}

// SerializeBinary сериализует односвязный список в бинарный формат
func (sll *SinglyLinkedList) SerializeBinary() ([]byte, error) {
	var result []byte
	current := sll.Head
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

// DeserializeBinary десериализует односвязный список из бинарного формата
func (sll *SinglyLinkedList) DeserializeBinary(data []byte) error {
	sll.Head = nil
	sll.Size = 0
	offset := 0

	for offset < len(data) {
		// Проверяем, достаточно ли байт для чтения длины строки
		if offset+4 > len(data) {
			return fmt.Errorf("недостаточно данных для чтения длины строки")
		}

		// Читаем длину строки
		length := binary.LittleEndian.Uint32(data[offset : offset+4])
		offset += 4

		// Проверяем, достаточно ли байт для чтения строки
		if offset+int(length) > len(data) {
			return fmt.Errorf("недостаточно данных для чтения строки")
		}

		// Читаем строку
		strData := string(data[offset : offset+int(length)])
		offset += int(length)
		sll.AddToTail(strData)
	}

	return nil
}
