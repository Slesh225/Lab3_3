package main

import (
	"bufio"
	"os"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q.Size != 0 || q.Front != nil || q.End != nil {
		t.Errorf("NewQueue() = %v; want empty queue", q)
	}
}

func TestQueuePush(t *testing.T) {
	q := NewQueue()
	q.Push("test")
	if q.Size != 1 || q.Front.Data != "test" || q.End.Data != "test" {
		t.Errorf("Push() = %v; want queue with one element 'test'", q)
	}

	q.Push("another")
	if q.Size != 2 || q.End.Data != "another" {
		t.Errorf("Push() = %v; want queue with two elements", q)
	}
}

func TestQueuePop(t *testing.T) {
	q := NewQueue()
	q.Push("first")
	q.Push("second")

	q.Pop()
	if q.Size != 1 || q.Front.Data != "second" {
		t.Errorf("Pop() = %v; want queue with one element 'second'", q)
	}

	q.Pop()
	if !q.isEmpty() {
		t.Errorf("Pop() = %v; want empty queue", q)
	}

	q.Pop() // Попытка удалить из пустой очереди
	if q.Size != 0 {
		t.Errorf("Pop() from empty queue = %v; want empty queue", q)
	}
}

func TestQueuePrint(t *testing.T) {
	q := NewQueue()
	q.Push("first")
	q.Push("second")
	// Здесь можно использовать захват вывода для проверки, но для простоты просто вызовем
	q.Print()
}

func TestQueueSaveToFile(t *testing.T) {
	q := NewQueue()
	q.Push("first")
	q.Push("second")

	filename := "test_queue.txt"
	err := q.SaveToFile(filename)
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

func TestQueueLoadFromFile(t *testing.T) {
	q := NewQueue()
	filename := "test_load_queue.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.WriteString("first\nsecond\n")
	file.Close()
	defer os.Remove(filename)

	err = q.LoadFromFile(filename)
	if err != nil {
		t.Errorf("LoadFromFile() error = %v; want nil", err)
	}

	if q.Size != 2 || q.Front.Data != "first" || q.End.Data != "second" {
		t.Errorf("LoadFromFile() = %v; want queue with two elements ['first', 'second']", q)
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue()
	if !q.isEmpty() {
		t.Errorf("isEmpty() = %v; want true", q.isEmpty())
	}

	q.Push("test")
	if q.isEmpty() {
		t.Errorf("isEmpty() = %v; want false", q.isEmpty())
	}
}
func TestQueueSerializeDeserialize(t *testing.T) {
	q := NewQueue()
	q.Push("first")
	q.Push("second")
	q.Push("third")
}

// TestSerializeTextq проверяет сериализацию очереди в текстовый формат (JSON)
func TestSerializeTextq(t *testing.T) {
	q := NewQueue()
	q.Push("first")
	q.Push("second")
	q.Push("third")

	serialized, err := q.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected := `["first","second","third"]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}

	// Сериализация пустой очереди
	emptyQ := NewQueue()
	serialized, err = emptyQ.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected = `[]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}
}

// TestSerializeBinary проверяет сериализацию очереди в бинарный формат
func TestSerializeBinaryq(t *testing.T) {
	q := NewQueue()
	q.Push("first")
	q.Push("second")
	q.Push("third")

	serialized, err := q.SerializeBinary()
	if err != nil {
		t.Errorf("SerializeBinary() error = %v; want nil", err)
	}

	// Ожидаемый результат: длина строки "first" (5) + строка "first" + длина строки "second" (6) + строка "second" + длина строки "third" (5) + строка "third"
	expected := []byte{
		5, 0, 0, 0, // длина "first"
		'f', 'i', 'r', 's', 't', // строка "first"
		6, 0, 0, 0, // длина "second"
		's', 'e', 'c', 'o', 'n', 'd', // строка "second"
		5, 0, 0, 0, // длина "third"
		't', 'h', 'i', 'r', 'd', // строка "third"
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

	// Сериализация пустой очереди
	emptyQ := NewQueue()
	serialized, err = emptyQ.SerializeBinary()
	if err != nil {
		t.Errorf("SerializeBinary() error = %v; want nil", err)
	}

	if len(serialized) != 0 {
		t.Errorf("SerializeBinary() = %v; want empty byte slice", serialized)
	}
}

// TestDeserializeTextq проверяет десериализацию очереди из текстового формата (JSON)
func TestDeserializeTextq(t *testing.T) {
	q := NewQueue()
	data := `["first","second","third"]`

	err := q.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if q.Size != 3 || q.Front.Data != "first" || q.End.Data != "third" {
		t.Errorf("DeserializeText() = %v; want queue with elements [first, second, third]", q)
	}

	// Десериализация пустой очереди
	emptyQ := NewQueue()
	data = `[]`
	err = emptyQ.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if emptyQ.Size != 0 || emptyQ.Front != nil || emptyQ.End != nil {
		t.Errorf("DeserializeText() = %v; want empty queue", emptyQ)
	}

	// Десериализация некорректных данных
	invalidData := `["first","second","third"`
	err = q.DeserializeText(invalidData)
	if err == nil {
		t.Errorf("DeserializeText() = %v; want error", err)
	}
}
func TestDeserializeBinaryq(t *testing.T) {
	q := NewQueue()
	data := []byte{
		5, 0, 0, 0, // длина "first"
		'f', 'i', 'r', 's', 't', // строка "first"
		6, 0, 0, 0, // длина "second"
		's', 'e', 'c', 'o', 'n', 'd', // строка "second"
		5, 0, 0, 0, // длина "third"
		't', 'h', 'i', 'r', 'd', // строка "third"
	}

	err := q.DeserializeBinary(data)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if q.Size != 3 || q.Front.Data != "first" || q.End.Data != "third" {
		t.Errorf("DeserializeBinary() = %v; want queue with elements [first, second, third]", q)
	}

	// Десериализация пустой очереди
	emptyQ := NewQueue()
	emptyData := []byte{} // Пустые данные
	err = emptyQ.DeserializeBinary(emptyData)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if emptyQ.Size != 0 || emptyQ.Front != nil || emptyQ.End != nil {
		t.Errorf("DeserializeBinary() = %v; want empty queue", emptyQ)
	}

	// Десериализация некорректных данных (недостаточно байт)
	invalidData := []byte{5, 0, 0, 0, 'f', 'i', 'r'} // Недостаточно байт для строки "first"
	err = q.DeserializeBinary(invalidData)
	if err == nil {
		t.Errorf("DeserializeBinary() = %v; want error", err)
	}
}
