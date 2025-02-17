package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processQuery(query string, array *Array, stack *Stack, queue *Queue, singlyList *SinglyLinkedList, doublyList *DoublyLinkedList, hashTable *HashTable, cbTree *BinaryTree, filename string) {
	tokens := strings.Split(query, " ")

	switch tokens[0] {
	case "MPUSH":
		if len(tokens) == 3 {
			index, _ := strconv.Atoi(tokens[1])
			value := tokens[2]
			array.Add(index, value)
		} else {
			fmt.Println("Ошибка: команда MPUSH требует 2 аргумента.")
		}
	case "MDEL":
		if len(tokens) == 2 {
			index, _ := strconv.Atoi(tokens[1])
			array.Remove(index)
		} else {
			fmt.Println("Ошибка: команда MDEL требует 1 аргумент.")
		}
	case "MGET":
		if len(tokens) == 2 {
			index, _ := strconv.Atoi(tokens[1])
			value := array.Get(index)
			fmt.Printf("Элемент по индексу %d: %s\n", index, value)
		} else {
			fmt.Println("Ошибка: команда MGET требует 1 аргумент.")
		}
	case "MREPLACE":
		if len(tokens) == 3 {
			index, _ := strconv.Atoi(tokens[1])
			value := tokens[2]
			array.Replace(index, value)
		} else {
			fmt.Println("Ошибка: команда MREPLACE требует 2 аргумента.")
		}
	case "SERT": // Новая команда SERT
		// Сериализация в текстовый формат
		serializedData, err := array.SerializeText()
		if err != nil {
			fmt.Println("Ошибка сериализации:", err)
		} else {
			// Сохраняем сериализованные данные в файл
			err = os.WriteFile(filename, []byte(serializedData), 0644)
			if err != nil {
				fmt.Println("Ошибка записи в файл:", err)
			} else {
				fmt.Println("Данные сериализованы и сохранены в файл.")
			}

			// Теперь десериализуем данные из файла
			loadedData, err := os.ReadFile(filename)
			if err != nil {
				fmt.Println("Ошибка чтения из файла:", err)
			} else {
				// Десериализация из текстового формата
				err = array.DeserializeText(string(loadedData))
				if err != nil {
					fmt.Println("Ошибка десериализации:", err)
				} else {
					fmt.Println("Данные успешно десериализованы из файла.")
				}
			}
		}

	case "SPUSH":
		if len(tokens) == 2 {
			value := tokens[1]
			stack.Push(value)
		} else {
			fmt.Println("Ошибка: команда SPUSH требует 1 аргумент.")
		}
	case "SPOP":
		stack.Pop()
	case "QPUSH":
		if len(tokens) == 2 {
			value := tokens[1]
			queue.Push(value)
		} else {
			fmt.Println("Ошибка: команда QPUSH требует 1 аргумент.")
		}
	case "QPOP":
		queue.Pop()
	case "LSADDHEAD":
		if len(tokens) == 2 {
			value := tokens[1]
			singlyList.AddToHead(value)
		} else {
			fmt.Println("Ошибка: команда LSADDHEAD требует 1 аргумент.")
		}
	case "LSADDTAIL":
		if len(tokens) == 2 {
			value := tokens[1]
			singlyList.AddToTail(value)
		} else {
			fmt.Println("Ошибка: команда LSADDTAIL требует 1 аргумент.")
		}
	case "LSDELHEAD":
		singlyList.RemoveHead()
	case "LSDELTAIL":
		singlyList.RemoveTail()
	case "LSDELVALUE":
		if len(tokens) == 2 {
			value := tokens[1]
			singlyList.RemoveByValue(value)
		} else {
			fmt.Println("Ошибка: команда LSDELVALUE требует 1 аргумент.")
		}
	case "LDADDHEAD":
		if len(tokens) == 2 {
			value := tokens[1]
			doublyList.AddToHead(value)
		} else {
			fmt.Println("Ошибка: команда LDADDHEAD требует 1 аргумент.")
		}
	case "LDADDTAIL":
		if len(tokens) == 2 {
			value := tokens[1]
			doublyList.AddToTail(value)
		} else {
			fmt.Println("Ошибка: команда LDADDTAIL требует 1 аргумент.")
		}
		if query == "SERIALIZE" {
			serializedData, err := doublyList.SerializeText()
			if err != nil {
				fmt.Println("Ошибка сериализации:", err)
			} else {
				fmt.Println("Сериализованные данные:", serializedData)
			}
		}

	case "LDDELHEAD":
		doublyList.RemoveFromHead()
	case "LDDELTAIL":
		doublyList.RemoveFromTail()
	case "LDDELVALUE":
		if len(tokens) == 2 {
			value := tokens[1]
			doublyList.RemoveByValue(value)
		} else {
			fmt.Println("Ошибка: команда LDDELVALUE требует 1 аргумент.")
		}
	case "HSET":
		if len(tokens) == 3 {
			key := tokens[1]
			value := tokens[2]
			hashTable.HSet(key, value)
		} else {
			fmt.Println("Ошибка: команда HSET требует 2 аргумента.")
		}
	case "HGET":
		if len(tokens) == 2 {
			key := tokens[1]
			hashTable.HGet(key)
		} else {
			fmt.Println("Ошибка: команда HGET требует 1 аргумент.")
		}
	case "HDEL":
		if len(tokens) == 2 {
			key := tokens[1]
			hashTable.HDel(key)
		} else {
			fmt.Println("Ошибка: команда HDEL требует 1 аргумент.")
		}
	case "HPRINT":
		hashTable.HPrint()
	case "TINSERT":
		if len(tokens) == 2 {
			digit, _ := strconv.Atoi(tokens[1])
			cbTree.Insert(digit)
		} else {
			fmt.Println("Ошибка: команда TINSERT требует 1 аргумент.")
		}
	case "TISCBT":
		if cbTree.IsComplete() {
			fmt.Println("Дерево является полным двоичным деревом.")
		} else {
			fmt.Println("Дерево не является полным двоичным деревом.")
		}
	case "TFIND":
		if len(tokens) == 2 {
			value, _ := strconv.Atoi(tokens[1])
			if cbTree.FindValue(value) {
				fmt.Printf("Значение %d найдено в дереве.\n", value)
			} else {
				fmt.Printf("Значение %d не найдено в дереве.\n", value)
			}
		} else {
			fmt.Println("Ошибка: команда TFIND требует 1 аргумент.")
		}
	case "TDISPLAY":
		cbTree.Display()
	case "PRINT":
		array.Print()
		stack.Print()
		queue.Print()
		singlyList.Print()
		doublyList.Print()
		hashTable.HPrint()
	default:
		fmt.Printf("Неизвестная команда: %s\n", tokens[0])
	}
}

func main() {
	var query, filename string
	array := NewArray(10)
	stack := NewStack()
	queue := NewQueue()
	singlyList := NewSinglyLinkedList()
	doublyList := NewDoublyLinkedList()
	hashTable := NewHashTable(10)
	cbTree := NewBinaryTree()

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--file" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
			i++
		}
		if arg == "--query" && i+1 < len(os.Args) {
			query = os.Args[i+1]
			i++
		}
	}

	if filename != "" && query != "" {
		command := strings.Split(query, " ")[0]

		switch command[0] {
		case 'M':
			array.LoadFromFile(filename)
		case 'S':
			stack.LoadFromFile(filename)
		case 'Q':
			queue.LoadFromFile(filename)
		case 'L':
			if command[1] == 'S' {
				singlyList.LoadFromFile(filename)
			} else if command[1] == 'D' {
				doublyList.LoadFromFile(filename)
			}
		case 'H':
			hashTable.LoadFromFile(filename)
		case 'T':
			cbTree.LoadFromFile(filename)
		case 'P':
			if command == "PRINT" {
				file, err := os.Open(filename)
				if err != nil {
					fmt.Printf("Ошибка: не удалось открыть файл %s\n", filename)
					return
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					fmt.Println(scanner.Text())
				}
			}
		default:
			fmt.Println("Ошибка: нераспознанный тип команды.")
			return
		}
	}

	if query != "" {
		processQuery(query, array, stack, queue, singlyList, doublyList, hashTable, cbTree, filename)
	} else {
		fmt.Println("Ошибка: запрос не указан.")
		return
	}

	if filename != "" && query != "" {
		command := strings.Split(query, " ")[0]

		switch command[0] {
		case 'M':
			array.SaveToFile(filename)
		case 'S':
			stack.SaveToFile(filename)
		case 'Q':
			queue.SaveToFile(filename)
		case 'L':
			if command[1] == 'S' {
				singlyList.SaveToFile(filename)
			} else if command[1] == 'D' {
				doublyList.SaveToFile(filename)
			}
		case 'H':
			hashTable.SaveToFile(filename)
		case 'T':
			cbTree.SaveToFile(filename)
		}
	}
}
