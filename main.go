package main

import "fmt"

/* func main() {
	number := 10
	// синтаксис указателя на переменную number 
	pointer := &number

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

func main() {
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
}