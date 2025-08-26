package main

import "fmt"

/* func main() {
	// 1-я секция -- объявление переменной
	// i := 1

	// 2-я секция -- условие, при котором будет выаполняться цикл
	// i <= 5

	// 3-я секция -- изменение от одной итерации к другой
	// i++
	// fmt.Println("До")

	number := 999

	for i := 1; i <= 5; i++ {
		score := 5
		fmt.Println("Hello, World!")
		fmt.Println("Lampochka", score)
		fmt.Println("Number:", number)

		score += 8 // не работает после завершения цикла. цикл начинается со старой переменной
	}
	// fmt.Println("После", score) // у переменных есть область видимости

	// "math/rand"
	// "time"

	// 1. бесконечные циклы
	// 2. ключевое слово break

}
 */

/* import (
	// "fmt"
	"math/rand"
	
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Рандомная цифра:", rand.Intn(10))

	// 0 1
	if rand.Intn(2) == 1 {
		fmt.Println("Вы попали в 1")
	} else {
		fmt.Println("Вы НЕ попали в 1")
	}
} */

import (
	// "fmt"
	"time"
	"math/rand"
)

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Новая итерация! i:", i)
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// бесконечный цикл
	score := 0

	for {
		fmt.Println("Новая итерация! i:", score)
		time.Sleep(1 * time.Second)
		score++

		// 0 1 2 3 4
		// 1
		if rand.Intn(4) == 1 {
			fmt.Println("============")
			fmt.Println("Счёт окончен")
			fmt.Println("============")
			break // прерывет цикл
		}
	}
}
