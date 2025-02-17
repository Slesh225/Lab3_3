package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

// Array представляет структуру массива
type Array struct {
	maxCapacity int
	size        int
	data        []string
}

// NewArray создает новый массив
func NewArray(capacity int) *Array {
	return &Array{
		maxCapacity: capacity,
		size:        0,
		data:        make([]string, capacity),
	}
}

// Add вставляет элемент по указанному индексу
func (a *Array) Add(index int, value string) {
	if index < 0 || index > a.size || a.size >= a.maxCapacity {
		fmt.Println("Неверный индекс или массив заполнен")
		return
	}
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.size++
}

// AddToTheEnd добавляет элемент в конец массива
func (a *Array) AddToTheEnd(value string) {
	if a.size >= a.maxCapacity {
		fmt.Println("Массив заполнен")
		return
	}
	a.data[a.size] = value
	a.size++
}

// Remove удаляет элемент по указанному индексу
func (a *Array) Remove(index int) {
	if index < 0 || index >= a.size {
		fmt.Println("Неверный индекс")
		return
	}
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
}

// Replace заменяет элемент по указанному индексу
func (a *Array) Replace(index int, value string) {
	if index < 0 || index >= a.size {
		fmt.Println("Неверный индекс")
		return
	}
	a.data[index] = value
}

// Print выводит элементы массива
func (a *Array) Print() {
	for i := 0; i < a.size; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(a.data[i])
	}
	fmt.Println()
}

// Length возвращает количество элементов в массиве
func (a *Array) Length() int {
	return a.size
}

// SaveToFile сохраняет массив в файл
func (a *Array) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	for i := 0; i < a.size; i++ {
		_, err := file.WriteString(a.data[i] + "\n")
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}
	}
	return nil
}

// LoadFromFile загружает массив из файла
func (a *Array) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	a.size = 0
	for scanner.Scan() && a.size < a.maxCapacity {
		a.data[a.size] = scanner.Text()
		a.size++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// Get возвращает элемент по указанному индексу
func (a *Array) Get(index int) string {
	if index < 0 || index >= a.size {
		fmt.Println("Неверный индекс")
		return ""
	}
	return a.data[index]
}

// Equals сравнивает два массива
func (a *Array) Equals(other *Array) bool {
	if other == nil {
		return false
	}
	if a.size != other.size {
		return false
	}
	for i := 0; i < a.size; i++ {
		if a.data[i] != other.data[i] {
			return false
		}
	}
	return true
}

// SerializeText сериализует массив в текстовом формате
func (a *Array) SerializeText() (string, error) {
	data, err := json.Marshal(a.data[:a.size])
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации в текстовый формат: %v", err)
	}
	return string(data), nil
}

// DeserializeText десериализует массив из текстового формата
func (a *Array) DeserializeText(data string) error {
	var temp []string
	err := json.Unmarshal([]byte(data), &temp)
	if err != nil {
		return fmt.Errorf("ошибка десериализации из текстового формата: %v", err)
	}
	a.size = 0
	for _, value := range temp {
		if a.size >= a.maxCapacity {
			break // Прекращаем заполнение, если массив заполнен
		}
		a.data[a.size] = value
		a.size++
	}
	return nil
}

// SerializeBinary сериализует массив в бинарном формате
func (a *Array) SerializeBinary() ([]byte, error) {
	var result []byte
	for i := 0; i < a.size; i++ {
		// Преобразуем строку в байты и добавляем её длину перед данными
		strBytes := []byte(a.data[i])
		length := uint32(len(strBytes))
		lengthBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(lengthBytes, length)
		result = append(result, lengthBytes...)
		result = append(result, strBytes...)
	}
	return result, nil
}

// DeserializeBinary десериализует массив из бинарного формата
func (a *Array) DeserializeBinary(data []byte) error {
	a.size = 0
	offset := 0
	for offset < len(data) && a.size < a.maxCapacity {
		// Читаем длину строки
		if offset+4 > len(data) {
			return fmt.Errorf("недостаточно данных для чтения длины строки")
		}
		length := binary.LittleEndian.Uint32(data[offset : offset+4])
		offset += 4

		// Читаем строку
		if offset+int(length) > len(data) {
			return fmt.Errorf("недостаточно данных для чтения строки")
		}
		strBytes := data[offset : offset+int(length)]
		a.data[a.size] = string(strBytes)
		a.size++
		offset += int(length)
	}
	return nil
}
