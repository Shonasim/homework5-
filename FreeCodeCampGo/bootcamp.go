package main

// Функции
func sub(x int, y int) int {
	return x - y
}
func concat(n1, n2 string) string {
	return n1 + n2
}

// Структуры
type messagetosend struct {
	message   string
	sender    user
	recepient user
	// Wheel is a field containing an anonymous struct
}
type user struct {
	name   string
	number int
}

func main() {
	//Задача1
	//x := sub(3, 2)
	//fmt.Println(x)

	//Задача2
	//a := concat("hello", "world")
	//fmt.Println(a)

}
