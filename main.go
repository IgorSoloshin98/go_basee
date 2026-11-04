//func main() {
//	var input int
//	fmt.Println("Введите время в часах (0-23): ")
//	_, err := fmt.Scan(&input)
//	if err != nil || input < 0 || input > 23 {
//		fmt.Println("Неверно задано время")
//		os.Exit(1)
//	}
//	var timeOfDay string
//	switch {
//	case input >= 6 && input < 12:
//		timeOfDay = "Утро"
//	case input >= 12 && input < 18:
//		timeOfDay = "День"
//	case input >= 18 && input < 23:
//		timeOfDay = "Вечер"
//	default:
//		timeOfDay = "Ночь"
//	}
//	fmt.Printf("Сейчас %02dч. - %s\n", input, timeOfDay)
//}

//ВВОД ОТ ПОЛЬЗОВАТЕЛЯ!!!!
//package main
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Printf("Введите строку: ")
//	scanner.Scan()
//	str := scanner.Text()
//	fmt.Println("Вы ввели:", str)
//}
/// мы пишем ааппаап апапап апапап
///выдаст ааппаап апапап апапап

//switch day {
//case 1:
//	fmt.Println("Понедельник")
//case 2:
//	fmt.Printlm("Вторник")
//case 3:
//	fmt.Println("Среда")
//case 4:
//	fmt.Println("Четверг")
//case 5:
//	fmt.Println("Пятница")
//case 6:
//	fmt.Println("Суббота")
//	fallthrough - означает что если кейс 6 выполнится то дальше
//	тоже провалимся
//case 7:
//	fmt.Println("Воскресенье")
//default:
//	fmt.Println("Неизвестный день")
//}
//или выраженние
//switch {
//case day < 1:
//	fmt.Println("Некорректный день")
//case day >= 1 && day <= 7:
//	fmt.Println("Рабочий день")
//case day >= 6 && day <= 7:
//	fmt.Println("Выходной день")
//default:
//	fmt.Println("нет такой недели")
//

//функции для написания пароля и проверки
//package main

//import (
//	"fmt"
//	"unicode/utf8"
//	"strings"
//)

//func main() {
//	isSecured := securePassword("rrvhjfgh")
//	fmt.Println(isSecured)
//}

//func securePassword(pass string) bool {
//	return utf8.RuneCountInString(pass) >= 6 &&
//		pass != strings.ToLower(pass) &&
//		pass != strings.ToUpper(pass)
//}
//или

//func securePassword(pass string) bool {
//	if utf8.RuneCountInString(pass) < 6 {
//		return false
//	}
//	if pass == strings.ToLower(pass) {
//		return false
//	}
//	if pass == strings.ToUpper(pass) {
//		return false
//	}
//	return true
//}

//defer

//func main() {
//	divide(10, 0)
//}

//func divide(n1, n2 int) {
//	defer fmt.Println("Конеч функкции divide")   //defer означает что все равно выполнится при любом условии
//	fmt.Println(n1 / n2)
//}

//func main() {
//	fmt.Println(Attack())

//	attackWithDamageBoost := DamageBoostDecorator(Attack)
//	attackWithCriticalHit := CriticalHitDecoration(attackWithDamageBoost)
//	attackWithSlowEffect := SlowEffectDecorator(attackWithCriticalHit)
//	fmt.Println(attackWithSlowEffect())
//}

//func Attack() string {
//	return "Атака выполнена!"
//}

//func DamageBoostDecorator(attackFunc func() string) func() string {
//	return func() string {
//		return "Вам улыбнулась удача, нанесение урона увеличено на 10%! " + attackFunc()
//	}
//}

//func CriticalHitDecoration(attackFunc func() string) func() string {
//	return func() string {
//		if rand.Float64() < 0.25 {
//			return "Критический адар! Урон удвоен! " + attackFunc()
//		}
//		return attackFunc()
//	}
//}

//func SlowEffectDecorator(attackFunc func() string) func() string {
//	return func() string {
//		return attackFunc() + " Цель замедлена на 2 хода!"
//	}

//package main

//import "fmt"

//func main() {
//	i := 1
//	for i < 11 {
//		fmt.Println("Привет", i)
//		i++
//	}
//	fmt.Println(i)

//	for i := range 10 {
//		fmt.Println(i + 1)
//	}
//
// /	str := "Ку-Ку!"
//
//		for _, r := range str {
//			fmt.Println(string(r))
//		}
//	}
//package main

//import (
//	"bufio"
//	"errors"
//	"fmt"
//	"os"
//	"unicode"
//)

//func main() {
//	input, err := GetInput()
//	if err != nil {
//		fmt.Printf("Ошибка при получении текста: %v\n", err)
//		os.Exit(1)
//	}

//	letters, gigits, spaces, punctuation := CountCharacters(input)
//	DisplayResults(letters, gigits, spaces, punctuation)
//}

//func GetInput() (string, error) {
//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Print("Введите текст: ")
//	if !scanner.Scan() {
//		fmt.Println()
//		if err := scanner.Err(); err != nil {
//			return "", fmt.Errorf("scan error: %w", err)
//		}
//		return "", errors.New("unable to read input")
//	}
//	return scanner.Text(), nil
//}

//func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
//	for _, char := range text {
//		switch {
//		case unicode.IsLetter(char):
//			letters++
//		case unicode.IsDigit(char):
//			digits++
//		case unicode.IsSpace(char):
//			spaces++
//		case unicode.IsPunct(char):
//			punctuation++
//		}
//	}
//	return
//}

//func DisplayResults(letters, digits, spaces, punctuation int) {
//	fmt.Printf("Количество букв: %d\n", letters)
//	fmt.Printf("Количество цифр: %d\n", digits)
//	fmt.Printf("Количество пробелов: %d\n", spaces)
//	fmt.Printf("Количество знаков препинания: %d\n", punctuation)
//}

// массивы
// var nums [5]int
// fmt.Println(nums)
//
//	0  1  2  3  4
//
// var nums = [5]int{2, 4, 6, 8, 0}
// fmt.Println(nums[1])     //4
// nums[1] = 400 //2 400 6 8 0
// fmt.Println(len(nums)) //5
// fruits := [...]string{"apple","banan", "cherry"}
//
// nums := [10]int{1: 10, 4: 20}
//package main

//import "fmt"

//func main() {

//	nums := [5]int{5, 23, -9, 1, 6}

//		for i := 0; i < len(nums); i++ {
//			fmt.Println(i, nums[i])
//		}
//		for _, value := range nums {
//			fmt.Println(value)
//		}
//	}
//package main

//import "fmt"

//func main() {
//	nums := [5]int{1, 2, 3, 4, 5}
//	modifyArray(&nums)
//	fmt.Println(nums)
//}

//	func modifyArray(arr *[5]int) {
//		arr[2] = 100
//		fmt.Println(arr)
//	}
//
// Секретный код практика
//package main

//import "fmt"

//func main() {
//	arr := [5]int{3, 8, 1, 8, 1}
//	fmt.Println("Секретный код:", generateSecretCode(arr))
//}

//func generateSecretCode(arr [5]int) string {
//	min, max := arr[0], arr[0]
//	middlePart := ""
//	for _, num := range arr {
//		if num < min {
//			min = num
//		}
//		if num > max {
//			max = num
//		}
//		if num%2 == 0 {
//			middlePart += fmt.Sprintf("Е%d", num)
//		} else {
//			middlePart += fmt.Sprintf("О%d", num)
//		}

//	}
//	return fmt.Sprintf("%d%s%d", min, middlePart, max)
//}
