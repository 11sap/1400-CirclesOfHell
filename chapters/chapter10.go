package chapters

import (
	"fmt"
	"math/rand"
)

func Run10() {
	tasks := []func(){
		ch10ex01, ch10ex02, ch10ex03, ch10ex04, ch10ex05, ch10ex06, ch10ex07,
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
