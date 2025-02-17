package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// TreeNode представляет узел в бинарном дереве
type TreeNode struct {
	Digit int
	Left  *TreeNode
	Right *TreeNode
}

// QueueNode представляет узел в очереди для узлов дерева
type QueueNode struct {
	Tree *TreeNode
	Next *QueueNode
}

// QueueTree представляет очередь для узлов дерева
type QueueTree struct {
	Front *QueueNode
	Rear  *QueueNode
	Count int
}

// NewQueueTree создает новую очередь для узлов дерева
func NewQueueTree() *QueueTree {
	return &QueueTree{}
}

// Enqueue добавляет узел дерева в очередь
func (q *QueueTree) Enqueue(node *TreeNode) {
	newNode := &QueueNode{Tree: node}
	if q.Rear == nil {
		q.Front = newNode
		q.Rear = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = newNode
	}
	q.Count++
}

// Dequeue удаляет и возвращает узел дерева из очереди
func (q *QueueTree) Dequeue() *TreeNode {
	if q.Front == nil {
		return nil
	}
	temp := q.Front
	q.Front = q.Front.Next
	if q.Front == nil {
		q.Rear = nil
	}
	temp.Next = nil
	q.Count--
	return temp.Tree
}

// IsEmpty проверяет, пуста ли очередь
func (q *QueueTree) IsEmpty() bool {
	return q.Front == nil
}

// BinaryTree представляет структуру бинарного дерева
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree создает новое бинарное дерево
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Insert добавляет новый узел в бинарное дерево
func (bt *BinaryTree) Insert(digit int) {
	newNode := &TreeNode{Digit: digit}
	if bt.Root == nil {
		bt.Root = newNode
		return
	}

	queue := NewQueueTree()
	queue.Enqueue(bt.Root)
	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if current.Left == nil {
			current.Left = newNode
			return
		} else {
			queue.Enqueue(current.Left)
		}

		if current.Right == nil {
			current.Right = newNode
			return
		} else {
			queue.Enqueue(current.Right)
		}
	}
}

// IsComplete проверяет, является ли бинарное дерево полным
func (bt *BinaryTree) IsComplete() bool {
	if bt.Root == nil {
		return false
	}

	queue := NewQueueTree()
	queue.Enqueue(bt.Root)
	nonFullNode := false

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if current.Left != nil {
			if nonFullNode {
				return false
			}
			queue.Enqueue(current.Left)
		} else {
			nonFullNode = true
		}

		if current.Right != nil {
			if nonFullNode {
				return false
			}
			queue.Enqueue(current.Right)
		} else {
			nonFullNode = true
		}
	}
	return true
}

// FindValue ищет значение в бинарном дереве
func (bt *BinaryTree) FindValue(value int) bool {
	return bt.findValue(bt.Root, value)
}

// findValue вспомогательная функция для поиска значения в бинарном дереве
func (bt *BinaryTree) findValue(current *TreeNode, value int) bool {
	if current == nil {
		return false
	}
	if current.Digit == value {
		return true
	}
	return bt.findValue(current.Left, value) || bt.findValue(current.Right, value)
}

// FindIndex находит значение по конкретному индексу
func (bt *BinaryTree) FindIndex(index int) {
	if index < 0 {
		fmt.Println("Неверный индекс.")
		return
	}

	if bt.Root == nil {
		fmt.Println("Дерево пустое.")
		return
	}

	queue := NewQueueTree()
	queue.Enqueue(bt.Root)
	currentIndex := 0

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		if currentIndex == index {
			fmt.Println("Значение:", current.Digit)
			return
		}
		currentIndex++

		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	fmt.Println("Значение не найдено.")
}

// Display печатает бинарное дерево
func (bt *BinaryTree) Display() {
	if bt.Root == nil {
		fmt.Println("Дерево пустое.")
		return
	}
	bt.printCBT(bt.Root, 0)
}

// printCBT вспомогательная функция для печати бинарного дерева
func (bt *BinaryTree) printCBT(current *TreeNode, level int) {
	if current != nil {
		bt.printCBT(current.Right, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("   ")
		}
		fmt.Println(current.Digit)
		bt.printCBT(current.Left, level+1)
	}
}

// Clear удаляет все узлы из бинарного дерева
func (bt *BinaryTree) Clear() {
	bt.clear(bt.Root)
	bt.Root = nil
}

// clear вспомогательная функция для удаления всех узлов из бинарного дерева
func (bt *BinaryTree) clear(node *TreeNode) {
	if node != nil {
		bt.clear(node.Left)
		bt.clear(node.Right)
		node.Left = nil
		node.Right = nil
	}
}

// LoadFromFile загружает бинарное дерево из файла
func (bt *BinaryTree) LoadFromFile(file string) error {
	bt.Clear()
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return fmt.Errorf("недопустимое значение в файле: %v", err)
		}
		bt.Insert(value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}
	return nil
}

// SaveToFile сохраняет бинарное дерево в файл
func (bt *BinaryTree) SaveToFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer f.Close()

	if bt.Root == nil {
		return nil // Пустое дерево, файл будет пустым
	}

	queue := NewQueueTree()
	queue.Enqueue(bt.Root)
	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if _, err := f.WriteString(fmt.Sprintf("%d\n", current.Digit)); err != nil {
			return fmt.Errorf("ошибка записи в файл: %v", err)
		}

		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}
	return nil
}

// SerializeText сериализует бинарное дерево в текстовый формат (JSON)
func (bt *BinaryTree) SerializeText() (string, error) {
	data, err := json.Marshal(bt)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации в текстовый формат: %v", err)
	}
	return string(data), nil
}

// DeserializeText десериализует бинарное дерево из текстового формата (JSON)
func (bt *BinaryTree) DeserializeText(data string) error {
	err := json.Unmarshal([]byte(data), bt)
	if err != nil {
		return fmt.Errorf("ошибка десериализации из текстового формата: %v", err)
	}
	return nil
}

// SerializeBinary сериализует бинарное дерево в бинарный формат
func (bt *BinaryTree) SerializeBinary() ([]byte, error) {
	var result []byte
	queue := NewQueueTree()
	queue.Enqueue(bt.Root)

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		if current == nil {
			result = append(result, 0) // Маркер для nil
			continue
		}

		// Записываем значение узла
		valueBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(valueBytes, uint32(current.Digit))
		result = append(result, valueBytes...)

		queue.Enqueue(current.Left)
		queue.Enqueue(current.Right)
	}

	return result, nil
}

// DeserializeBinary десериализует бинарное дерево из бинарного формата
func (bt *BinaryTree) DeserializeBinary(data []byte) error {
	bt.Root = nil
	queue := NewQueueTree()
	index := 0

	if len(data) == 0 {
		return nil
	}

	// Читаем корневой узел
	if index+4 > len(data) {
		return fmt.Errorf("недостаточно данных для чтения корневого узла")
	}
	value := int(binary.LittleEndian.Uint32(data[index : index+4]))
	index += 4
	bt.Root = &TreeNode{Digit: value}
	queue.Enqueue(bt.Root)

	for !queue.IsEmpty() && index < len(data) {
		current := queue.Dequeue()

		// Левый дочерний узел
		if index >= len(data) {
			break
		}
		if data[index] == 0 {
			index++
		} else {
			if index+4 > len(data) {
				return fmt.Errorf("недостаточно данных для чтения левого узла")
			}
			value := int(binary.LittleEndian.Uint32(data[index : index+4]))
			index += 4
			current.Left = &TreeNode{Digit: value}
			queue.Enqueue(current.Left)
		}

		// Правый дочерний узел
		if index >= len(data) {
			break
		}
		if data[index] == 0 {
			index++
		} else {
			if index+4 > len(data) {
				return fmt.Errorf("недостаточно данных для чтения правого узла")
			}
			value := int(binary.LittleEndian.Uint32(data[index : index+4]))
			index += 4
			current.Right = &TreeNode{Digit: value}
			queue.Enqueue(current.Right)
		}
	}

	return nil
}
