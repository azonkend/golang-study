package main

import "fmt"
// defer вызов анонимной функции
// Анонирмная функция выполняется в самом конце(отложенный до конца вызов)
/* func main() {
	fmt.Println("Я main и я начался")
	defer func() {
		fmt.Println("Я main и я закончился")
	}()

	//hello()	
	foo()
}

func hello() {
	fmt.Println("Я hello и я начался")
	defer func() {
		fmt.Println("Я hello и я закончился")
	}()

	fmt.Println("Я функция hello 1")
	fmt.Println("Я функция hello 2")
	fmt.Println("Я функция hello 3")
}

// Анонимная функция кладётся в stack отложенных вызовов
// defer кладутся в stack сверху вниз, а достаются в обратном порядке 
func foo() {
	defer func() {
		fmt.Println("Defer 1")
	}()

	defer func() {
		fmt.Println("Defer 2")
	}()

	defer func() {
		fmt.Println("Defer 3")
	}()

	fmt.Println("foo")
 }
*/
//=============================================================================================//

func main() {
	fmt.Println("Я main и я начался")
	defer func() {
		fmt.Println("Я main и я закончился")
	}()

	database()
}

func database() {
	// Создать подключение
	defer func() { // с помощью этой функии мы можем закрыть подключение на любом этапе кода, чтобы не закрывать в каждой фунции вручную
		// Закрыть подключение
	}()
	// функции
	// циклы

	// Закрыть подключение
}