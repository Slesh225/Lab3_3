package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// captureStdout захватывает вывод в stdout для тестирования функций, которые печатают в консоль.
func captureStdout(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout
	var output string
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	output = buf.String()
	return output
}

// Тесты для SinglyLinkedList
func TestNewSinglyLinkedList(t *testing.T) {
	sll := NewSinglyLinkedList()
	assert.NotNil(t, sll, "Новый список не должен быть nil")
	assert.Nil(t, sll.Head, "Head нового списка должен быть nil")
	assert.Equal(t, 0, sll.Size, "Размер нового списка должен быть 0")
}

func TestSinglyLinkedList_AddToHead(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToHead("test")
	assert.Equal(t, "test", sll.Head.Data, "Элемент в начале списка должен быть 'test'")
	assert.Equal(t, 1, sll.Size, "Размер списка должен быть 1 после добавления элемента")
}

func TestSinglyLinkedList_AddToTail(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test")
	assert.Equal(t, "test", sll.Head.Data, "Элемент в конце списка должен быть 'test'")
	assert.Equal(t, 1, sll.Size, "Размер списка должен быть 1 после добавления элемента")

	sll.AddToTail("test2")
	assert.Equal(t, "test2", sll.Head.Next.Data, "Второй элемент должен быть 'test2'")
	assert.Equal(t, 2, sll.Size, "Размер списка должен быть 2 после добавления второго элемента")
}

func TestSinglyLinkedList_RemoveHead(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToHead("test1")
	sll.AddToHead("test2")

	sll.RemoveHead()
	assert.Equal(t, "test1", sll.Head.Data, "Head должен быть 'test1' после удаления 'test2'")
	assert.Equal(t, 1, sll.Size, "Размер списка должен быть 1 после удаления элемента")

	sll.RemoveHead()
	assert.Nil(t, sll.Head, "Head должен быть nil после удаления всех элементов")
	assert.Equal(t, 0, sll.Size, "Размер списка должен быть 0 после удаления всех элементов")
}

func TestSinglyLinkedList_RemoveTail(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")

	sll.RemoveTail()
	assert.Equal(t, "test1", sll.Head.Data, "Head должен быть 'test1' после удаления 'test2'")
	assert.Equal(t, 1, sll.Size, "Размер списка должен быть 1 после удаления элемента")

	sll.RemoveTail()
	assert.Nil(t, sll.Head, "Head должен быть nil после удаления всех элементов")
	assert.Equal(t, 0, sll.Size, "Размер списка должен быть 0 после удаления всех элементов")
}

func TestSinglyLinkedList_RemoveByValue(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	sll.RemoveByValue("test2")
	assert.Equal(t, "test1", sll.Head.Data, "Head должен быть 'test1'")
	assert.Equal(t, "test3", sll.Head.Next.Data, "Следующий элемент должен быть 'test3'")
	assert.Equal(t, 2, sll.Size, "Размер списка должен быть 2 после удаления элемента")

	sll.RemoveByValue("test1")
	assert.Equal(t, "test3", sll.Head.Data, "Head должен быть 'test3' после удаления 'test1'")
	assert.Equal(t, 1, sll.Size, "Размер списка должен быть 1 после удаления элемента")

	sll.RemoveByValue("test3")
	assert.Nil(t, sll.Head, "Head должен быть nil после удаления всех элементов")
	assert.Equal(t, 0, sll.Size, "Размер списка должен быть 0 после удаления всех элементов")
}

func TestSinglyLinkedList_Search(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	node := sll.Search("test2")
	assert.NotNil(t, node, "Узел с значением 'test2' должен быть найден")
	assert.Equal(t, "test2", node.Data, "Найденный узел должен содержать значение 'test2'")

	node = sll.Search("test4")
	assert.Nil(t, node, "Узел с значением 'test4' не должен быть найден")
}

func TestSinglyLinkedList_Print(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	output := captureStdout(func() {
		sll.Print()
	})
	assert.Equal(t, "test1 test2 test3 \n", output, "Вывод должен быть 'test1 test2 test3'")
}

func TestSinglyLinkedList_SaveAndLoadFromFile(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	err := sll.SaveToFile("test.txt")
	assert.NoError(t, err, "Ошибка при сохранении в файл")

	newSll := NewSinglyLinkedList()
	err = newSll.LoadFromFile("test.txt")
	assert.NoError(t, err, "Ошибка при загрузке из файла")

	output := captureStdout(func() {
		newSll.Print()
	})
	assert.Equal(t, "test1 test2 test3 \n", output, "Вывод должен быть 'test1 test2 test3'")

	os.Remove("test.txt")
}

func TestSinglyLinkedList_LoadFromFile_NonExistentFile(t *testing.T) {
	sll := NewSinglyLinkedList()
	err := sll.LoadFromFile("nonexistent.txt")
	assert.Error(t, err, "Ожидается ошибка при загрузке из несуществующего файла")
}

func TestSinglyLinkedList_SaveToFile_WriteError(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")

	err := sll.SaveToFile("/invalid/path/test.txt")
	assert.Error(t, err, "Ожидается ошибка при сохранении в невалидный путь")
}

func TestSinglyLinkedList_RemoveHead_EmptyList(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.RemoveHead()
	assert.Nil(t, sll.Head, "Head должен остаться nil при удалении из пустого списка")
	assert.Equal(t, 0, sll.Size, "Размер списка должен остаться 0 при удалении из пустого списка")
}

func TestSinglyLinkedList_RemoveTail_EmptyList(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.RemoveTail()
	assert.Nil(t, sll.Head, "Head должен остаться nil при удалении из пустого списка")
	assert.Equal(t, 0, sll.Size, "Размер списка должен остаться 0 при удалении из пустого списка")
}

func TestSinglyLinkedList_RemoveByValue_EmptyList(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.RemoveByValue("test")
	assert.Nil(t, sll.Head, "Head должен остаться nil при удалении из пустого списка")
	assert.Equal(t, 0, sll.Size, "Размер списка должен остаться 0 при удалении из пустого списка")
}

func TestSinglyLinkedList_Search_EmptyList(t *testing.T) {
	sll := NewSinglyLinkedList()
	node := sll.Search("test")
	assert.Nil(t, node, "Узел не должен быть найден в пустом списке")
}

func TestSinglyLinkedList_Print_EmptyList(t *testing.T) {
	sll := NewSinglyLinkedList()
	output := captureStdout(func() {
		sll.Print()
	})
	assert.Equal(t, "\n", output, "Вывод пустого списка должен быть пустой строкой")
}

func TestSinglyLinkedList_AddToHead_MultipleElements(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToHead("test1")
	sll.AddToHead("test2")
	sll.AddToHead("test3")

	assert.Equal(t, "test3", sll.Head.Data, "Head должен быть 'test3'")
	assert.Equal(t, "test2", sll.Head.Next.Data, "Следующий элемент должен быть 'test2'")
	assert.Equal(t, "test1", sll.Head.Next.Next.Data, "Третий элемент должен быть 'test1'")
	assert.Equal(t, 3, sll.Size, "Размер списка должен быть 3")
}

func TestSinglyLinkedList_AddToTail_MultipleElements(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	assert.Equal(t, "test1", sll.Head.Data, "Head должен быть 'test1'")
	assert.Equal(t, "test2", sll.Head.Next.Data, "Следующий элемент должен быть 'test2'")
	assert.Equal(t, "test3", sll.Head.Next.Next.Data, "Третий элемент должен быть 'test3'")
	assert.Equal(t, 3, sll.Size, "Размер списка должен быть 3")
}
func TestSinglyLinkedList_AddAndRemoveCombinations(t *testing.T) {
	sll := NewSinglyLinkedList()

	// Добавляем элементы в начало и конец
	sll.AddToHead("head1")
	sll.AddToTail("tail1")
	sll.AddToHead("head2")
	sll.AddToTail("tail2")

	assert.Equal(t, 4, sll.Size, "Размер списка должен быть 4")
	assert.Equal(t, "head2", sll.Head.Data, "Head должен быть 'head2'")
	assert.Equal(t, "tail2", sll.Head.Next.Next.Next.Data, "Последний элемент должен быть 'tail2'")

	// Удаляем элементы из начала и конца
	sll.RemoveHead()
	sll.RemoveTail()

	assert.Equal(t, 2, sll.Size, "Размер списка должен быть 2")
	assert.Equal(t, "head1", sll.Head.Data, "Head должен быть 'head1'")
	assert.Equal(t, "tail1", sll.Head.Next.Data, "Последний элемент должен быть 'tail1'")
}
func TestSinglyLinkedList_RemoveByValue_MiddleElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	sll.RemoveByValue("test2")
	assert.Equal(t, 2, sll.Size, "Размер списка должен быть 2")
	assert.Equal(t, "test1", sll.Head.Data, "Head должен быть 'test1'")
	assert.Equal(t, "test3", sll.Head.Next.Data, "Следующий элемент должен быть 'test3'")
}
func TestSinglyLinkedList_RemoveByValue_TailElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	sll.RemoveByValue("test3")
	assert.Equal(t, 2, sll.Size, "Размер списка должен быть 2")
	assert.Equal(t, "test1", sll.Head.Data, "Head должен быть 'test1'")
	assert.Equal(t, "test2", sll.Head.Next.Data, "Следующий элемент должен быть 'test2'")
}
func TestSinglyLinkedList_RemoveByValue_HeadElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	sll.RemoveByValue("test1")
	assert.Equal(t, 2, sll.Size, "Размер списка должен быть 2")
	assert.Equal(t, "test2", sll.Head.Data, "Head должен быть 'test2'")
	assert.Equal(t, "test3", sll.Head.Next.Data, "Следующий элемент должен быть 'test3'")
}
func TestSinglyLinkedList_RemoveByValue_NonExistentElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	sll.RemoveByValue("test4")
	assert.Equal(t, 3, sll.Size, "Размер списка должен остаться 3")
	assert.Equal(t, "test1", sll.Head.Data, "Head должен быть 'test1'")
	assert.Equal(t, "test3", sll.Head.Next.Next.Data, "Последний элемент должен быть 'test3'")
}
func TestSinglyLinkedList_AddAfterRemovingAllElements(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	sll.RemoveHead()
	sll.RemoveHead()
	sll.RemoveHead()

	sll.AddToTail("newTest1")
	sll.AddToTail("newTest2")

	assert.Equal(t, 2, sll.Size, "Размер списка должен быть 2")
	assert.Equal(t, "newTest1", sll.Head.Data, "Head должен быть 'newTest1'")
	assert.Equal(t, "newTest2", sll.Head.Next.Data, "Следующий элемент должен быть 'newTest2'")
}
func TestSinglyLinkedList_SaveAndLoadEmptyList(t *testing.T) {
	sll := NewSinglyLinkedList()

	err := sll.SaveToFile("empty.txt")
	assert.NoError(t, err, "Ошибка при сохранении пустого списка в файл")

	newSll := NewSinglyLinkedList()
	err = newSll.LoadFromFile("empty.txt")
	assert.NoError(t, err, "Ошибка при загрузке пустого списка из файла")

	assert.Equal(t, 0, newSll.Size, "Размер загруженного списка должен быть 0")
	assert.Nil(t, newSll.Head, "Head загруженного списка должен быть nil")

	os.Remove("empty.txt")
}
func TestSinglyLinkedList_AddAfterLoadingFromFile(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	err := sll.SaveToFile("test.txt")
	assert.NoError(t, err, "Ошибка при сохранении в файл")

	newSll := NewSinglyLinkedList()
	err = newSll.LoadFromFile("test.txt")
	assert.NoError(t, err, "Ошибка при загрузке из файла")

	newSll.AddToTail("newTest1")
	newSll.AddToTail("newTest2")

	assert.Equal(t, 5, newSll.Size, "Размер списка должен быть 5")
	assert.Equal(t, "test1", newSll.Head.Data, "Head должен быть 'test1'")
	assert.Equal(t, "newTest2", newSll.Head.Next.Next.Next.Next.Data, "Последний элемент должен быть 'newTest2'")

	os.Remove("test.txt")
}
func TestSinglyLinkedList_Search_MiddleElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	node := sll.Search("test2")
	assert.NotNil(t, node, "Узел с значением 'test2' должен быть найден")
	assert.Equal(t, "test2", node.Data, "Найденный узел должен содержать значение 'test2'")
}
func TestSinglyLinkedList_Search_TailElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	node := sll.Search("test3")
	assert.NotNil(t, node, "Узел с значением 'test3' должен быть найден")
	assert.Equal(t, "test3", node.Data, "Найденный узел должен содержать значение 'test3'")
}
func TestSinglyLinkedList_Search_HeadElement(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("test1")
	sll.AddToTail("test2")
	sll.AddToTail("test3")

	node := sll.Search("test1")
	assert.NotNil(t, node, "Узел с значением 'test1' должен быть найден")
	assert.Equal(t, "test1", node.Data, "Найденный узел должен содержать значение 'test1'")
}

func TestFindNodeByValue_sll(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("first")
	sll.AddToTail("second")
	sll.AddToTail("third")

	// Поиск существующего значения
	node := sll.Search("second")
	if node == nil || node.Data != "second" {
		t.Errorf("findNodeByValue() = %v; want node with value 'second'", node)
	}

	// Поиск несуществующего значения
	node = sll.Search("fourth")
	if node != nil {
		t.Errorf("findNodeByValue() = %v; want nil", node)
	}
}

// TestSerializeTextsll проверяет сериализацию односвязного списка в текстовый формат (JSON)
func TestSerializeTextsll(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("first")
	sll.AddToTail("second")
	sll.AddToTail("third")

	serialized, err := sll.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected := `["first","second","third"]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}

	// Сериализация пустого списка
	emptySll := NewSinglyLinkedList()
	serialized, err = emptySll.SerializeText()
	if err != nil {
		t.Errorf("SerializeText() error = %v; want nil", err)
	}

	expected = `[]`
	if serialized != expected {
		t.Errorf("SerializeText() = %v; want %v", serialized, expected)
	}
}

// TestDeserializeTextsll проверяет десериализацию односвязного списка из текстового формата (JSON)
func TestDeserializeTextsll(t *testing.T) {
	sll := NewSinglyLinkedList()
	data := `["first","second","third"]`

	err := sll.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if sll.Size != 3 || sll.Head.Data != "first" {
		t.Errorf("DeserializeText() = %v; want list with elements [first, second, third]", sll)
	}

	// Десериализация пустого списка
	emptySll := NewSinglyLinkedList()
	data = `[]`
	err = emptySll.DeserializeText(data)
	if err != nil {
		t.Errorf("DeserializeText() error = %v; want nil", err)
	}

	if emptySll.Size != 0 || emptySll.Head != nil {
		t.Errorf("DeserializeText() = %v; want empty list", emptySll)
	}

	// Десериализация некорректных данных
	invalidData := `["first","second","third"`
	err = sll.DeserializeText(invalidData)
	if err == nil {
		t.Errorf("DeserializeText() = %v; want error", err)
	}
}

// TestSerializeBinarysll проверяет сериализацию односвязного списка в бинарный формат
func TestSerializeBinarysll(t *testing.T) {
	sll := NewSinglyLinkedList()
	sll.AddToTail("first")
	sll.AddToTail("second")
	sll.AddToTail("third")

	serialized, err := sll.SerializeBinary()
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

	// Сериализация пустого списка
	emptySll := NewSinglyLinkedList()
	serialized, err = emptySll.SerializeBinary()
	if err != nil {
		t.Errorf("SerializeBinary() error = %v; want nil", err)
	}

	if len(serialized) != 0 {
		t.Errorf("SerializeBinary() = %v; want empty byte slice", serialized)
	}
}

// TestDeserializeBinarysll проверяет десериализацию односвязного списка из бинарного формата
func TestDeserializeBinarysll(t *testing.T) {
	sll := NewSinglyLinkedList()
	data := []byte{
		5, 0, 0, 0, // длина "first"
		'f', 'i', 'r', 's', 't', // строка "first"
		6, 0, 0, 0, // длина "second"
		's', 'e', 'c', 'o', 'n', 'd', // строка "second"
		5, 0, 0, 0, // длина "third"
		't', 'h', 'i', 'r', 'd', // строка "third"
	}

	err := sll.DeserializeBinary(data)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if sll.Size != 3 || sll.Head.Data != "first" {
		t.Errorf("DeserializeBinary() = %v; want list with elements [first, second, third]", sll)
	}

	// Десериализация пустого списка
	emptySll := NewSinglyLinkedList()
	emptyData := []byte{} // Пустые данные
	err = emptySll.DeserializeBinary(emptyData)
	if err != nil {
		t.Errorf("DeserializeBinary() error = %v; want nil", err)
	}

	if emptySll.Size != 0 || emptySll.Head != nil {
		t.Errorf("DeserializeBinary() = %v; want empty list", emptySll)
	}

	// Десериализация некорректных данных (недостаточно байт)
	invalidData := []byte{5, 0, 0, 0, 'f', 'i', 'r'} // Недостаточно байт для строки "first"
	err = sll.DeserializeBinary(invalidData)
	if err == nil {
		t.Errorf("DeserializeBinary() = %v; want error", err)
	}
}
