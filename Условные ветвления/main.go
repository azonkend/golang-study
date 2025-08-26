package main

import "fmt"

/* func main() {
	score := 5

	if score > 10 {
		fmt.Println("Ты красавчик!")
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}
	
} */

/* func main() {
	score := 5

	if score > 10 {
		if score > 15 {
			fmt.Println("Ты мега-красавчик!")
		} else {
			fmt.Println("Ты красавчик!")
		}
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}
	
} */

/* func main() {
	score := 5

	// сократили вложенность
	if score > 15 {
		fmt.Println("Ты мега-красавчик!")
	} else if score > 10 {
		fmt.Println("Ты красавчик!") 
	} else {
		fmt.Println("Тебе нужно ещё многому научиться")
	}
		
} */

/* func main() {
	score := 20
	
	if score == 12 {
		fmt.Println("Дюжина!")
	} else if score == 21 {
		fmt.Println("Очко!")
	} else if score == 50 {
		fmt.Println("Полтинник!")
	} else {
		fmt.Println("Ты не попал ни в какую категорию!")
	}
} */

/* func main() {
	subscribed := true
	
	if subscribed {
		fmt.Println("Вы подписаны на канал")
	} else {
		fmt.Println("Вы НЕ подписаны на канал")
	}
} */

func main() {
	score := 21

	if score < 6 || score > 16 { // Оператор ИЛИ || всё выражение делает истинным, если хотя бы одно условие истинно (если все усл. ложные, то false)
		fmt.Println("Ты попал взону")
	} else {
		fmt.Println("Ты НЕ попал в зону")
	}

	if score > 8 && score < 15 { // Оператор И && всё выражение делает истинным, только если все условия соответствуют истине (если одно условие false, то всё выражение false)
		fmt.Println("Ты попал в яблочко")
	} else {
		fmt.Println("Ты НЕ попал в яблочко")
	}
}

// также есть операторы <=, >=, !=, (!true == false, !false == true)