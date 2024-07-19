package main

import (
	"fmt"
)

// Задача 1:
func task1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Задача 2:
func task2() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}
}

// Задача 3:
func task3() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}
}

// Задача 4:
func task4() {
	a, b := 0, 1
	for i := 1; i <= 10; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

// Задание 5:
func task5() {
	a := 56
	b := 98
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Println("НОД:", a)
}

// Задача 6: FizzBuzz
func task6() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Задача 7: Проверка на простое число
func task7() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	isPrime := true
	if num <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Println(num, "простое число")
	} else {
		fmt.Println(num, "не является простым числом")
	}
}

// Задача 8: Вывод всех делителей числа
func task8() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	fmt.Print("Делители: ")
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// Задача 9: Сумма цифр числа
func task9() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	fmt.Println("Сумма цифр:", sum)
}

// Задача 10: Запрос положительного числа
func task10() {
	var num int
	for {
		fmt.Print("Введите положительное число: ")
		fmt.Scan(&num)
		if num > 0 {
			break
		}
	}
	fmt.Println("Вы ввели:", num)
}

// Задача 11: Произведение чисел до n
func task11() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)
	prod := 1
	for i := 1; i <= n; i++ {
		prod *= i
		if prod > 1000 {
			fmt.Println("Произведение превысило 1000")
			return
		}
	}
	fmt.Println("Произведение:", prod)
}

// Задача 12: Проверка на палиндром
func task12() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	original := num
	reversed := 0
	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	if original == reversed {
		fmt.Println(original, "является палиндромом")
	} else {
		fmt.Println(original, "не является палиндромом")
	}
}

// Задача 13: Пирамида из символов '*'
func task13() {
	var n int
	fmt.Print("Введите высоту пирамиды: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Задача 14: Таблица умножения до n
func task14() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

// Задача 15: Треугольник Паскаля
func task15() {
	var n int
	fmt.Print("Введите высоту треугольника Паскаля: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		number := 1
		for j := 0; j <= i; j++ {
			fmt.Printf("%4d", number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}

// Задача 16: Факториал до отрицательного числа
func task16() {
	for {
		var num int
		fmt.Print("Введите число: ")
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		fact := 1
		for i := 1; i <= num; i++ {
			fact *= i
		}
		fmt.Println("Факториал:", fact)
	}
}

// Задача 17: Сумма двух чисел
func task17() {
	for {
		var a, b int
		fmt.Print("Введите два числа: ")
		fmt.Scan(&a, &b)
		fmt.Println("Сумма:", a+b)
	}
}

// Задача 18: Числа от 1 до 100, пропуск квадратов
func task18() {
	for i := 1; i <= 100; i++ {
		isSquare := false
		for j := 1; j*j <= i; j++ {
			if j*j == i {
				isSquare = true
				break
			}
		}
		if !isSquare {
			fmt.Println(i)
		}
	}
}

// Задача 19: Простые числа от 1 до 50
func task19() {
	for i := 2; i <= 50; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}

// Задача 20: Числа от 1 до 30, кроме делящихся на 2 или 3
func task20() {
	for i := 1; i <= 30; i++ {
		if i%2 != 0 && i%3 != 0 {
			fmt.Println(i)
		}
	}
}

// Задача 21: Числа от 1 до 50, прекращение на кубе
func task21() {
	for i := 1; i <= 50; i++ {
		for j := 1; j*j*j <= i; j++ {
			if j*j*j == i {
				return
			}
		}
		fmt.Println(i)
	}
}

// Задача 22: Числа от 1 до 100, прекращение при сумме > 200
func task22() {
	sum := 0
	for i := 1; i <= 100; i++ {
		if sum+i > 200 {
			return
		}
		fmt.Println(i)
		sum += i
	}
}

// Задача 23: Прекращение при вводе числа, кратного 7
func task23() {
	for {
		var num int
		fmt.Print("Введите число: ")
		fmt.Scan(&num)
		if num%7 == 0 {
			break
		}
	}
}

func main() {
	// Выполнение каждой задачи по очереди
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9()
	task10()
	task11()
	task12()
	task13()
	task14()
	task15()
	task16()
	task17()
	task18()
	task19()
	task20()
	task21()
	task22()
	task23()
}
