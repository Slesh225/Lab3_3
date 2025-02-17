package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

// Queue представляет структуру очереди
type Queue struct {
	Front *Node // Public
	End   *Node // Public
	Size  int   // Public
}

// NewQueue создает новую очередь (Public)
func NewQueue() *Queue {
	return &Queue{}
}

// Push добавляет новый элемент в конец очереди (Public)
func (q *Queue) Push(value string) {
	newNode := &Node{Data: value}
	if q.End == nil {
		q.Front = newNode
		q.End = newNode
	} else {
		q.End.Next = newNode
		q.End = newNode
	}
	q.Size++
}

// Pop удаляет передний элемент из очереди (Public)
func (q *Queue) Pop() {
	if q.Front == nil {
		fmt.Println("Очередь пуста.")
		return
	}
	temp := q.Front
	q.Front = q.Front.Next
	if q.Front == nil {
		q.End = nil
	}
	temp.Next = nil
	q.Size--
}

// Print печатает элементы очереди (Public)
func (q *Queue) Print() {
	temp := q.Front
	for temp != nil {
		fmt.Print(temp.Data, " ")
		temp = temp.Next
	}
	fmt.Println()
}

// SaveToFile сохраняет очередь в файл (Public)
func (q *Queue) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	temp := q.Front
	for temp != nil {
		_, err := file.WriteString(temp.Data + "\n")
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}
		temp = temp.Next
	}
	return nil
}

// LoadFromFile загружает очередь из файла (Public)
func (q *Queue) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		q.Push(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// isEmpty проверяет, пуста ли очередь (Private)
func (q *Queue) isEmpty() bool {
	return q.Front == nil
}

// SerializeText сериализует очередь в текстовый формат (JSON)
func (q *Queue) SerializeText() (string, error) {
	data := []string{} // Инициализация пустого слайса
	current := q.Front
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

// DeserializeText десериализует очередь из текстового формата (JSON)
func (q *Queue) DeserializeText(data string) error {
	var temp []string
	err := json.Unmarshal([]byte(data), &temp)
	if err != nil {
		return fmt.Errorf("ошибка десериализации из текстового формата: %v", err)
	}
	q.Front = nil
	q.End = nil
	q.Size = 0
	for _, value := range temp {
		q.Push(value)
	}
	return nil
}

// SerializeBinary сериализует очередь в бинарный формат
func (q *Queue) SerializeBinary() ([]byte, error) {
	var result []byte
	current := q.Front
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

// DeserializeBinary десериализует очередь из бинарного формата
func (q *Queue) DeserializeBinary(data []byte) error {
	q.Front = nil
	q.End = nil
	q.Size = 0
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
		q.Push(strData)
	}

	return nil
}
