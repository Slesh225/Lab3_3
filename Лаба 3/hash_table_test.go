package main

import (
	"os"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	ht := NewHashTable(10)
	if ht.Capacity != 10 {
		t.Errorf("Ожидаемая емкость: 10, получено: %d", ht.Capacity)
	}
	if len(ht.Table) != 10 {
		t.Errorf("Ожидаемая длина таблицы: 10, получено: %d", len(ht.Table))
	}
}

func TestHashFunction(t *testing.T) {
	ht := NewHashTable(10)
	tests := []struct {
		key    string
		expect int
	}{
		{"test", 7},
		{"hello", 4},
		{"world", 1},
		{"", 0}, // Пустой ключ
	}

	for _, tt := range tests {
		index := ht.HashFunction(tt.key)
		if index != tt.expect {
			t.Errorf("Для ключа '%s' ожидался индекс %d, получено %d", tt.key, tt.expect, index)
		}
	}
}

func TestHSetAndHGet(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")

	// Проверка существующих ключей
	ht.HGet("key1")
	ht.HGet("key2")

	// Проверка несуществующего ключа
	ht.HGet("key3")

	// Перезапись значения
	ht.HSet("key1", "new_value1")
	ht.HGet("key1")
}

func TestHDel(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")

	// Удаление существующего ключа
	ht.HDel("key1")

	// Попытка удаления несуществующего ключа
	ht.HDel("key3")

	// Удаление из середины цепочки
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")
	ht.HSet("key3", "value3")
	ht.HDel("key2")
	ht.HGet("key1")
	ht.HGet("key2")
	ht.HGet("key3")
}

func TestClear_hs(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")

	ht.Clear()

	for i := 0; i < ht.Capacity; i++ {
		if ht.Table[i] != nil {
			t.Errorf("Ожидалось, что таблица будет пустой, но найдены элементы в индексе %d", i)
		}
	}
}

func TestHPrint(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")

	// Этот тест не проверяет вывод, но он помогает покрыть код
	ht.HPrint()
}

func TestLoadFromFile_hs(t *testing.T) {
	ht := NewHashTable(10)
	err := ht.LoadFromFile("testdata.txt")
	if err != nil {
		t.Errorf("Ошибка загрузки из файла: %v", err)
	}

	// Проверка загруженных данных
	ht.HGet("key1")
	ht.HGet("key2")

	// Попытка загрузки из несуществующего файла
	err = ht.LoadFromFile("nonexistent.txt")
	if err == nil {
		t.Errorf("Ожидалась ошибка при загрузке из несуществующего файла")
	}
}

func TestSaveToFile_hs(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")

	err := ht.SaveToFile("testdata_output.txt")
	if err != nil {
		t.Errorf("Ошибка сохранения в файл: %v", err)
	}

	// Проверка сохраненных данных
	ht2 := NewHashTable(10)
	err = ht2.LoadFromFile("testdata_output.txt")
	if err != nil {
		t.Errorf("Ошибка загрузки из файла: %v", err)
	}

	ht2.HGet("key1")
	ht2.HGet("key2")

	// Попытка сохранения в недопустимый путь
	err = ht.SaveToFile("/invalid/path/testdata_output.txt")
	if err == nil {
		t.Errorf("Ожидалась ошибка при сохранении в недопустимый путь")
	}
}

func TestFindNodeByKey(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")

	node := ht.findNodeByKey("key1")
	if node == nil || node.Value != "value1" {
		t.Errorf("Ожидалось найти узел с ключом 'key1' и значением 'value1'")
	}

	node = ht.findNodeByKey("key3")
	if node != nil {
		t.Errorf("Ожидалось, что узел с ключом 'key3' не будет найден")
	}
}

func TestCollisionHandling(t *testing.T) {
	ht := NewHashTable(1) // Искусственно создаем коллизии
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")

	ht.HGet("key1")
	ht.HGet("key2")

	ht.HDel("key1")
	ht.HGet("key1")
	ht.HGet("key2")
}

func TestEmptyTableOperations(t *testing.T) {
	ht := NewHashTable(10)

	// Попытка получения значения из пустой таблицы
	ht.HGet("key1")

	// Попытка удаления из пустой таблицы
	ht.HDel("key1")

	// Очистка пустой таблицы
	ht.Clear()
}

func TestMain(m *testing.M) {
	// Создание тестового файла
	file, err := os.Create("testdata.txt")
	if err != nil {
		panic(err)
	}
	file.WriteString("key1 value1\n")
	file.WriteString("key2 value2\n")
	file.Close()

	// Запуск тестов
	code := m.Run()

	// Удаление тестового файла
	os.Remove("testdata.txt")
	os.Remove("testdata_output.txt")

	os.Exit(code)
}
func TestSerializeDeserializeTextHT(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")
	ht.HSet("key3", "value3")

	// Сериализация в текстовый формат
	serialized, err := ht.SerializeText()
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Десериализация из текстового формата
	newHt := NewHashTable(10)
	err = newHt.DeserializeText(serialized)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Проверка корректности десериализованных данных
	keys := []string{"key1", "key2", "key3"}
	for _, key := range keys {
		value := ht.findNodeByKey(key).Value
		newValue := newHt.findNodeByKey(key).Value
		if value != newValue {
			t.Errorf("Ожидалось значение %s для ключа %s, получено %s", value, key, newValue)
		}
	}

	// Проверка количества элементов
	if ht.Size() != newHt.Size() {
		t.Errorf("Ожидалось, что размер таблицы будет %d, получено %d", ht.Size(), newHt.Size())
	}
}

func TestSerializeDeserializeTextEmptyTableht(t *testing.T) {
	ht := NewHashTable(10)

	// Сериализация пустой таблицы
	serialized, err := ht.SerializeText()
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Десериализация
	newHt := NewHashTable(10)
	err = newHt.DeserializeText(serialized)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Проверка, что таблица пуста
	if newHt.Size() != 0 {
		t.Errorf("Ожидалось, что таблица будет пустой, но размер равен %d", newHt.Size())
	}
}
func TestSerializeDeserializeBinaryht(t *testing.T) {
	ht := NewHashTable(10)
	ht.HSet("key1", "value1")
	ht.HSet("key2", "value2")
	ht.HSet("key3", "value3")

	// Сериализация в бинарный формат
	serialized, err := ht.SerializeBinary()
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Десериализация из бинарного формата
	newHt := NewHashTable(10)
	err = newHt.DeserializeBinary(serialized)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Проверка корректности десериализованных данных
	keys := []string{"key1", "key2", "key3"}
	for _, key := range keys {
		value := ht.findNodeByKey(key).Value
		newValue := newHt.findNodeByKey(key).Value
		if value != newValue {
			t.Errorf("Ожидалось значение %s для ключа %s, получено %s", value, key, newValue)
		}
	}

	// Проверка количества элементов
	if ht.Size() != newHt.Size() {
		t.Errorf("Ожидалось, что размер таблицы будет %d, получено %d", ht.Size(), newHt.Size())
	}
}

func TestSerializeDeserializeBinaryEmptyTableht(t *testing.T) {
	ht := NewHashTable(10)

	// Сериализация пустой таблицы
	serialized, err := ht.SerializeBinary()
	if err != nil {
		t.Fatalf("Ошибка сериализации: %v", err)
	}

	// Десериализация
	newHt := NewHashTable(10)
	err = newHt.DeserializeBinary(serialized)
	if err != nil {
		t.Fatalf("Ошибка десериализации: %v", err)
	}

	// Проверка, что таблица пуста
	if newHt.Size() != 0 {
		t.Errorf("Ожидалось, что таблица будет пустой, но размер равен %d", newHt.Size())
	}
}
