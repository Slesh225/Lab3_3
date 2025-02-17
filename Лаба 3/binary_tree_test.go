package main

import (
	"bufio"
	"os"
	"testing"
)

// TestNewBinaryTree проверяет создание нового бинарного дерева
func TestNewBinaryTree(t *testing.T) {
	bt := NewBinaryTree()
	if bt.Root != nil {
		t.Errorf("NewBinaryTree() = %v; want empty tree", bt)
	}
}

// TestInsert проверяет вставку элементов в бинарное дерево
func TestInsert(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	if bt.Root == nil || bt.Root.Digit != 10 {
		t.Errorf("Insert() = %v; want tree with root 10", bt)
	}

	bt.Insert(5)
	if bt.Root.Left == nil || bt.Root.Left.Digit != 5 {
		t.Errorf("Insert() = %v; want tree with left child 5", bt)
	}

	bt.Insert(15)
	if bt.Root.Right == nil || bt.Root.Right.Digit != 15 {
		t.Errorf("Insert() = %v; want tree with right child 15", bt)
	}

	// Вставка дубликатов
	bt.Insert(10)
	if bt.Root.Digit != 10 || bt.Root.Left.Digit != 5 || bt.Root.Right.Digit != 15 {
		t.Errorf("Insert() = %v; want tree with elements [10, 5, 15]", bt)
	}
}

// TestIsComplete проверяет, является ли бинарное дерево полным
func TestIsComplete(t *testing.T) {
	bt := NewBinaryTree()
	if bt.IsComplete() {
		t.Errorf("IsComplete() = %v; want false for empty tree", bt.IsComplete())
	}

	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)
	if !bt.IsComplete() {
		t.Errorf("IsComplete() = %v; want true for complete tree", bt.IsComplete())
	}

	bt.Insert(3)
	if !bt.IsComplete() {
		t.Errorf("IsComplete() = %v; want true for complete tree", bt.IsComplete())
	}

	bt.Insert(7)
	if !bt.IsComplete() {
		t.Errorf("IsComplete() = %v; want true for complete tree", bt.IsComplete())
	}

	bt.Insert(12)
	if bt.IsComplete() {
		t.Errorf("IsComplete() = %v; want false for incomplete tree", bt.IsComplete())
	}
}

// TestFindValue проверяет поиск значения в бинарном дереве
func TestFindValue(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)

	if !bt.FindValue(10) {
		t.Errorf("FindValue() = %v; want true for value 10", bt.FindValue(10))
	}

	if !bt.FindValue(5) {
		t.Errorf("FindValue() = %v; want true for value 5", bt.FindValue(5))
	}

	if !bt.FindValue(15) {
		t.Errorf("FindValue() = %v; want true for value 15", bt.FindValue(15))
	}

	if bt.FindValue(20) {
		t.Errorf("FindValue() = %v; want false for value 20", bt.FindValue(20))
	}

	// Поиск в пустом дереве
	emptyBt := NewBinaryTree()
	if emptyBt.FindValue(10) {
		t.Errorf("FindValue() = %v; want false for empty tree", emptyBt.FindValue(10))
	}
}

// TestFindIndex проверяет поиск значения по индексу
func TestFindIndex(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)

	// Проверка корректных индексов
	bt.FindIndex(0) // 10
	bt.FindIndex(1) // 5
	bt.FindIndex(2) // 15

	// Проверка несуществующего индекса
	bt.FindIndex(3) // Несуществующий индекс

	// Проверка отрицательного индекса
	bt.FindIndex(-1) // Неверный индекс

	// Проверка пустого дерева
	emptyBt := NewBinaryTree()
	emptyBt.FindIndex(0) // Дерево пустое
}

// TestDisplay проверяет печать бинарного дерева
func TestDisplay(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Display()

	// Печать пустого дерева
	emptyBt := NewBinaryTree()
	emptyBt.Display()
}

// TestClear проверяет очистку бинарного дерева
func TestClear(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)

	bt.Clear()
	if bt.Root != nil {
		t.Errorf("Clear() = %v; want empty tree", bt)
	}

	// Очистка пустого дерева
	emptyBt := NewBinaryTree()
	emptyBt.Clear()
	if emptyBt.Root != nil {
		t.Errorf("Clear() = %v; want empty tree", emptyBt)
	}
}

// TestSaveToFile проверяет сохранение бинарного дерева в файл
func TestSaveToFile(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)

	filename := "test_tree.txt"
	err := bt.SaveToFile(filename)
	if err != nil {
		t.Errorf("SaveToFile() error = %v; want nil", err)
	}
	defer os.Remove(filename)

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

	if len(lines) != 3 || lines[0] != "10" || lines[1] != "5" || lines[2] != "15" {
		t.Errorf("File content = %v; want ['10', '5', '15']", lines)
	}

	// Сохранение пустого дерева
	emptyBt := NewBinaryTree()
	emptyFilename := "empty_tree.txt"
	err = emptyBt.SaveToFile(emptyFilename)
	if err != nil {
		t.Errorf("SaveToFile() error = %v; want nil", err)
	}
	defer os.Remove(emptyFilename)

	// Проверка содержимого пустого файла
	file, err = os.Open(emptyFilename)
	if err != nil {
		t.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) != 0 {
		t.Errorf("File content = %v; want empty file", lines)
	}
}

// TestLoadFromFile проверяет загрузку бинарного дерева из файла
func TestLoadFromFile(t *testing.T) {
	bt := NewBinaryTree()
	filename := "test_load_tree.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.WriteString("10\n5\n15\n")
	file.Close()
	defer os.Remove(filename)

	err = bt.LoadFromFile(filename)
	if err != nil {
		t.Errorf("LoadFromFile() error = %v; want nil", err)
	}

	if bt.Root == nil || bt.Root.Digit != 10 || bt.Root.Left.Digit != 5 || bt.Root.Right.Digit != 15 {
		t.Errorf("LoadFromFile() = %v; want tree with elements [10, 5, 15]", bt)
	}

	// Загрузка из пустого файла
	emptyFilename := "empty_load_tree.txt"
	file, err = os.Create(emptyFilename)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	file.Close()
	defer os.Remove(emptyFilename)

	emptyBt := NewBinaryTree()
	err = emptyBt.LoadFromFile(emptyFilename)
	if err != nil {
		t.Errorf("LoadFromFile() error = %v; want nil", err)
	}

	if emptyBt.Root != nil {
		t.Errorf("LoadFromFile() = %v; want empty tree", emptyBt)
	}

	// Загрузка из несуществующего файла
	nonexistentFilename := "nonexistent_file.txt"
	err = bt.LoadFromFile(nonexistentFilename)
	if err == nil {
		t.Errorf("LoadFromFile() = %v; want error", err)
	}
}

// TestQueueTree проверяет работу очереди для узлов дерева
func TestQueueTree(t *testing.T) {
	queue := NewQueueTree()
	if !queue.IsEmpty() {
		t.Errorf("NewQueueTree() = %v; want empty queue", queue)
	}

	node1 := &TreeNode{Digit: 10}
	node2 := &TreeNode{Digit: 5}
	queue.Enqueue(node1)
	queue.Enqueue(node2)

	if queue.IsEmpty() {
		t.Errorf("Enqueue() = %v; want non-empty queue", queue)
	}

	dequeued := queue.Dequeue()
	if dequeued.Digit != 10 {
		t.Errorf("Dequeue() = %v; want node with digit 10", dequeued)
	}

	dequeued = queue.Dequeue()
	if dequeued.Digit != 5 {
		t.Errorf("Dequeue() = %v; want node with digit 5", dequeued)
	}

	if !queue.IsEmpty() {
		t.Errorf("Dequeue() = %v; want empty queue", queue)
	}

	// Попытка удалить из пустой очереди
	dequeued = queue.Dequeue()
	if dequeued != nil {
		t.Errorf("Dequeue() = %v; want nil", dequeued)
	}
}

// TestSerializeText проверяет сериализацию бинарного дерева в текстовый формат (JSON)
func TestSerializeText(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)

	serialized, err := bt.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected := `{"Root":{"Digit":10,"Left":{"Digit":5,"Left":null,"Right":null},"Right":{"Digit":15,"Left":null,"Right":null}}}`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}

	// Сериализация пустого дерева
	emptyBt := NewBinaryTree()
	serialized, err = emptyBt.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected = `{"Root":null}`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}
}

// TestDeserializeText проверяет десериализацию бинарного дерева из текстового формата (JSON)
func TestDeserializeText(t *testing.T) {
	bt := NewBinaryTree()
	data := `{"Root":{"Digit":10,"Left":{"Digit":5,"Left":null,"Right":null},"Right":{"Digit":15,"Left":null,"Right":null}}}`

	err := bt.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if bt.Root == nil || bt.Root.Digit != 10 || bt.Root.Left.Digit != 5 || bt.Root.Right.Digit != 15 {
		t.Errorf("DeserializeText() = %v; want tree with elements [10, 5, 15]", bt)
	}

	// Десериализация пустого дерева
	emptyBt := NewBinaryTree()
	data = `{"Root":null}`
	err = emptyBt.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if emptyBt.Root != nil {
		t.Errorf("DeserializeText() = %v; want empty tree", emptyBt)
	}

	// Десериализация некорректных данных
	invalidData := `{"Root":{"Digit":10,"Left":{"Digit":5,"Left":null,"Right":null},"Right":{"Digit":15,"Left":null,"Right":null}`
	err = bt.DeserializeText(invalidData)
	if err == nil {
		t.Errorf("DeserializeText() = %v; want error", err)
	}
}

// TestSerializeBinary проверяет сериализацию бинарного дерева в бинарный формат
func TestSerializeBinary(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)

	serialized, err := bt.SerializeBinary()
	if err != nil {
		t.Errorf("SerializeBinary() error = %v; want nil", err)
	}

	// Ожидаемый результат: 10 (4 байта) + 5 (4 байта) + 15 (4 байта)
	expected := []byte{
		10, 0, 0, 0, // 10
		5, 0, 0, 0, // 5
		15, 0, 0, 0, // 15
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

	// Сериализация пустого дерева
	emptyBt := NewBinaryTree()
	serialized, err = emptyBt.SerializeBinary()
	if err != nil {
		t.Errorf("SerializeBinary() error = %v; want nil", err)
	}

	if len(serialized) != 0 {
		t.Errorf("SerializeBinary() = %v; want empty byte slice", serialized)
	}
}

// TestDeserializeBinary проверяет десериализацию бинарного дерева из бинарного формата
// TestDeserializeBinaryBT проверяет десериализацию бинарного дерева из бинарного формата
func TestDeserializeBinaryBT(t *testing.T) {
	// Создаем бинарное дерево и заполняем его данными
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(7)

	// Сериализуем дерево в бинарный формат
	serialized, err := bt.SerializeBinary()
	if err != nil {
		t.Fatalf("SerializeBinary() error = %v; want nil", err)
	}

	// Создаем новое дерево и десериализуем данные
	newBt := NewBinaryTree()
	err = newBt.DeserializeBinary(serialized)
	if err != nil {
		t.Fatalf("DeserializeBinary() error = %v; want nil", err)
	}

	// Проверяем, что десериализованное дерево совпадает с исходным
	if !treesEqual(bt.Root, newBt.Root) {
		t.Errorf("DeserializeBinary() = %v; want %v", newBt, bt)
	}

	// Десериализация пустого дерева
	emptyBt := NewBinaryTree()
	emptyData := []byte{} // Пустые данные
	err = emptyBt.DeserializeBinary(emptyData)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}
	if emptyBt.Root != nil {
		t.Errorf("DeserializeBinary() = %v; want empty tree", emptyBt)
	}

	// Десериализация некорректных данных (недостаточно байт)
	invalidData := []byte{10, 0, 0, 0, 5, 0, 0} // Недостаточно байт для второго узла
	err = bt.DeserializeBinary(invalidData)
	if err == nil {
		t.Errorf("DeserializeBinary() = %v; want error", err)
	}
}

// treesEqual рекурсивно проверяет, равны ли два дерева
func treesEqual(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Digit != node2.Digit {
		return false
	}
	return treesEqual(node1.Left, node2.Left) && treesEqual(node1.Right, node2.Right)
}
