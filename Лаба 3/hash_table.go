package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// HashNode представляет узел в хэш-таблице
type HashNode struct {
	Key   string    // Public
	Value string    // Public
	Next  *HashNode // Public
}

// HashTable представляет структуру хэш-таблицы
type HashTable struct {
	Capacity int         // Public
	Table    []*HashNode // Public
}

// NewHashTable создает новую хэш-таблицу (Public)
func NewHashTable(size int) *HashTable {
	return &HashTable{
		Capacity: size,
		Table:    make([]*HashNode, size),
	}
}

// Size возвращает количество элементов в хэш-таблице
func (ht *HashTable) Size() int {
	size := 0
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			size++
			current = current.Next
		}
	}
	return size
}

// HashFunction вычисляет индекс хэша для заданного ключа (Public)
func (ht *HashTable) HashFunction(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (hash*31 + int(ch)) % ht.Capacity
	}
	return hash
}

// HSet вставляет или обновляет пару ключ-значение в хэш-таблице (Public)
func (ht *HashTable) HSet(key, value string) {
	index := ht.HashFunction(key)
	current := ht.Table[index]

	for current != nil {
		if current.Key == key {
			current.Value = value
			return
		}
		current = current.Next
	}

	newNode := &HashNode{Key: key, Value: value, Next: ht.Table[index]}
	ht.Table[index] = newNode
}

// HGet извлекает значение, связанное с ключом (Public)
func (ht *HashTable) HGet(key string) {
	index := ht.HashFunction(key)
	current := ht.Table[index]

	for current != nil {
		if current.Key == key {
			fmt.Printf("Значение для ключа [%s]: %s\n", key, current.Value)
			return
		}
		current = current.Next
	}

	fmt.Printf("Ключ [%s] не найден.\n", key)
}

// HDel удаляет пару ключ-значение из хэш-таблицы (Public)
func (ht *HashTable) HDel(key string) {
	index := ht.HashFunction(key)
	current := ht.Table[index]
	var prev *HashNode

	for current != nil {
		if current.Key == key {
			if prev == nil {
				ht.Table[index] = current.Next
			} else {
				prev.Next = current.Next
			}
			return
		}
		prev = current
		current = current.Next
	}

	fmt.Printf("Ключ [%s] не найден для удаления.\n", key)
}

// Clear удаляет все элементы из хэш-таблицы (Public)
func (ht *HashTable) Clear() {
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			temp := current
			current = current.Next
			temp.Next = nil
		}
		ht.Table[i] = nil
	}
}

// HPrint печатает содержимое хэш-таблицы (Public)
func (ht *HashTable) HPrint() {
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		if current != nil {
			fmt.Printf("[%d]: ", i)
			for current != nil {
				fmt.Printf("%s => %s ", current.Key, current.Value)
				current = current.Next
			}
			fmt.Println()
		}
	}
}

// LoadFromFile загружает хэш-таблицу из файла (Public)
func (ht *HashTable) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			ht.HSet(parts[0], parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// SaveToFile сохраняет хэш-таблицу в файл (Public)
func (ht *HashTable) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()

	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			_, err := file.WriteString(fmt.Sprintf("%s %s\n", current.Key, current.Value))
			if err != nil {
				return fmt.Errorf("ошибка записи в файл: %v", err)
			}
			current = current.Next
		}
	}
	return nil
}

// findNodeByKey вспомогательная функция для поиска узла по ключу (Private)
func (ht *HashTable) findNodeByKey(key string) *HashNode {
	index := ht.HashFunction(key)
	current := ht.Table[index]
	for current != nil {
		if current.Key == key {
			return current
		}
		current = current.Next
	}
	return nil
}

// SerializeText сериализует хэш-таблицу в текстовый формат (JSON)
func (ht *HashTable) SerializeText() (string, error) {
	var data []string
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			data = append(data, fmt.Sprintf("%s:%s", current.Key, current.Value))
			current = current.Next
		}
	}
	result, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации в текстовый формат: %v", err)
	}
	return string(result), nil
}

// DeserializeText десериализует хэш-таблицу из текстового формата (JSON)
func (ht *HashTable) DeserializeText(data string) error {
	var temp []string
	err := json.Unmarshal([]byte(data), &temp)
	if err != nil {
		return fmt.Errorf("ошибка десериализации из текстового формата: %v", err)
	}
	ht.Clear()
	for _, pair := range temp {
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			ht.HSet(parts[0], parts[1])
		}
	}
	return nil
}

// SerializeBinary сериализует хэш-таблицу в бинарный формат
func (ht *HashTable) SerializeBinary() ([]byte, error) {
	var result []byte
	for i := 0; i < ht.Capacity; i++ {
		current := ht.Table[i]
		for current != nil {
			// Записываем длину ключа
			keyLength := uint32(len(current.Key))
			keyLengthBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(keyLengthBytes, keyLength)
			result = append(result, keyLengthBytes...)

			// Записываем ключ
			result = append(result, []byte(current.Key)...)

			// Записываем длину значения
			valueLength := uint32(len(current.Value))
			valueLengthBytes := make([]byte, 4)
			binary.LittleEndian.PutUint32(valueLengthBytes, valueLength)
			result = append(result, valueLengthBytes...)

			// Записываем значение
			result = append(result, []byte(current.Value)...)

			current = current.Next
		}
	}
	return result, nil
}

// DeserializeBinary десериализует хэш-таблицу из бинарного формата
func (ht *HashTable) DeserializeBinary(data []byte) error {
	ht.Clear()
	offset := 0
	for offset < len(data) {
		// Читаем длину ключа
		keyLength := binary.LittleEndian.Uint32(data[offset : offset+4])
		offset += 4

		// Читаем ключ
		key := string(data[offset : offset+int(keyLength)])
		offset += int(keyLength)

		// Читаем длину значения
		valueLength := binary.LittleEndian.Uint32(data[offset : offset+4])
		offset += 4

		// Читаем значение
		value := string(data[offset : offset+int(valueLength)])
		offset += int(valueLength)

		ht.HSet(key, value)
	}
	return nil
}
