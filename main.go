package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(charName, charClass string) string {
	const baseDamage = 5
	var randMin, randMax int

	switch charClass {
	case "warrior":
		randMin, randMax = 3, 5
	case "mage":
		randMin, randMax = 5, 10
	case "healer":
		randMin, randMax = -3, -1
	default:
		return "неизвестный класс персонажа"
	}

	attack := baseDamage + randint(randMin, randMax)

	return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, attack)
}

func defence(charName, charClass string) string {
	const baseDefence = 10
	var randMin, randMax int

	switch charClass {
	case "warrior":
		randMin, randMax = 5, 10
	case "mage":
		randMin, randMax = -2, 2
	case "healer":
		randMin, randMax = 2, 5
	default:
		return "неизвестный класс персонажа"
	}

	defence := baseDefence + randint(randMin, randMax)
	return fmt.Sprintf("%s блокировал %d урона.", charName, defence)
}

func special(charName, charClass string) string {
	var skill string
	var count int

	switch charClass {
	case "warrior":
		skill, count = "Выносливость", 105
	case "mage":
		skill, count = "Атака", 45
	case "healer":
		skill, count = "Защита", 40
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s применил специальное умение `%s %d`", charName, skill, count)
}

// здесь обратите внимание на имена параметров
func startTraining(charName, charClass string) string {
	switch {
	case charClass == "warrior":
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	case charClass == "mage":
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	case charClass == "healer":
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		switch {
		case cmd == "attack":
			fmt.Println(attack(charName, charClass))
		case cmd == "defence":
			fmt.Println(defence(charName, charClass))
		case cmd == "special":
			fmt.Println(special(charName, charClass))
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiceCharClass() string {
	var approveChoice string
	var charClass string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)

		switch {
		case charClass == "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case charClass == "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case charClass == "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		default:
			fmt.Println("Неизвестный класс. Попробуй еще раз.")
			continue
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}

	return charClass
}

func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := choiceCharClass()

	fmt.Println(startTraining(charName, charClass))
}

func randint(min, max int) int {
	return rand.Intn(max-min+1) + min
}
