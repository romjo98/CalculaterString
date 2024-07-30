package operations

import (
	"calculatorstrings/read"
	"fmt"
	"strconv"
	"strings"
)

func Calculate() {
	for { // Добавляем бесконечный цикл
		str1, operatorStr, str2Interface := read.ReadInput()

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
