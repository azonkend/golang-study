package main

import "fmt"

/* func main() {
	score := 5
	fmt.Println("Счёт:", score)
	fmt.Println("До вызова функции")
	hello("Кыся")
	hello("Пёся")
	fmt.Println("После вызова функции")
}

// функции могут принимать аргументы с указанием типа. можно записывать через запятую
func hello(arg string) {
	fmt.Println("Hello 1")
	fmt.Println("Hello 2")
	fmt.Println("Hello 3")
	fmt.Println("Hello 4")
	fmt.Println("Указанный аргумент функции:", arg)
} */

//=============================================================================================//

/* func main() {
	number := 5
	text := "Привет"

	fmt.Println("number:", number)
	fmt.Println("text:", text)

	foo(number, text) // копирование переменных из main

	fmt.Println("number:", number)
	fmt.Println("text:", text)
} 

func foo(n int, t string) {
	fmt.Println("<Начало функции foo>")
	fmt.Println("n:", n)
	fmt.Println("t:", t)

	n = 100
	t = "Меня поменяли :)"

	fmt.Println("n:", n)
	fmt.Println("t:", t)
	fmt.Println("<Конец функции foo>")
} */

//=============================================================================================//

/* func main() {
	fmt.Println("До вызова функции")
	number := sum(1, 3)
	fmt.Println("После вызова функции")
	fmt.Println("result:", number)

	text1 := "Hello "
	text2 := "World!"
	text3 := sumString(text1, text2)
	fmt.Println(text3)
}

func sum(a int, b int) int { // int вне скобок, это тот тип, который мы должны вернуть из функции
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	s := a + b

	return s
}

func sumString(s1 string, s2 string) string {
	s := s1 + s2 

	return s
} */

//=============================================================================================//

/* func main() {
	s := str()

	fmt.Println(s)
}

func str() string {
	subscribed := false

	if subscribed {
		return "Ты трушный"
	} else {
		return "Ты фолсный"
	}
} */

//=============================================================================================//

var number int = 5 // глобальная переменная, доступна во всех функциях(обычно задаются значения в разных частях программы)


func main() {
	greeting("")
}
// return в функции, которая ничего не возвращает
func greeting(name string) {
	if name == "" {
		fmt.Println("Вы передали пустое имя ...")
		return
	}

	fmt.Println("Привет, уважаемый", name)
}

func square(n int) int {
	s := n * n

	return s
}