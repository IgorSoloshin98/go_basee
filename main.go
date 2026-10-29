// \n -перенос строки
// \t -табуляцию делает
// \" - поставить кавычки на слов
// %f -флоат
// %t - булевое
// %d%% - будет процент

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

//ввод от пользователя
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

//func getFullLength(str string) (int, int) {
//	bytes := len(str)
//	runes := utf8.RuneCountInString(str)
//	return bytes, runes
//}
//или
//func getFullLength(str string) (int, int) {
//	return len(str), utf8.RuneCountInString(str)
//}

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

//Система оценок
//Нашему проекту приходится работать с разными учебными заведениями и не всегда система
//оценок похожа. Необходимо создать программу, которая запрашивает у пользователя оценку в
//диапазоне от 0 до 100 включительно и выводит соответствующую буквенную оценку (A, B, C, D, F).
//Детали
//Необходимо создать функцию, которая будет запрашивать у пользователя ввод числовой
//оценки (целое число) в диапазоне от 0 до 100. Необходимо убедиться, что введенное значение
//находится в указанном диапазоне. Если значение вне диапазона или введено некорректно, программа
//должна возвращать ошибку из функции, а при ее обработке выводить сообщение об ошибке и
//завершать выполнение.
//В зависимости от введенной оценки, программа должна присвоить ей соответствующую
//буквенную оценку по следующей шкале:
//90–100: A
//80–89: B
//70–79: C
//60–69: D
//Ниже 60: F

//package main

//import (
//	"errors"
//	"fmt"
//	"math/rand/v2"
//)

//func main() {
//	score, err := getScore()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}

//	letterGrade := getLetterGrade(score)
//	fmt.Printf("Ваша буквенная оценка: %s\n", letterGrade)
//}

//func getLetterGrade(score int) string {
//	switch {
//	case score >= 90:
//		return "A"
//	case score >= 80:
//		return "B"
//	case score >= 70:
//		return "C"
//	case score >= 60:
//		return "D"
//	default:
//		return "F"
//	}
//}

//func getScore() (int, error) {
//	var score int
//	fmt.Printf("Введите оценку (0-100): ")
//	if _, err := fmt.Scan(&score); err != nil {
//		return 0, fmt.Errorf("Ошибка ввода: %w", err)
//	}

//	if score < 0 || score > 100 {
//		return 0, errors.New("Оценка должна быть в диапазоне от 0 и до 100 вкл")
//	}

//	return score, nil
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

package main

import "fmt"

func main() {
	i := 1
	for i < 11 {
		fmt.Println("Привет", i)
		i++
	}
	fmt.Println(i)

	for i := range 10 {
		fmt.Println(i + 1)
	}
}
