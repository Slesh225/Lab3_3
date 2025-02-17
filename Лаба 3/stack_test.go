package main

import (
	"bufio"
	"os"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s.Size != 0 || s.Top != nil {
		t.Errorf("NewStack() = %v; want empty stack", s)
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack()
	s.Push("first")
	if s.Size != 1 || s.Top.Data != "first" {
		t.Errorf("Push() = %v; want stack with one element 'first'", s)
	}

	s.Push("second")
	if s.Size != 2 || s.Top.Data != "second" {
		t.Errorf("Push() = %v; want stack with two elements ['second', 'first']", s)
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push("first")
	s.Push("second")

	s.Pop()
	if s.Size != 1 || s.Top.Data != "first" {
		t.Errorf("Pop() = %v; want stack with one element 'first'", s)
	}

	s.Pop()
	if !s.isEmpty() {
		t.Errorf("Pop() = %v; want empty stack", s)
	}

	s.Pop() // Попытка удалить из пустого стека
	if s.Size != 0 {
		t.Errorf("Pop() from empty stack = %v; want empty stack", s)
	}
}

func TestStackPrint(t *testing.T) {
	s := NewStack()
	s.Push("first")
	s.Push("second")
	// Здесь можно использовать захват вывода для проверки, но для простоты просто вызовем
	s.Print()
}

func TestStackSaveToFile(t *testing.T) {
	s := NewStack()
	s.Push("first")
	s.Push("second")

	filename := "test_stack.txt"
	err := s.SaveToFile(filename)
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

	if len(lines) != 2 || lines[0] != "second" || lines[1] != "first" {
		t.Errorf("File content = %v; want ['second', 'first']", lines)
	}
}

func TestStackLoadFromFile(t *testing.T) {
	s := NewStack()
	filename := "test_load_stack.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.WriteString("first\nsecond\n")
	file.Close()
	defer os.Remove(filename)

	err = s.LoadFromFile(filename)
	if err != nil {
		t.Errorf("LoadFromFile() error = %v; want nil", err)
	}

	if s.Size != 2 || s.Top.Data != "second" {
		t.Errorf("LoadFromFile() = %v; want stack with two elements ['second', 'first']", s)
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := NewStack()
	if !s.isEmpty() {
		t.Errorf("isEmpty() = %v; want true", s.isEmpty())
	}

	s.Push("test")
	if s.isEmpty() {
		t.Errorf("isEmpty() = %v; want false", s.isEmpty())
	}
}

func TestStackSaveToFileError(t *testing.T) {
	s := NewStack()
	s.Push("first")

	// Попытка сохранить в несуществующую директорию
	err := s.SaveToFile("/invalid/path/test_stack.txt")
	if err == nil {
		t.Errorf("SaveToFile() with invalid path = %v; want error", err)
	}
}

func TestStackLoadFromFileError(t *testing.T) {
	s := NewStack()

	// Попытка загрузить из несуществующего файла
	err := s.LoadFromFile("nonexistent_file.txt")
	if err == nil {
		t.Errorf("LoadFromFile() with nonexistent file = %v; want error", err)
	}
}

func TestStackPrintEmpty(t *testing.T) {
	s := NewStack()
	// Печать пустого стека
	s.Print()
}

// TestSerializeTextStack проверяет сериализацию стека в текстовый формат (JSON)
func TestSerializeTextStack(t *testing.T) {
	s := NewStack()
	s.Push("first")
	s.Push("second")
	s.Push("third")

	serialized, err := s.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected := `["third","second","first"]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}

	// Сериализация пустого стека
	emptyS := NewStack()
	serialized, err = emptyS.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected = `[]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}
}

// TestDeserializeTextStack проверяет десериализацию стека из текстового формата (JSON)
func TestDeserializeTextStack(t *testing.T) {
	s := NewStack()
	data := `["third","second","first"]`

	err := s.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if s.Size != 3 || s.Top.Data != "third" {
		t.Errorf("DeserializeText() = %v; want stack with elements [third, second, first]", s)
	}

	// Десериализация пустого стека
	emptyS := NewStack()
	data = `[]`
	err = emptyS.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if emptyS.Size != 0 || emptyS.Top != nil {
		t.Errorf("DeserializeText() = %v; want empty stack", emptyS)
	}

	// Десериализация некорректных данных
	invalidData := `["first","second","third"`
	err = s.DeserializeText(invalidData)
	if err == nil {
		t.Errorf("DeserializeText() = %v; want error", err)
	}
}

// TestSerializeBinaryStack проверяет сериализацию стека в бинарный формат
func TestSerializeBinaryStack(t *testing.T) {
	s := NewStack()
	s.Push("third")
	s.Push("second")
	s.Push("first")

	serialized, err := s.SerializeBinary()
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

	// Сериализация пустого стека
	emptyS := NewStack()
	serialized, err = emptyS.SerializeBinary()
	if err != nil {
		t.Errorf("SerializeBinary() error = %v; want nil", err)
	}

	if len(serialized) != 0 {
		t.Errorf("SerializeBinary() = %v; want empty byte slice", serialized)
	}
}

// TestDeserializeBinaryStack проверяет десериализацию стека из бинарного формата
func TestDeserializeBinaryStack(t *testing.T) {
	s := NewStack()
	data := []byte{
		5, 0, 0, 0, // длина "first"
		'f', 'i', 'r', 's', 't', // строка "first"
		6, 0, 0, 0, // длина "second"
		's', 'e', 'c', 'o', 'n', 'd', // строка "second"
		5, 0, 0, 0, // длина "third"
		't', 'h', 'i', 'r', 'd', // строка "third"
	}

	err := s.DeserializeBinary(data)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if s.Size != 3 || s.Top.Data != "third" {
		t.Errorf("DeserializeBinary() = %v; want stack with elements [third, second, first]", s)
	}

	// Десериализация пустого стека
	emptyS := NewStack()
	emptyData := []byte{} // Пустые данные
	err = emptyS.DeserializeBinary(emptyData)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if emptyS.Size != 0 || emptyS.Top != nil {
		t.Errorf("DeserializeBinary() = %v; want empty stack", emptyS)
	}

	// Десериализация некорректных данных (недостаточно байт)
	invalidData := []byte{5, 0, 0, 0, 'f', 'i', 'r'} // Недостаточно байт для строки "first"
	err = s.DeserializeBinary(invalidData)
	if err == nil {
		t.Errorf("DeserializeBinary() = %v; want error", err)
	}
}
