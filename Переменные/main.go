package main

import "fmt"

/* func main() {
	text := "Game Over"

	// : пишется только когда создаём переменную и оно определяет тип данных для этой переменной(присвоить строку переменной с изначальным числом нельзя)
	
	score := 111
	score = 50 + 10
	score = score + 10

	fmt.Println(text)
	fmt.Println(score)
} */

/* func main() {
	text1 := "Get Ready"
	text2 := "Game Over"
	score := 0

	fmt.Println(text1, text2, score) // Get Ready Game Over 0
	fmt.Println("Ваш счёт:", score)
} */

/* func main() {
	text := "Hello"
	// Дробные и целые числа - это разные типы данных
	// Целое число - int
	// Дробное число - float
	
	score := 1
	drob := 0.5
	subscribed := true

	fmt.Println(text)
	fmt.Println(score)
	fmt.Println(drob)
	fmt.Println("Подписан?", subscribed)
} */

/* func main() {
	score := 10.0

	score = score + 10
	score = score - 5
	score = score * 2
	// при делении любого целого числа на целое, результат будет целым.
	// чтобы получить дробный результат, то и делимое число нужно привести к дроби
	score = score / 4 
	score += 10 // сложение с присваиванием(то же самое и с другими операциями)
	score++ // инкремент(увеличевает на 1)
	score-- // дикремент(уменьшает на 1)


	fmt.Println(score)
} */

/* func main() {
	score := 11
	celoe := score / 3
	ostatok := score % 3

	fmt.Println("Счёт:", score, celoe, ostatok)
} */

/* func main() {
	text := "hello"
	// text = text + "world" 

	fmt.Println(text + " world")
} */

/* func main() {
	// Удлинённый способ записи переменных. В отличие от короткого способа в них можно ничего не присваивать сразу.
	var number int 
	var text string 

	fmt.Println("number:", number)
	fmt.Println("text:", text)

	number = 15
	text = "Hello, World!"

	fmt.Println("number:", number)
	fmt.Println("text:", text)
} */

func main() {
	var number int // знач. по умолчанию 0
	var text string // знач. по умолчанию ""
	var drob float64 // знач. по умолчанию 0.0 (цифра означает кол-во занимаемой оперативной памяти)
	var boolean bool // знач. по умолчанию false

	fmt.Println("number:", number)
	fmt.Println("text:", text)
	fmt.Println("drob:", drob)
	fmt.Println("boolean:", boolean)
}