package main

import "fmt"

//=============================================================================================//

func main() {
	number := 10

	fmt.Println(&number) // &number - указатель на переменную number
}

//=============================================================================================//

/* func main() {
	var number int = 10

	var pointer *int = &number // *int - указатель на тип данных int

	fmt.Println(number)
	fmt.Println(pointer)
} */

//=============================================================================================//

/* func main() {
	number := 10
	
	pointer := &number // синтаксис указателя на переменную number 

	//fmt.Println(pointer) // Показывает просто адрес в памяти
	foo(pointer)
}

// функция будет ожидать аргумент n, который является указателем на
// тип данных int
func foo(n *int) {
	fmt.Println(n)

	// разыменование указателя(получаем доступ к переменной, на которую 
	// этот указатель указывает
	fmt.Println(*n)
} */

//=============================================================================================//

// Без указателя
/* func main() {
	number := 5

	boo(number)

	fmt.Println(number)
}

func boo(n int) {
	n = 10 // Просто поменяли переменную n, которая существует только в рамках этой функции. 
		   // Поменяли коппию переменной number, а не саму переменную.
	fmt.Println(n)
} */

//=============================================================================================//

// Изменение значения переменной через указатель
// Указатели в main и foo - это два разных указателя на один и тот же адрес в памяти
/* func main() {
	number := 10
	
	pointer := &number

	foo(pointer)

	fmt.Println(number)
	//fmt.Println("Разыменовывание указателя pointer:", *pointer)	// разыменование через оригинальный указатель
}
// Передали в функцию указатель на number
// Когда мы передаём указатель в функцию, то мы создаём КОППИЮ этого указателя, с тем же адресом в памяти
func foo(n *int) {
	fmt.Println("<--Начало переменной foo-->")
	fmt.Println(n)
	fmt.Println(*n) // Значение переменной, которая под этим указателем

	*n = 5 // получаем доступ к переменной через *(указатель) и меняем её
	fmt.Println("<--Конец переменной foo-->")
} */

//=============================================================================================//

/* func main() {
	name := "Игорь"

	pointer := &name

	fmt.Println("Имя до изменения:", name)

	changeName(pointer)

	fmt.Println("Имя после изменения:", name)
}

func changeName(s *string) {
	*s = "Иван"
} */

//=============================================================================================//

/* func main() {
	name := "Игорь"

	pointer := &name

	fmt.Println("Имя до изменения:", name)

	changeName(pointer)

	fmt.Println("Имя после изменения:", name)
}

// этот указатель будет ссылаться уже на другую переменную
func changeName(s *string) {
	str := "hello"
	*s = &str
} */

//=============================================================================================//

// nil указатели - указатели, которые никуда не указывают

/* func main(){
	number := 15

	var ptr *int

	fmt.Println("number:", number)
	fmt.Println("ptr:", ptr)

	// fmt.Println("Разыменование nil-указателя:", *ptr) // будет ошибка

	if ptr != nil {
		fmt.Println("Разыменование:", *ptr)
	} else {
		fmt.Println("Получен nil-указатель")
	}
} */

// Указатели нужны для того, чтобы работать с одним и тем же объектом в памяти в разных частях программы
// Например работа с подключением к базе данных
// Когда мы что-то передаём в функцию - оно копируется







































/* func main() { 
	coffee := "Espresso" 
	pointer := &coffee // Создаем указатель pointer, который хранит адрес переменной coffee

	fmt.Println("Название кофе для переменной coffee:", coffee)
	// Выводим адрес в памяти, где хранится переменная coffee
	fmt.Println("Место в памяти:", pointer)
	// Выводим адрес указателя pointer
	fmt.Println("Poiner address:", pointer)
	fmt.Println("=========================================")

	coffeeCoppy := coffee // Создаем новую переменную coffeeCoppy и присваиваем ей значение переменной coffee
	
	fmt.Println("Название кофе для переменнной coffeeCoppy:", coffeeCoppy)
	// Выводим адрес в памяти, где хранится переменная coffeeCoppy
	fmt.Println("Место в памяти:", &coffeeCoppy)
	// Выводим адрес указателя на переменную coffeeCoppy
	fmt.Println("Poiner address:", &coffeeCoppy)
} */

/* func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price:", coffeePrice)
	// STEP 1
	// Compile Time (code your write):	 var coffeePrice = 4.50
	// Runtime (what maschine sees):	[some memory address] = 4.50

	// STEP 2
	// Compile Time (code your write):	fmt.Println("Coffee price:",coffeePrice)
	// Runtime (what maschine sees):	fmt.Println([some memory address], [memory address (same as on step 1)]
	fmt.Println("Memory addres of price 4.50", &coffeePrice)

	coffeePrice = 5.00
	fmt.Println("Memory addres of price 5.00", &coffeePrice)

	// poinerCoffePrice := &coffeePrice
	var pointerToCoffePrice *float64 = &coffeePrice // Эта переменная содержит указатель на место в памяти, где хранится знач. типа float64
	*pointerToCoffePrice = 7.50 // Доступ к месту в памяти для изменения значения(без * мы не получим доступ, т.к. это не переменная а ссылка)

	fmt.Println("Updated coffeePrice value in memory", coffeePrice)
} */