package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() (string, string, interface{}) {
	reader := bufio.NewReader(os.Stdin) //читаем ввод с консоли, os.Stdin константа, представляющая поток ввода (консоль)
	fmt.Print("Введите: ")
	input, _ := reader.ReadString('\n') //сохраняем ввод в переменную, до начала новой строки
	input = strings.TrimSpace(input)    //Удаляем пробелы с начала и конца строки

	// Проверка на кавычки
	if input[0] != '"' {
		panic("Ошибка: Первый аргумент должен быть в кавычках.")
	}

	// Разделяем строку по пробелам, игнорируя пробелы в кавычках
	parts := []string{}          //объявляем срез который будет хранить разделимые части строки
	inQuotes := false            //переменная типа булево которая будет определять находится ли символ внутри кавычек или нет
	currentPart := ""            //переменная будет хранить текущую часть строки
	for _, char := range input { // пробегаемся по каждому символу input
		if char == '"' { //тут если текущий символ равен "
			inQuotes = !inQuotes //то значение inQuotes меняется на противоположное
		} else if char == ' ' && !inQuotes { //иначе если символ равен пробелу и не находится внутри кавычек то выполняем действие
			if currentPart != "" {
				parts = append(parts, currentPart) //Если текущая часть не пуста то добавляем в срез parts, а currentPart очищаем
				currentPart = ""
			}
		} else {
			currentPart += string(char) //В противном случае если символ не пробел, или находится внутри кавычек, добавляем к currentPart
		}
	}
	if currentPart != "" {
		parts = append(parts, currentPart) //Если после цикла currentPart не пуста, то она добавляется в срез
	}
	//Проверяем что бы по итогу у нас было 3 части
	if len(parts) < 3 {
		fmt.Println("Неверный формат ввода: не хватает элементов")
		return "", "", nil
	}

	// Извлекаем значения
	str1 := parts[0]
	operator := parts[1]
	str2 := parts[2]

	// Проверка длины строки в кавычках
	if len(str1) > 10 {
		panic("Ошибка: Входящая строка в кавычках превышает 10 символов.")
	}
	if len(str2) > 10 {
		panic("Ошибка: Входящая строка в кавычках превышает 10 символов.")
	}

	// Удаляем кавычки из str1
	str1 = strings.Trim(str1, "\"")

	// Проверяем, является ли str2 числом
	var value interface{}                           //объявляем переменную любого типа
	if num, err := strconv.Atoi(str2); err == nil { //пытаемся преобразовать строку в целое число, если получается
		// Проверка на диапазон числа
		if num >= 1 && num <= 10 {
			value = num
		} else {
			// Число вне диапазона 1-10, обрабатываем как строку
			value = str2
		}
	} else {
		// Если str2 не число, то это строка
		value = str2
	}

	return str1, operator, value
}

func Calculate() {
	for { // Добавляем бесконечный цикл
		str1, operatorStr, str2Interface := ReadInput()

		switch operatorStr {
		case "+":
			if str2, ok := str2Interface.(string); ok {
				result := str1 + str2
				// Обрезаем результат до 40 символов
				if len(result) > 40 {
					result = result[:40] + "..."
				}
				fmt.Printf("\"%s\"\n", result)
			} else if num, ok := str2Interface.(int); ok {
				// Если str2 - число, преобразуем его в строку
				result := str1 + strconv.Itoa(num)
				// Обрезаем результат до 40 символов
				if len(result) > 40 {
					result = result[:40] + "..."
				}
				fmt.Printf("\"%s\"\n", result)
			} else {
				panic("Неверный формат ввода")
			}
		case "-":
			// Обработка вычитания строк (удалить str2 из str1)
			if str2, ok := str2Interface.(string); ok {
				result := strings.ReplaceAll(str1, str2, "") // Используем ReplaceAll
				// Обрезаем результат до 40 символов
				if len(result) > 40 {
					result = result[:40] + "..."
				}
				fmt.Printf("\"%s\"\n", result)
			} else {
				panic("Неверный формат ввода")
			}
		case "*":
			// Повторение строки (умножение на число)
			if num, ok := str2Interface.(int); ok {
				result := strings.Repeat(str1, num)
				// Обрезаем результат до 40 символов
				if len(result) > 40 {
					result = result[:40] + "..."
				}
				fmt.Printf("\"%s\"\n", result)
			} else {
				panic("Число должно быть от 1 до 10")
			}
		case "/":
			// Взятие подстроки
			if num, ok := str2Interface.(int); ok {
				if num > len(str1) {
					panic("Число не может превышать длину строки")
				}
				result := str1[:num]
				// Обрезаем результат до 40 символов
				if len(result) > 40 {
					result = result[:40] + "..."
				}
				fmt.Printf("\"%s\"\n", result)
			} else {
				panic("Неверный формат ввода")
			}
		default:
			panic("Неподдерживаемый оператор")
		}
	}
}

func main() {
	Calculate()
}
