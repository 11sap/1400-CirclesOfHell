package chapters

import (
	"fmt"
	"math/rand"
)

func Run10() {
	tasks := []func(){
		ch10ex01, ch10ex02, ch10ex03, ch10ex04, ch10ex05, ch10ex06, ch10ex07, ch10ex08,
		ch10ex09, ch10ex10, ch10ex11, ch10ex12, ch10ex12, ch10ex13, ch10ex14, ch10ex15,
		ch10ex16, ch10ex17, ch10ex18, ch10ex19, ch10ex20, ch10ex21, ch10ex22, ch10ex23,
	}

	for {
		fmt.Print("\nГлава 10. Выберите задачу (1-28): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 28 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 28! Для выхода введите 0.")
		}
	}
}

func ch10ex01() {
	fmt.Println("а) 8 случайных чисел от 0 до 1:")
	for i := 0; i < 8; i++ {
		fmt.Printf("%.3f ", rand.Float64())
	}
	fmt.Println()

	var k int
	fmt.Print("\nб) Введите k: ")
	fmt.Scan(&k)
	fmt.Printf("%d случайных чисел от 0 до 1:\n", k)
	for i := 0; i < k; i++ {
		fmt.Printf("%.3f ", rand.Float64())
	}
	fmt.Println()

	fmt.Println("\nв) 15 случайных чисел от 25 до 26:")
	for i := 0; i < 15; i++ {
		fmt.Printf("%.3f ", 25+rand.Float64())
	}
	fmt.Println()

	fmt.Println("\nг) 20 случайных чисел от 0 до 15:")
	for i := 0; i < 20; i++ {
		fmt.Printf("%.2f ", 15*rand.Float64())
	}
	fmt.Println()

	var a, b float64
	fmt.Print("\nд) Введите a и b: ")
	fmt.Scan(&a, &b)
	kRandom := rand.Intn(int(a)) + 1
	fmt.Printf("k = %d случайных чисел от 0 до %.2f:\n", kRandom, b)
	for i := 0; i < kRandom; i++ {
		fmt.Printf("%.2f ", b*rand.Float64())
	}
	fmt.Println()

	fmt.Println("\nе) 10 случайных чисел от -40 до 40:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%.2f ", 80*rand.Float64()-40)
	}
	fmt.Println()

	var m int
	fmt.Print("\nж) Введите m, a, b: ")
	fmt.Scan(&m, &a, &b)
	kRandom = rand.Intn(m) + 1
	fmt.Printf("k = %d случайных чисел от %.2f до %.2f:\n", kRandom, a, b)
	rangeSize := b - a
	for i := 0; i < kRandom; i++ {
		fmt.Printf("%.2f ", a+rangeSize*rand.Float64())
	}
	fmt.Println()
}

func ch10ex02() {
	fmt.Println("а) 10 случайных целых чисел от 0 до 10:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rand.Intn(11))
	}
	fmt.Println()

	var k, a int
	fmt.Print("\nб) Введите k и a: ")
	fmt.Scan(&k, &a)
	fmt.Printf("%d случайных целых чисел от 0 до %d:\n", k, a)
	for i := 0; i < k; i++ {
		fmt.Printf("%d ", rand.Intn(a+1))
	}
	fmt.Println()

	fmt.Println("\nв) 20 случайных целых чисел от 10 до 20:")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", 10+rand.Intn(11))
	}
	fmt.Println()

	fmt.Print("\nг) Введите k и a: ")
	fmt.Scan(&k, &a)
	fmt.Printf("%d случайных целых чисел от -10 до %d:\n", k, a)
	for i := 0; i < k; i++ {
		fmt.Printf("%d ", -10+rand.Intn(a+11))
	}
	fmt.Println()

	var b int
	fmt.Print("\nд) Введите a и b: ")
	fmt.Scan(&a, &b)
	kRandom := rand.Intn(15) + 1
	fmt.Printf("k = %d случайных целых чисел от %d до %d:\n", kRandom, a, b)
	for i := 0; i < kRandom; i++ {
		fmt.Printf("%d ", a+rand.Intn(b-a+1))
	}
	fmt.Println()
}

func ch10ex03() {
	var a, b int
	fmt.Print("Введите a и b: ")
	fmt.Scan(&a, &b)

	m := rand.Intn(20) + 1
	n := rand.Intn(20) + 1
	fmt.Printf("m = %d, n = %d\n", m, n)

	fmt.Printf("%d целых чисел от 0 до %d:\n", n, b)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", rand.Intn(b+1))
	}
	fmt.Println()

	fmt.Printf("%d неотрицательных действительных чисел (0..%.2f):\n", m, float64(m))
	for i := 0; i < m; i++ {
		fmt.Printf("%.2f ", float64(m)*rand.Float64())
	}
	fmt.Println()
}

func ch10ex04() {
	fmt.Println("50 четных чисел от 0 до 3 включительно (только единицы и двойки):")
	for i := 0; i < 50; i++ {
		num := 2 * rand.Intn(2)
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func ch10ex05() {
	fmt.Println("30 нечетных целых чисел от 0 до 5 включительно:")
	for i := 0; i < 30; i++ {
		num := 2*rand.Intn(3) + 1
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func ch10ex06() {
	fmt.Println("50 целых чисел (0 или 1):")
	zeros := 0
	ones := 0

	for i := 0; i < 50; i++ {
		num := rand.Intn(2)
		fmt.Printf("%d ", num)
		if num == 0 {
			zeros++
		} else {
			ones++
		}
	}

	fmt.Printf("\nКоличество нулей: %d\n", zeros)
	fmt.Printf("Количество единиц: %d\n", ones)
}

func ch10ex07() {
	fmt.Println("а) Два различных целых числа:")
	a := rand.Intn(3)
	b := rand.Intn(4)
	for b == a {
		b = rand.Intn(4)
	}
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Println("\nб) Три различных целых числа:")
	a = 1 + rand.Intn(3)
	b = rand.Intn(4)
	c := 1 + rand.Intn(4)

	for b == a {
		b = rand.Intn(4)
	}
	for c == a || c == b {
		c = 1 + rand.Intn(4)
	}
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	fmt.Println("\nв) 15 чисел (7 двоек и 8 троек):")
	numbers := make([]int, 15)

	for i := 0; i < 7; i++ {
		numbers[i] = 2
	}

	for i := 7; i < 15; i++ {
		numbers[i] = 3
	}

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", numbers[i])
	}
	fmt.Println()
}

func ch10ex08() {
	fmt.Println("10.8. Бросание монеты:")
	coin := rand.Intn(2)
	if coin == 0 {
		fmt.Println("Выпала решка (0)")
	} else {
		fmt.Println("Выпал орел (1)")
	}
}

func ch10ex09() {
	fmt.Println("10.9. Относительная частота выпадения монеты:")

	fmt.Println("При 100 бросках:")
	zeros := 0
	ones := 0
	for i := 0; i < 100; i++ {
		if rand.Intn(2) == 0 {
			zeros++
		} else {
			ones++
		}
	}
	fmt.Printf("0 (решка): %d (%.2f%%)\n", zeros, float64(zeros)/100*100)
	fmt.Printf("1 (орел):  %d (%.2f%%)\n", ones, float64(ones)/100*100)

	fmt.Println("\nПри 1000 бросках:")
	zeros = 0
	ones = 0
	for i := 0; i < 1000; i++ {
		if rand.Intn(2) == 0 {
			zeros++
		} else {
			ones++
		}
	}
	fmt.Printf("0 (решка): %d (%.2f%%)\n", zeros, float64(zeros)/1000*100)
	fmt.Printf("1 (орел):  %d (%.2f%%)\n", ones, float64(ones)/1000*100)
}

func ch10ex10() {
	fmt.Println("10.10. Игра 'Угадай четность'")

	fmt.Println("Выберите вариант игры:")
	fmt.Println("1 - Один раунд")
	fmt.Println("2 - Несколько раундов")
	fmt.Println("3 - Игра до отказа")

	var choice int
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:

		fmt.Println("\nа) Один раунд:")

		compNum := rand.Intn(100) + 1

		var answer int
		fmt.Print("Четное (введите 2) или нечетное (введите 1)? ")
		fmt.Scan(&answer)

		compParity := 1
		if compNum%2 == 0 {
			compParity = 2
		}

		fmt.Printf("Компьютер загадал число %d (", compNum)
		if compNum%2 == 0 {
			fmt.Print("четное)\n")
		} else {
			fmt.Print("нечетное)\n")
		}

		if answer == compParity {
			fmt.Println("Вы угадали!")
		} else {
			fmt.Println("Вы не угадали.")
		}

	case 2:
		fmt.Println("\nб) Несколько раундов:")

		var n int
		fmt.Print("Сколько раундов? ")
		fmt.Scan(&n)

		playerWins := 0
		compWins := 0

		for round := 0; round < n; round++ {
			compNum := rand.Intn(100) + 1
			compParity := 1
			if compNum%2 == 0 {
				compParity = 2
			}

			var answer int
			fmt.Printf("Раунд %d: Четное (2) или нечетное (1)? ", round+1)
			fmt.Scan(&answer)

			if answer == compParity {
				playerWins++
				fmt.Printf("Угадано! Число было %d\n", compNum)
			} else {
				compWins++
				fmt.Printf("Не угадано! Число было %d\n", compNum)
			}
		}

		fmt.Printf("\nСчет %d:%d ", playerWins, compWins)
		if playerWins > compWins {
			fmt.Println("в вашу пользу. Вы выиграли!")
		} else if playerWins < compWins {
			fmt.Println("в пользу компьютера. Вы проиграли!")
		} else {
			fmt.Println("Ничья!")
		}

	case 3:
		fmt.Println("\nв) Игра до отказа:")

		correct := 0
		incorrect := 0

		for {
			compNum := rand.Intn(100) + 1
			compParity := 1
			if compNum%2 == 0 {
				compParity = 2
			}

			var answer int
			fmt.Print("Четное (2) или нечетное (1)? (0 - выход) ")
			fmt.Scan(&answer)

			if answer == 0 {
				break
			}

			if answer == compParity {
				correct++
				fmt.Printf("Верно! Число было %d\n", compNum)
			} else {
				incorrect++
				fmt.Printf("Неверно! Число было %d\n", compNum)
			}
		}

		fmt.Printf("\nИтог: верных ответов - %d, неверных - %d\n", correct, incorrect)

	default:
		fmt.Println("Неверный выбор")
	}
}

func ch10ex11() {
	fmt.Println("10.11. Бросание игрального кубика:")
	dice := rand.Intn(6) + 1
	fmt.Printf("Выпало: %d\n", dice)
}

func ch10ex12() {
	fmt.Println("10.12. Два игрока бросают кубик:")

	player1 := rand.Intn(6) + 1
	player2 := rand.Intn(6) + 1

	fmt.Printf("Игрок 1: %d\n", player1)
	fmt.Printf("Игрок 2: %d\n", player2)

	if player1 > player2 {
		fmt.Println("Выиграл игрок 1!")
	} else if player2 > player1 {
		fmt.Println("Выиграл игрок 2!")
	} else {
		fmt.Println("Ничья!")
	}
}

func ch10ex13() {
	fmt.Println("10.13. Игра в кубик")

	fmt.Println("Выберите вариант игры:")
	fmt.Println("1 - Каждый бросает кубик два раза")
	fmt.Println("2 - Несколько раундов (каждый бросает по одному разу)")

	var choice int
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("\nВариант 1: каждый бросает кубик два раза")

		player1Score := 0
		fmt.Println("Игрок 1 бросает:")
		for i := 0; i < 2; i++ {
			roll := rand.Intn(6) + 1
			fmt.Printf("  Бросок %d: %d\n", i+1, roll)
			player1Score += roll
		}
		fmt.Printf("Сумма игрока 1: %d\n\n", player1Score)

		player2Score := 0
		fmt.Println("Игрок 2 бросает:")
		for i := 0; i < 2; i++ {
			roll := rand.Intn(6) + 1
			fmt.Printf("  Бросок %d: %d\n", i+1, roll)
			player2Score += roll
		}
		fmt.Printf("Сумма игрока 2: %d\n\n", player2Score)

		if player1Score > player2Score {
			fmt.Println("Выиграл игрок 1!")
		} else if player2Score > player1Score {
			fmt.Println("Выиграл игрок 2!")
		} else {
			fmt.Println("Ничья!")
		}

	case 2:
		fmt.Println("\nВариант 2: несколько раундов")

		var rounds int
		fmt.Print("Сколько раундов? ")
		fmt.Scan(&rounds)

		player1Wins := 0
		player2Wins := 0
		draws := 0

		for round := 1; round <= rounds; round++ {
			player1 := rand.Intn(6) + 1
			player2 := rand.Intn(6) + 1

			if player1 > player2 {
				player1Wins++
			} else if player2 > player1 {
				player2Wins++
			} else {
				draws++
			}
		}

		fmt.Println("\nРезультаты:")
		fmt.Printf("Игрок 1 выиграл: %d раз\n", player1Wins)
		fmt.Printf("Игрок 2 выиграл: %d раз\n", player2Wins)
		fmt.Printf("Ничья: %d раз\n", draws)

		if player1Wins > player2Wins {
			fmt.Println("\nОбщий победитель: Игрок 1!")
		} else if player2Wins > player1Wins {
			fmt.Println("\nОбщий победитель: Игрок 2!")
		} else {
			fmt.Println("\nОбщая ничья!")
		}

	default:
		fmt.Println("Неверный выбор")
	}
}

func ch10ex14() {
	fmt.Println("10.14. Три игрока бросают кубики:")

	var k int
	fmt.Print("Сколько кубиков бросает каждый игрок? ")
	fmt.Scan(&k)

	scores := [3]int{}

	for player := 0; player < 3; player++ {
		fmt.Printf("\nИгрок %d:\n", player+1)
		for i := 0; i < k; i++ {
			roll := rand.Intn(6) + 1
			fmt.Printf("  Кубик %d: %d\n", i+1, roll)
			scores[player] += roll
		}
		fmt.Printf("Сумма: %d\n", scores[player])
	}

	maxScore := scores[0]
	winners := []int{0}

	for i := 1; i < 3; i++ {
		if scores[i] > maxScore {
			maxScore = scores[i]
			winners = []int{i}
		} else if scores[i] == maxScore {
			winners = append(winners, i)
		}
	}

	fmt.Println("\nРезультаты:")
	if len(winners) == 1 {
		fmt.Printf("Победил игрок %d со счетом %d!\n", winners[0]+1, maxScore)
	} else {
		fmt.Printf("Ничья между игроками ")
		for i, w := range winners {
			if i > 0 {
				fmt.Print(" и ")
			}
			fmt.Printf("%d", w+1)
		}
		fmt.Printf(" со счетом %d!\n", maxScore)
	}
}

func ch10ex15() {
	fmt.Println("10.15. Частота выпадения чисел на кубике:")

	fmt.Println("100 бросков:")
	counts := make([]int, 7)
	for i := 0; i < 100; i++ {
		roll := rand.Intn(6) + 1
		counts[roll]++
	}

	for i := 1; i <= 6; i++ {
		fmt.Printf("%d: %d (%.1f%%)\n", i, counts[i], float64(counts[i])/100*100)
	}

	fmt.Println("\n1000 бросков:")
	counts = make([]int, 7)
	for i := 0; i < 1000; i++ {
		roll := rand.Intn(6) + 1
		counts[roll]++
	}

	for i := 1; i <= 6; i++ {
		fmt.Printf("%d: %d (%.1f%%)\n", i, counts[i], float64(counts[i])/1000*100)
	}
}

func ch10ex16() {
	fmt.Println("10.16. Выбор кости домино:")

	side1 := rand.Intn(7)
	side2 := rand.Intn(7)

	fmt.Printf("Выбрана кость %d-%d\n", side1, side2)
}

func ch10ex17() {
	fmt.Println("10.17. Выбор двух костей домино:")

	side1a := rand.Intn(7)
	side1b := rand.Intn(7)

	side2a := rand.Intn(7)
	side2b := rand.Intn(7)

	fmt.Printf("Кость 1: %d-%d\n", side1a, side1b)
	fmt.Printf("Кость 2: %d-%d\n", side2a, side2b)

	canConnect := false

	if side1a == side2a || side1a == side2b ||
		side1b == side2a || side1b == side2b {
		canConnect = true
	}

	if canConnect {
		fmt.Println("Кости можно приставить одна к другой")
	} else {
		fmt.Println("Кости нельзя приставить")
	}
}

func ch10ex18() {
	fmt.Println("10.18. Проверка таблицы умножения")

	fmt.Println("а) Один вопрос:")
	a := rand.Intn(9) + 1
	b := rand.Intn(9) + 1

	var answer int
	fmt.Printf("Сколько будет %d × %d? ", a, b)
	fmt.Scan(&answer)

	if answer == a*b {
		fmt.Println("Правильно!")
	} else {
		fmt.Printf("Неправильно! Правильный ответ: %d\n", a*b)
	}
}

func ch10ex18b() {
	fmt.Println("\nб) 20 вопросов:")

	correct := 0
	incorrect := 0

	for i := 0; i < 20; i++ {
		a := rand.Intn(9) + 1
		b := rand.Intn(9) + 1

		var answer int
		fmt.Printf("Вопрос %d: %d × %d = ", i+1, a, b)
		fmt.Scan(&answer)

		if answer == a*b {
			correct++
		} else {
			incorrect++
			fmt.Printf("Неправильно! Правильный ответ: %d\n", a*b)
		}
	}

	fmt.Printf("\nРезультат: правильных - %d, неправильных - %d\n", correct, incorrect)
}

func ch10ex18c() {
	fmt.Println("\nв) Вопросы до ответа 0:")

	questionNum := 1
	for {
		a := rand.Intn(9) + 1
		b := rand.Intn(9) + 1

		var answer int
		fmt.Printf("Вопрос %d: %d × %d = (0 - выход) ", questionNum, a, b)
		fmt.Scan(&answer)

		if answer == 0 {
			break
		}

		if answer == a*b {
			fmt.Println("Правильно!")
		} else {
			fmt.Printf("Неправильно! Правильный ответ: %d\n", a*b)
		}

		questionNum++
	}
}

func ch10ex19() {
	fmt.Println("10.19. Выбор карты одной масти:")

	ranks := []string{"6", "7", "8", "9", "10", "валет", "дама", "король", "туз"}
	rank := ranks[rand.Intn(len(ranks))]

	fmt.Printf("Выбрана карта: %s\n", rank)
}

func ch10ex20() {
	fmt.Println("10.20. Выбор карты из полной колоды:")

	suits := []string{"пик", "треф", "бубен", "червей"}
	ranks := []string{"шестерка", "семерка", "восьмерка", "девятка",
		"десятка", "валет", "дама", "король", "туз"}

	suit := suits[rand.Intn(len(suits))]
	rank := ranks[rand.Intn(len(ranks))]

	var suffix string
	switch suit {
	case "бубен", "червей":
		suffix = ""
	default:
		suffix = "а"
	}

	fmt.Printf("Выбрана %s %s%s\n", rank, suit, suffix)
}

func ch10ex21() {
	fmt.Println("10.21. Сравнение двух карт:")

	suits := []string{"пик", "треф", "бубен", "червей"}
	ranks := []string{"шестерка", "семерка", "восьмерка", "девятка",
		"десятка", "валет", "дама", "король", "туз"}

	suit1 := suits[rand.Intn(len(suits))]
	rank1 := ranks[rand.Intn(len(ranks))]

	suit2 := suits[rand.Intn(len(suits))]
	rank2 := ranks[rand.Intn(len(ranks))]
	for suit2 == suit1 && rank2 == rank1 {
		suit2 = suits[rand.Intn(len(suits))]
		rank2 = ranks[rand.Intn(len(ranks))]
	}

	var suffix1, suffix2 string
	if suit1 == "бубен" || suit1 == "червей" {
		suffix1 = ""
	} else {
		suffix1 = "а"
	}
	if suit2 == "бубен" || suit2 == "червей" {
		suffix2 = ""
	} else {
		suffix2 = "а"
	}

	fmt.Printf("Игрок 1: %s %s%s\n", rank1, suit1, suffix1)
	fmt.Printf("Игрок 2: %s %s%s\n", rank2, suit2, suffix2)

	suitValue := map[string]int{
		"пик":    1,
		"треф":   2,
		"бубен":  3,
		"червей": 4,
	}

	rankValue := make(map[string]int)
	for i, r := range ranks {
		rankValue[r] = i + 1
	}

	player1Wins := false
	player2Wins := false

	if suitValue[suit1] > suitValue[suit2] {
		player1Wins = true
	} else if suitValue[suit2] > suitValue[suit1] {
		player2Wins = true
	} else {

		if rankValue[rank1] > rankValue[rank2] {
			player1Wins = true
		} else if rankValue[rank2] > rankValue[rank1] {
			player2Wins = true
		}
	}

	fmt.Print("\nРезультат: ")
	if player1Wins {
		fmt.Println("Выиграл игрок 1!")
	} else if player2Wins {
		fmt.Println("Выиграл игрок 2!")
	} else {
		fmt.Println("Ничья! (маловероятно, но возможно)")
	}
}

func ch10ex22() {
	fmt.Println("10.22. Неоднократный выбор карт:")

	suits := []string{"пик", "треф", "бубен", "червей"}
	ranks := []string{"шестерка", "семерка", "восьмерка", "девятка",
		"десятка", "валет", "дама", "король", "туз"}

	var rounds int
	fmt.Print("Сколько раундов? ")
	fmt.Scan(&rounds)

	player1Wins := 0
	player2Wins := 0

	for round := 1; round <= rounds; round++ {
		fmt.Printf("\nРаунд %d:\n", round)

		suit1 := suits[rand.Intn(len(suits))]
		rank1 := ranks[rand.Intn(len(ranks))]

		suit2 := suits[rand.Intn(len(suits))]
		rank2 := ranks[rand.Intn(len(ranks))]

		suitValue := map[string]int{"пик": 1, "треф": 2, "бубен": 3, "червей": 4}
		rankValue := make(map[string]int)
		for i, r := range ranks {
			rankValue[r] = i + 1
		}

		if suitValue[suit1] > suitValue[suit2] ||
			(suitValue[suit1] == suitValue[suit2] && rankValue[rank1] > rankValue[rank2]) {
			player1Wins++
			fmt.Println("  Выиграл игрок 1")
		} else if suitValue[suit2] > suitValue[suit1] ||
			(suitValue[suit1] == suitValue[suit2] && rankValue[rank2] > rankValue[rank1]) {
			player2Wins++
			fmt.Println("  Выиграл игрок 2")
		} else {
			fmt.Println("  Ничья")
		}
	}

	fmt.Printf("\nИтог: игрок 1 - %d, игрок 2 - %d\n", player1Wins, player2Wins)
	if player1Wins > player2Wins {
		fmt.Println("Общий победитель: игрок 1!")
	} else if player2Wins > player1Wins {
		fmt.Println("Общий победитель: игрок 2!")
	} else {
		fmt.Println("Общая ничья!")
	}
}

func ch10ex23() {
	fmt.Println("10.23. Козырная масть:")

	suits := []string{"пик", "треф", "бубен", "червей"}
	ranks := []string{"шестерка", "семерка", "восьмерка", "девятка",
		"десятка", "валет", "дама", "король", "туз"}

	trumpSuit := suits[rand.Intn(len(suits))]
	fmt.Printf("Козырная масть: %s\n\n", trumpSuit)

	suit := suits[rand.Intn(len(suits))]
	rank := ranks[rand.Intn(len(ranks))]

	var suffix string
	if suit == "бубен" || suit == "червей" {
		suffix = ""
	} else {
		suffix = "а"
	}

	fmt.Printf("Выбрана карта: %s %s%s\n", rank, suit, suffix)

	if suit == trumpSuit {
		fmt.Println("Это козырная карта!")
	} else {
		fmt.Println("Это некозырная карта.")
	}
}
