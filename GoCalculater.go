package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// цикл чтоб после ответа снова спрашивал

	for {
		userInput := userInputs()
		err := func(inputArr []string) {
			if err := checkErrors(inputArr); err != "" {
				fmt.Println(err)
				os.Exit(0)
			}
		}
		err(userInput)
	}
}

func userInputs() (str []string) { // Объявление функции userInputs, которая возвращает срез строк
	reader := bufio.NewReader(os.Stdin) // Создание нового объекта типа Reader для чтения ввода пользователя из стандартного потока ввода
	fmt.Println("Введите выражение ")   // Вывод сообщения на экран, запрашивающего ввод выражения
	text, _ := reader.ReadString('\n')  // Чтение строки из стандартного потока ввода до символа '\n'
	text = strings.TrimSpace(text)      // Удаление пробелов и других символов-разделителей в начале и конце строки
	str = strings.Split(text, " ")      // Разделение строки на подстроки по пробелам и сохранение результатов в срез строк
	return str                          // Возвращение среза строк с подстроками введенного пользователем выражения
}

// функция с всеми вариациями ошибок
func checkErrors(str []string) (err string) {
	var inputUser = &str
	var lenUser = len(*inputUser)
	var lenOperation int
	var listOperation = []string{"+", "-", "*", "/"}
	var system = systemCheck(*inputUser)
	var systemsOperation = &system
	var result = calculator(*inputUser, *systemsOperation)

	for _, i := range *inputUser {
		for _, l := range listOperation {
			if l == i {
				lenOperation += 1
			}
		}
	}

	switch {
	case lenUser <= 2:
		return "Ошибка, так как строка не является математической операцией."
	case lenOperation > 1:
		return "Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	case *systemsOperation == 3:
		return "Ошибка, так как используются одновременно разные системы счисления."
	case *systemsOperation == 0:
		return "Ошибка, так как строка не является математической операцией."
	case checkNum(*inputUser, *systemsOperation) == false:
		if *systemsOperation == 1 {
			return "Ошибка, так как формат ввода цифр не удовлетворяет заданию - от 1 до 10"
		} else if *systemsOperation == 2 {
			return "Ошибка, так как формат ввода цифр не удовлетворяет заданию - от I до X"
		}
	case result <= 0 && *systemsOperation == 2:
		return "Ошибка, так как в римской системе нет отрицательных чисел."
	}

	if *systemsOperation == 2 {
		fmt.Println(formArabInRoman(result))
	} else if *systemsOperation == 1 {
		fmt.Println(result)
	}
	return err
}

// проверка на систему счисления
func systemCheck(str []string) (i int) {
	firstArabNum := systemArabNum(str[0])
	secondArabNum := systemArabNum(str[2])
	firstRomNum := systemRomNum(str[0])
	secondRomNum := systemRomNum(str[2])

	switch {
	case firstArabNum != 0 && secondArabNum != 0:
		return 1
	case firstRomNum != 0 && secondRomNum != 0:
		return 2
	case (firstArabNum != 0 && secondRomNum != 0) ||
		(secondArabNum != 0 && firstRomNum != 0):
		return 3
	}
	return i
}

func systemArabNum(a string) int {
	firstNum, err := strconv.Atoi(a)
	if err != nil {

	}
	return firstNum
}

func systemRomNum(s string) (result int) {
	arr := map[string]int{"I": 1, "IV": 4, "V": 5, "X": 10,
		"L": 50, "C": 100, "D": 500, "M": 1000}
	result = 0

	for i := range s {
		if i < len(s)-1 && arr[s[i:i+1]] < arr[s[i+1:i+2]] {
			result -= arr[s[i:i+1]]
		} else {
			result += arr[s[i:i+1]]
		}
	}
	return result
}

func checkNum(str []string, i int) (tf bool) {
	var a int
	var b int

	if i == 1 {
		a = systemArabNum(str[0])
		b = systemArabNum(str[2])
	} else if i == 2 {
		a = systemRomNum(str[0])
		b = systemRomNum(str[2])
	}
	if (1 <= a && a <= 10) && (1 <= b && b <= 10) {
		return true
	}
	return tf
}

func calculator(str []string, i int) (result int) {
	s := func(str []string, i int) (a, b int) {
		if i == 1 { // арабская система
			a = systemArabNum(str[0])
			b = systemArabNum(str[2])
		} else if i == 2 { // римская система
			a = systemRomNum(str[0])
			b = systemRomNum(str[2])
		}
		return a, b
	}

	a, b := s(str, i)
	var operand = &str[1]
	switch *operand {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = a / b
	}
	return result
}

func formArabInRoman(i int) (result string) {
	var romNums = map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}
	var k = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 40, 50, 90, 100}
	var numbers []int

	for i != 0 {
		for _, a := range k {
			if a == i {
				i -= a
				numbers = append(numbers, a)
			} else if a < i {
				i -= a
				numbers = append(numbers, a)
				break
			}
		}
	}

	var rom string

	for _, k := range numbers {
		for key, value := range romNums {
			if key != k {
				continue
			} else if key == k {
				rom += value
			}
		}
	}

	result += rom
	return result
}
