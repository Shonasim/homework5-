package main

import (
	"fmt"
)

func updateValue(p *int) {
	*p += *p
}
func swap(a, b *int) {
	*a, *b = *b, *a
}

// UpdateName обновляет имя человека
func UpdateName(name *string, newName string) {
	*name = newName
}

// UpdateAge обновляет возраст человека
func UpdateAge(age *int, newAge int) {
	*age = newAge
}

// PrintPerson выводит данные о человеке
func PrintPerson(name *string, age *int) {
	fmt.Printf("Имя: %s, Возраст: %d\n", *name, *age)
}
func square(a *int) {
	*a *= *a
	fmt.Println(&a)
}
func zero(xPtr *int) {
	*xPtr = 0
}
func Swaap(x *int, y *int) {
	*x, *y = *y, *x
}
func main() {
	//var x int = 10
	//var ptr *int
	//ptr = &x
	//fmt.Println(x)
	//fmt.Println(&x)
	//fmt.Println(ptr)
	//fmt.Println(*ptr)
	//
	//x := 5
	//fmt.Println(x)
	//updateValue(&x)
	//fmt.Println(x)
	//a, b := 2, 3
	//fmt.Println(a, b)
	//swap(&a, &b)
	//fmt.Println(a, b)

	// Инициализация данных о человеке
	//name := "Иван"
	//age := 30
	//
	//// Вывод начальных данных о человеке
	//PrintPerson(&name, &age)
	//
	//// Обновление имени и возраста
	//UpdateName(&name, "Алексей")
	//UpdateAge(&age, 35)
	//
	//// Вывод обновленных данных о человеке
	//PrintPerson(&name, &age)

	//x := 2
	//var c *int
	//c = &x
	//fmt.Println(c, "", &x)
	//*c = 3
	//fmt.Println(x)
	//x := 2
	//square(&x)
	//fmt.Println(x)
	//fmt.Println(&x)
	//
	//a := 2
	//var b *int
	//b = &a
	//*b += 5
	//fmt.Println(*b+1, a)
	//fmt.Println(*b, a)

	//x := 5
	//zero(&x)
	//fmt.Println(x) // x is 0
	a, b := 1, 2
	Swaap(&a, &b)
	fmt.Println(a, b)

}

//
