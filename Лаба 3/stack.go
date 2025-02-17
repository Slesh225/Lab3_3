package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

// Stack представляет структуру стека
type Stack struct {
	Top  *Node // Public
	Size int   // Public
}

// NewStack создает новый стек (Public)
func NewStack() *Stack {
	return &Stack{}
}

// Push добавляет новый элемент на вершину стека (Public)
func (s *Stack) Push(value string) {
	newNode := &Node{Data: value, Next: s.Top}
	s.Top = newNode
	s.Size++
}

// Pop удаляет верхний элемент из стека (Public)
func (s *Stack) Pop() {
	if s.Top == nil {
		fmt.Println("Стек пуст.")
		return
	}
	temp := s.Top
	s.Top = s.Top.Next
	temp.Next = nil
	s.Size--
}

// Print выводит элементы стека (Public)
func (s *Stack) Print() {
	temp := s.Top
	for temp != nil {
		fmt.Print(temp.Data, " ")
		temp = temp.Next
	}
	fmt.Println()
}

// SaveToFile сохраняет стек в файл (Public)
func (s *Stack) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	temp := s.Top
	for temp != nil {
		_, err := file.WriteString(temp.Data + "\n")
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}
		temp = temp.Next
	}
	return nil
}

// LoadFromFile загружает стек из файла (Public)
func (s *Stack) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s.Push(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// isEmpty проверяет, пуст ли стек (Private)
func (s *Stack) isEmpty() bool {
	return s.Top == nil
}

// SerializeText сериализует стек в текстовый формат (JSON)
func (s *Stack) SerializeText() (string, error) {
	data := []string{} // Инициализация пустого слайса
	current := s.Top
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

// DeserializeText десериализует стек из текстового формата (JSON)
func (s *Stack) DeserializeText(data string) error {
	var temp []string
	err := json.Unmarshal([]byte(data), &temp)
	if err != nil {
		return fmt.Errorf("ошибка десериализации из текстового формата: %v", err)
	}
	s.Top = nil
	s.Size = 0
	for i := len(temp) - 1; i >= 0; i-- {
		s.Push(temp[i])
	}
	return nil
}

// SerializeBinary сериализует стек в бинарный формат
func (s *Stack) SerializeBinary() ([]byte, error) {
	var result []byte
	current := s.Top
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

// DeserializeBinary десериализует стек из бинарного формата
func (s *Stack) DeserializeBinary(data []byte) error {
	s.Top = nil
	s.Size = 0
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
		s.Push(strData)
	}

	return nil
}
