package main

import "fmt"

/* 
	Структуры - это логическое объединение нескольких переменных для действий с ними
*/

//=============================================================================================//

type User struct {
	Name        string 	// ""
	Age 		int		// 0
	PhoneNumber string	// ""
	IsClose 	bool 	// false
	Rating 		float64	// 0.0
}

func main() {
	user := User{}
	
	fmt.Println("User:", user) // User: { 0  false 0}
}

//=============================================================================================//

/* type User struct {
	Name        string 	
	Age 		int	
	PhoneNumber string	
	IsClose 	bool 	
	Rating 		float64	
}

func main() {
	user := User{
		Name: "Сергей",
		// Age: 64, // При отсутсвии одного параметра будет значение по умолчанию
		PhoneNumber: "+7 (911) 911-11-22",
		IsClose: true,
		Rating: 5.5,
	}
	
	fmt.Println("User:", user)
	fmt.Println("Имя:", user.Name) // Досуп к отдельному параметру
	// можно работать с полями структур
	fmt.Println("Имя до изменения:", user.Name)
	user.Name = "Игорь"
	fmt.Println("Имя после изменения:", user.Name)
} */

//=============================================================================================//

// Методы

/* type User struct {
	Name 		string		
	Rating 		float64
}

// Метод - функция, которая относится к конкретной структуре (синтаксический сахар для обычной функции, он даёт возможность узнать,
// что она относится к конкретной структуре)
func (u User) Greeting() { 
	fmt.Println("Всем привет!")
	fmt.Println("Меня зовут", u.Name)
	fmt.Println("Мой рейтинг:", u.Rating)
	fmt.Println("")
}

func GoodBye(u User) {
	fmt.Println("Всем пока!")
	fmt.Println("Меня звали", u.Name)
	fmt.Println("Мой рейтинг был:", u.Rating)
	fmt.Println("")
}

func (u User) RaitingUP(rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}

func main() {
	user := User{
		Name: "Вася",
		Rating: 6.0,
	}
	user.Greeting()
	GoodBye(user)

	user.RaitingUP(4.0)
} */

//=============================================================================================//

/* type User struct {
	Name 		string		
	Rating 		float64
} 

// Если при передаче переменной в качестве арг. функции никаких указателей не используется, то эта переменная копируется.
//  
func RaitingUP(u User, rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}

func main() {
	user := User{
		Name: "Вася",
		Rating: 6.0,
	}
	
	fmt.Println("Рейтинг до:", user.Rating) // Rating в исходном User не изменился.
	RaitingUP(user, 5.0)	 // Мы изменили только его коппию
	fmt.Println("Рейтинг после:", user.Rating) 
}  */

//=============================================================================================//

//  Исправление с помощью указателей

/* type User struct {
	Name 		string		
	Rating 		float64
}

func RatingUP(u *User, rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}

func main() {
	user := User{
		Name: "Вася",
		Rating: 6.0,
	}
	
	ptr := &user

	fmt.Println("User до:", user)
	fmt.Println("")

	RatingUP(ptr, 3.0)	 
	fmt.Println("User после:", user)
} */

//=============================================================================================//

/* type User struct {
	Name 		string		
	Rating 		float64
}

// Аргумент, содержащий сруктуру называется РЕСИВЕР(может быть по значению, когда созда1тся коппия, 
// и по указателю, когда передаётся указатель)
// С коппией мы не сможем менять исходную структуру(можно применять, если нам не требуется менять структуру)
func (u *User) RaitingUP(rating float64) { 
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошёл валидацию")
	}
}

func main() {
	user := User{
		Name: "Вася",
		Rating: 6.0,
	}
	
	fmt.Println("User до:", user)
	fmt.Println("")

	user.RaitingUP(1.0)	 
	fmt.Println("User после:", user)
} */

//=============================================================================================//

/* type User struct {
	Name			string
	Age				int
	PhoneNumber		string
	IsClose			bool 
	Rating			float64
}

// конструктор
func NewUser(
	name string, 
	age int, 
	phoneNumber string, 
	isClose bool, 
	rating float64,
) User {
	fmt.Println("Валидирую имя")
	if name == "" {
		fmt.Println("Имя не прошло валидацию :(")
		return User{}
	}
	fmt.Println("Валидирую возраст")
	if age <=0 || age >= 150 {
		fmt.Println("Возраст не прошёл валидацию :(")
		return User{}
	}
	fmt.Println("Валидирую номер телефона")
	if phoneNumber == "" {
		fmt.Println("Номер телефона не прошёл валидацию :(")
		return User{}
	}
	fmt.Println("Валидирую рейтинг")
	if rating < 0.0 || rating > 10.0 {
		fmt.Println("Рейтинг не прошёл валидацию :(")
		return User{}
	}
	return User{
		Name: name,
		Age: age,
		PhoneNumber: phoneNumber,
		IsClose: isClose,
		Rating: rating,
	}
} 

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	} 
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	} 
}

func (u *User) CloseAccount() {
	// Закрыть аккаунт
	u.IsClose = true
}

func (u *User) OpenAccount() {
	// Открыть аккаунт
	u.IsClose = false
}

func (u *User) RatingUP(rating float64) {
	if u.Rating + rating <= 10.0 {
		u.Rating += rating
	}
}

func (u *User) RatingDOWN(rating float64) {
	if u.Rating + rating >= 0.0 {
		u.Rating -= rating
	}
}

func main() {
	user := NewUser(
		"Даниил",
		50,
		"+79114455555",
		false,
		4.5,
	)
	

	user.ChangeAge(4)

	fmt.Println("User Age:", user.GetAge())
	fmt.Println("User:", user)
} */

/* 
	Если поменять значения в структуре напрямую или изначально написать неверные значения,
	то методы работать не будут!

	Неверные значения при создании можно пофиксить с помощью конструктора
*/