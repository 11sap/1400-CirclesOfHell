package chapters

import (
	"fmt"
	"math"
	"math/rand"
)

func Run11() {
	tasks := []func(){
		ch11ex01, ch11ex02, ch11ex03, ch11ex04, ch11ex05, ch11ex06, ch11ex07, ch11ex08,
		ch11ex09, ch11ex10, ch11ex11, ch11ex12, ch11ex13, ch11ex14, ch11ex15, ch11ex16,
		ch11ex17, ch11ex18, ch11ex19, ch11ex20, ch11ex21, ch11ex22, ch11ex23, ch11ex24,
		ch11ex25, ch11ex26, ch11ex27, ch11ex28, ch11ex29, ch11ex30, ch11ex31, ch11ex32,
		ch11ex33, ch11ex34, ch11ex35, ch11ex36, ch11ex37, ch11ex38, ch11ex39, ch11ex40,
		ch11ex41, ch11ex42, ch11ex43, ch11ex44, ch11ex45, ch11ex46, ch11ex47, ch11ex48,
		ch11ex49, ch11ex50, ch11ex51, ch11ex52,
	}

	for {
		fmt.Print("\nГлава 11. Выберите задачу (1-261): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 261 {
			tasks[n-1]()
		} else {
			fmt.Println("От 1 до 261! Для выхода введите 0.")
		}
	}
}

func printArray(arr []int, name string) {
	fmt.Printf("%s: [", name)
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Println("]")
}

func printFloatArray(arr []float64, name string) {
	fmt.Printf("%s: [", name)
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", v)
	}
	fmt.Println("]")
}

func ch11ex01() {
	fmt.Println("11.1. Заполнение массива из 8 элементов")

	arr := [8]int{37, 0, 50, 46, 34, 46, 0, 13}

	fmt.Print("Массив: [")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
		if i < len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

func ch11ex02() {
	fmt.Println("11.2. Массив чисел от 1 до 10")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}

	printArray(arr[:], "Массив")
}

func ch11ex03() {
	fmt.Println("11.3. Массив чисел от 10 до 1")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = 10 - i
	}

	printArray(arr[:], "Массив")
}

func ch11ex04() {
	fmt.Println("11.4. Первые 10 четных чисел")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = 2 * (i + 1)
	}

	printArray(arr[:], "Массив")
}

func ch11ex05() {
	fmt.Println("11.5. Первые 10 нечетных чисел")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = 2*i + 1
	}

	printArray(arr[:], "Массив")
}

func ch11ex06() {
	fmt.Println("11.6. Первые 10 чисел, делящихся на 3")

	var arr [10]int
	count := 0
	num := 1

	for count < 10 {
		if num%3 == 0 {
			arr[count] = num
			count++
		}
		num++
	}

	printArray(arr[:], "Массив")
}

func ch11ex07() {
	fmt.Println("11.7. Первые 10 степеней числа 2")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = int(math.Pow(2, float64(i)))
	}

	printArray(arr[:], "Массив")
}

func ch11ex08() {
	fmt.Println("11.8. Первые 15 чисел Фибоначчи")

	var arr [15]int
	arr[0] = 1
	arr[1] = 1

	for i := 2; i < 15; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	printArray(arr[:], "Массив")
}

func ch11ex09() {
	fmt.Println("11.9. 10 членов арифметической прогрессии")

	var arr [10]int
	a := 3
	d := 2

	for i := 0; i < 10; i++ {
		arr[i] = a + i*d
	}

	printArray(arr[:], "Массив")
}

func ch11ex10() {
	fmt.Println("11.10. 10 членов геометрической прогрессии")

	var arr [10]int
	a := 3
	q := 2

	arr[0] = a
	for i := 1; i < 10; i++ {
		arr[i] = arr[i-1] * q
	}

	printArray(arr[:], "Массив")
}

func ch11ex11() {
	fmt.Println("11.11. Первые 10 точных квадратов")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = (i + 1) * (i + 1)
	}

	printArray(arr[:], "Массив")
}

func ch11ex12() {
	fmt.Println("11.12. Первые 10 точных кубов")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = (i + 1) * (i + 1) * (i + 1)
	}

	printArray(arr[:], "Массив")
}

func ch11ex13() {
	fmt.Println("11.13. Степени числа 2")

	n := 10
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = int(math.Pow(2, float64(i+1)))
	}

	printArray(arr, "Массив")
}

func ch11ex14() {
	fmt.Println("11.14. Цифры числа в обратном порядке")

	n := 123456
	var arr [6]int

	temp := n
	index := 0

	for temp > 0 {
		arr[index] = temp % 10
		temp /= 10
		index++
	}

	fmt.Printf("Число: %d\n", n)
	fmt.Print("Цифры в обратном порядке: [")
	for i := 0; i < index; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", arr[i])
	}
	fmt.Println("]")
}

func ch11ex15() {
	fmt.Println("11.15. Убывающая и возрастающая последовательности")

	arr1 := make([]int, 8)
	for i := 0; i < 8; i++ {
		arr1[i] = 8 - i
	}
	printArray(arr1, "а) Убывающая")

	arr2 := make([]int, 8)
	for i := 0; i < 8; i++ {
		arr2[i] = i + 1
	}
	printArray(arr2, "б) Возрастающая")
}

func ch11ex16() {
	fmt.Println("11.16. Арифметическая и геометрическая прогрессии")

	a1 := 2.0
	d := 3.0
	arr1 := make([]float64, 10)

	for i := 0; i < 10; i++ {
		arr1[i] = a1 + float64(i)*d
	}
	printFloatArray(arr1, "а) Арифметическая прогрессия")

	a2 := 2.0
	q := 1.5
	arr2 := make([]float64, 10)
	arr2[0] = a2

	for i := 1; i < 10; i++ {
		arr2[i] = arr2[i-1] * q
	}
	printFloatArray(arr2, "б) Геометрическая прогрессия")
}

func ch11ex17() {
	fmt.Println("11.17. Первые 10 чисел Фибоначчи")

	arr := make([]int, 10)
	arr[0] = 1
	if len(arr) > 1 {
		arr[1] = 1
	}

	for i := 2; i < 10; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	printArray(arr, "Массив Фибоначчи")
}

func ch11ex18() {
	fmt.Println("11.18. Числа, делящиеся на 13 и 17")

	arr1 := make([]int, 20)
	count := 0
	num := 300

	for count < 20 {
		if num%13 == 0 && num%17 == 0 {
			arr1[count] = num
			count++
		}
		num++
	}

	printArray(arr1, "а) Делятся на 13 и 17")

	arr2 := make([]int, 15)
	count = 0
	num = 2

	for count < 15 {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			arr2[count] = num
			count++
		}
		num++
	}

	printArray(arr2, "б) Первые 15 простых чисел")
}

func ch11ex19() {
	fmt.Println("11.19. Первые 10 простых чисел, начиная с 100")

	arr := make([]int, 10)
	count := 0
	num := 100

	for count < 10 {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			arr[count] = num
			count++
		}
		num++
	}

	printArray(arr, "Простые числа")
}

func ch11ex20() {
	fmt.Println("11.20. Проверка таблицы умножения")

	questions := make([][2]int, 20)
	answers := make([]int, 20)

	for i := 0; i < 20; i++ {
		questions[i][0] = 2 + rand.Intn(8)
		questions[i][1] = 2 + rand.Intn(8)
	}

	correct := 0
	for i := 0; i < 20; i++ {
		a, b := questions[i][0], questions[i][1]
		fmt.Printf("Вопрос %d: Сколько будет %d × %d? ", i+1, a, b)
		fmt.Scan(&answers[i])

		if answers[i] == a*b {
			correct++
			fmt.Println("Правильно!")
		} else {
			fmt.Printf("Неправильно! Правильный ответ: %d\n", a*b)
		}
	}

	fmt.Printf("\nРезультат: %d правильных из 20\n", correct)
}

func ch11ex21() {
	fmt.Println("11.21. Массив из 20 неповторяющихся элементов")

	arr := make([]int, 20)
	used := make(map[int]bool)

	for i := 0; i < 20; i++ {
		for {
			num := rand.Intn(100) + 1
			if !used[num] {
				arr[i] = num
				used[num] = true
				break
			}
		}
	}

	printArray(arr, "Массив")
}

func ch11ex22() {
	fmt.Println("11.22. Вывод элементов массива")

	arr := []int{-5, 10, 0, 150, -3, 75, 99, 200, -1, 50}
	printArray(arr, "Исходный массив")

	fmt.Print("а) Неотрицательные элементы: ")
	for _, v := range arr {
		if v >= 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	fmt.Print("б) Элементы не превышающие 100: ")
	for _, v := range arr {
		if v <= 100 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func ch11ex23() {
	fmt.Println("11.23. Четные элементы и элементы, оканчивающиеся нулем")

	arr := []int{10, 25, 30, 47, 50, 63, 70, 88, 90, 105}
	printArray(arr, "Исходный массив")

	fmt.Print("а) Четные элементы: ")
	for _, v := range arr {
		if v%2 == 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	fmt.Print("б) Элементы, оканчивающиеся нулем: ")
	for _, v := range arr {
		if v%10 == 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func ch11ex24() {
	fmt.Println("11.24. Двузначные и трехзначные элементы")

	arr := []int{5, 23, 456, 78, 9, 123, 45, 678, 90, 1000}
	printArray(arr, "Исходный массив")

	fmt.Print("а) Двузначные элементы: ")
	for _, v := range arr {
		if v >= 10 && v <= 99 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	fmt.Print("б) Трехзначные элементы: ")
	for _, v := range arr {
		if v >= 100 && v <= 999 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func ch11ex25() {
	fmt.Println("11.25. Элементы с четными и нечетными индексами")

	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	printArray(arr, "Исходный массив")

	fmt.Print("а) Элементы с четными индексами: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("б) Элементы с нечетными индексами: ")
	for i := 2; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func ch11ex26() {
	fmt.Println("11.26. Сначала неотрицательные, затем отрицательные элементы")

	arr := []int{-5, 10, -3, 0, 7, -1, 15, -8, 2, -4}
	printArray(arr, "Исходный массив")

	fmt.Print("Неотрицательные элементы: ")
	for _, v := range arr {
		if v >= 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	fmt.Print("Отрицательные элементы: ")
	for _, v := range arr {
		if v < 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func ch11ex27() {
	fmt.Println("11.27. Сначала четные, затем нечетные элементы")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printArray(arr, "Исходный массив")

	fmt.Print("Четные элементы: ")
	for _, v := range arr {
		if v%2 == 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()

	fmt.Print("Нечетные элементы: ")
	for _, v := range arr {
		if v%2 != 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func ch11ex28() {
	fmt.Println("11.28. Индексы элементов, оканчивающихся цифрой 0")

	arr := []int{10, 25, 30, 47, 50, 63, 70, 88, 90, 105}
	printArray(arr, "Исходный массив")

	fmt.Print("Индексы элементов, оканчивающихся на 0: ")
	for i, v := range arr {
		if v%10 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func ch11ex29() {
	fmt.Println("11.29. Дни января без осадков")

	precipitation := []int{5, 0, 3, 0, 0, 7, 0, 2, 0, 4,
		0, 6, 0, 0, 8, 0, 1, 0, 3, 0,
		0, 5, 0, 2, 0, 0, 9, 0, 4, 0, 0}

	fmt.Print("Дни без осадков: ")
	for i, p := range precipitation {
		if p == 0 {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}

func ch11ex30() {
	fmt.Println("11.30. Команды с менее чем 3 победами")

	wins := []int{5, 2, 7, 1, 3, 0, 4, 2, 6, 1,
		8, 3, 2, 0, 5, 1, 4, 2, 3, 0}

	fmt.Print("Номера команд с менее чем 3 победами: ")
	for i, w := range wins {
		if w < 3 {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}

func ch11ex31() {
	fmt.Println("11.31. Элементы с четных и нечетных позиций")

	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	printArray(arr, "Исходный массив")

	fmt.Print("Элементы на четных позициях (2, 4, 6...): ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	fmt.Print("Элементы на нечетных позициях (1, 3, 5...): ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func ch11ex32() {
	fmt.Println("11.32. Преобразование массива действительных чисел")

	arr := []float64{-3.5, 2.1, -1.8, 4.2, -5.0, 0.0, 3.7, -2.3}
	printFloatArray(arr, "Исходный массив")

	arrA := make([]float64, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v < 0 {
			arrA[i] = math.Abs(v)
		}
	}
	printFloatArray(arrA, "а) Абсолютные величины отрицательных")

	arrB := make([]float64, len(arr))
	copy(arrB, arr)
	for i := 1; i < len(arrB); i += 2 {
		if arrB[i] >= 0 {
			arrB[i] = math.Sqrt(arrB[i])
		}
	}
	printFloatArray(arrB, "б) Квадратные корни элементов с четными номерами")
}

func ch11ex33() {
	fmt.Println("11.33. Преобразование массива действительных чисел")

	arr := []float64{12.5, 8.2, 15.0, 3.7, 20.1, 9.6, 25.0, 4.9}
	printFloatArray(arr, "Исходный массив")

	arrA := make([]float64, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v > 10 {
			arrA[i] = math.Sqrt(v)
		}
	}
	printFloatArray(arrA, "а) Квадратные корни элементов > 10")

	arrB := make([]float64, len(arr))
	copy(arrB, arr)
	for i := 1; i < len(arrB); i += 2 {
		arrB[i] = math.Abs(arrB[i])
	}
	printFloatArray(arrB, "б) Абсолютные величины элементов с четными номерами")
}

func ch11ex34() {
	fmt.Println("11.34. Преобразование массива действительных чисел")

	arr := []float64{2.5, -3.1, 4.2, -1.8, 5.0, -2.3, 3.7, -4.9}
	printFloatArray(arr, "Исходный массив")

	k1 := 2
	k2 := 4

	arrA := make([]float64, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v > 0 {
			arrA[i] = v - arrA[k1]
		} else {
			arrA[i] = v - arrA[k2]
		}
	}
	printFloatArray(arrA, fmt.Sprintf("а) Вычитание (k1=%d, k2=%d)", k1, k2))

	arrB := make([]float64, len(arr))
	copy(arrB, arr)
	for i := range arrB {
		if i%2 == 0 {
			arrB[i] -= 1
		} else {
			arrB[i] += 1
		}
	}
	printFloatArray(arrB, "б) +/-1 в зависимости от номера")
}

func ch11ex35() {
	fmt.Println("11.35. Преобразование массива действительных чисел")

	arr := []float64{-2.5, 3.1, -4.2, 1.8, -5.0, 2.3, -3.7, 4.9}
	printFloatArray(arr, "Исходный массив")

	m1 := 1
	m2 := 3

	arrA := make([]float64, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v < 0 {
			arrA[i] = v + arrA[m1]
		} else {
			arrA[i] = v + arrA[m2]
		}
	}
	printFloatArray(arrA, fmt.Sprintf("а) Сложение (m1=%d, m2=%d)", m1, m2))

	arrB := make([]float64, len(arr))
	copy(arrB, arr)
	for i := range arrB {
		if i%2 == 0 {
			arrB[i] -= 1
		} else {
			arrB[i] += 1
		}
	}
	printFloatArray(arrB, "б) -/+1 в зависимости от номера")
}

func ch11ex36() {
	fmt.Println("11.36. Преобразование массива действительных чисел")

	arr := []float64{2.5, 0.0, -3.1, 4.2, 0.0, -1.8, 5.0, -2.3}
	printFloatArray(arr, "Исходный массив")

	K1 := 2
	n := 10.0

	arrA := make([]float64, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v > 0 {
			arrA[i] = v - arrA[K1]
		} else if v < 0 {
			arrA[i] = v - n
		}

	}
	printFloatArray(arrA, fmt.Sprintf("а) Преобразование (K1=%d, n=%.1f)", K1, n))

	a := 2.0
	b := 3.0
	arrB := make([]float64, len(arr))
	copy(arrB, arr)
	for i, v := range arrB {
		if v == 0 {
			arrB[i] = v + n
		} else if v > 0 {
			arrB[i] = v - a
		} else {
			arrB[i] = v + b
		}
	}
	printFloatArray(arrB, fmt.Sprintf("б) Преобразование (n=%.1f, a=%.1f, b=%.1f)", n, a, b))
}

func ch11ex37() {
	fmt.Println("11.37. Преобразование массива действительных чисел")

	arr := []float64{-2.5, 0.0, 3.1, -4.2, 0.0, 1.8, -5.0, 2.3}
	printFloatArray(arr, "Исходный массив")

	a1 := 1
	b := 5.0

	arrA := make([]float64, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v < 0 {
			arrA[i] = v + arrA[a1]
		} else if v == 0 {
			arrA[i] = v - b
		}

	}
	printFloatArray(arrA, fmt.Sprintf("а) Преобразование (a1=%d, b=%.1f)", a1, b))

	A := 2.0
	b2 := 3.0
	c := 1.0
	arrB := make([]float64, len(arr))
	copy(arrB, arr)
	for i, v := range arrB {
		if v > 0 {
			arrB[i] = v - A
		} else if v < 0 {
			arrB[i] = v - b2
		} else {
			arrB[i] = v + c
		}
	}
	printFloatArray(arrB, fmt.Sprintf("б) Преобразование (A=%.1f, b=%.1f, c=%.1f)", A, b2, c))
}

func ch11ex38() {
	fmt.Println("11.38. Преобразование массива целых чисел")

	arr := []int{14, 25, 34, 47, 50, 63, 74, 88, 94, 105}
	printArray(arr, "Исходный массив")

	arrA := make([]int, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v%10 == 4 {
			arrA[i] = v / 2
		}
	}
	printArray(arrA, "а) Элементы, оканчивающиеся на 4, уменьшены вдвое")

	arrB := make([]int, len(arr))
	copy(arrB, arr)
	for i, v := range arrB {
		if v%2 == 0 {
			arrB[i] = v * v
		} else {
			arrB[i] = v * 2
		}
	}
	printArray(arrB, "б) Четные - квадраты, нечетные - удвоены")

	a := 5
	b := 3
	arrC := make([]int, len(arr))
	copy(arrC, arr)
	for i, v := range arrC {
		if v%2 == 0 {
			arrC[i] = v + a
		}
		if i%2 == 0 {
			arrC[i] = arrC[i] - b
		}
	}
	printArray(arrC, fmt.Sprintf("в) Преобразование (a=%d, b=%d)", a, b))
}

func ch11ex39() {
	fmt.Println("11.39. Преобразование массива целых чисел")

	arr := []int{10, 25, 30, 47, 50, 63, 70, 88, 90, 105}
	printArray(arr, "Исходный массив")

	arrA := make([]int, len(arr))
	copy(arrA, arr)
	for i, v := range arrA {
		if v%10 == 0 {
			arrA[i] = 0
		}
	}
	printArray(arrA, "а) Элементы, кратные 10, заменены нулями")

	arrB := make([]int, len(arr))
	copy(arrB, arr)
	for i, v := range arrB {
		if v%2 != 0 {
			arrB[i] = v * 2
		} else {
			arrB[i] = v / 2
		}
	}
	printArray(arrB, "б) Нечетные удвоены, четные уменьшены вдвое")

	m := 5
	n := 3
	arrC := make([]int, len(arr))
	copy(arrC, arr)
	for i, v := range arrC {
		if v%2 != 0 {
			arrC[i] = v - m
		}
		if i%2 != 0 {
			arrC[i] = arrC[i] + n
		}
	}
	printArray(arrC, fmt.Sprintf("в) Преобразование (m=%d, n=%d)", m, n))
}

func ch11ex40() {
	fmt.Println("11.40. Расчеты с элементами массива")

	arr := []float64{4.0, 9.0, 16.0, 25.0, 36.0, 49.0, 64.0, 81.0}
	printFloatArray(arr, "Исходный массив")

	fmt.Print("а) Квадратные корни элементов: ")
	for _, v := range arr {
		fmt.Printf("%.2f ", math.Sqrt(v))
	}
	fmt.Println()

	i, j := 2, 5
	average := (arr[i] + arr[j]) / 2
	fmt.Printf("б) Среднее арифметическое элементов arr[%d]=%.1f и arr[%d]=%.1f: %.2f\n",
		i, arr[i], j, arr[j], average)
}

func ch11ex41() {
	fmt.Println("11.41. Проверка элементов массива")

	arr := []int{16, 25, -36, 49, 64, 81, 100, 121, -8, 50}
	printArray(arr, "Исходный массив")

	s := 3
	if s-1 < len(arr) {
		element := arr[s-1]
		isPositive := element > 0
		fmt.Printf("а) %d-й элемент (%d) положительный: %v\n",
			s, element, isPositive)
	}

	k := 5
	if k-1 < len(arr) {
		element := arr[k-1]
		isEven := element%2 == 0
		fmt.Printf("б) %d-й элемент (%d) четный: %v\n",
			k, element, isEven)
	}

	if k-1 < len(arr) && s-1 < len(arr) {
		kElement := arr[k-1]
		sElement := arr[s-1]
		if kElement > sElement {
			fmt.Printf("в) %d-й элемент (%d) больше %d-го (%d)\n",
				k, kElement, s, sElement)
		} else if sElement > kElement {
			fmt.Printf("в) %d-й элемент (%d) больше %d-го (%d)\n",
				s, sElement, k, kElement)
		} else {
			fmt.Printf("в) %d-й и %d-й элементы равны (%d)\n",
				k, s, kElement)
		}
	}
}

func ch11ex42() {
	fmt.Println("11.42. Операции с элементами массива")

	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	printArray(arr, "Исходный массив")

	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Printf("а) Сумма всех элементов: %d\n", sum)

	product := 1
	for _, v := range arr {
		product *= v
	}
	fmt.Printf("б) Произведение всех элементов: %d\n", product)

	sumSquares := 0
	for _, v := range arr {
		sumSquares += v * v
	}
	fmt.Printf("в) Сумма квадратов всех элементов: %d\n", sumSquares)

	sumFirstSix := 0
	for i := 0; i < 6 && i < len(arr); i++ {
		sumFirstSix += arr[i]
	}
	fmt.Printf("г) Сумма шести первых элементов: %d\n", sumFirstSix)

	var k1, k2 int
	fmt.Print("д) Введите k1 и k2 (k2 > k1): ")
	fmt.Scan(&k1, &k2)

	if k1 < 1 || k2 > len(arr) || k2 <= k1 {
		fmt.Println("  Некорректные индексы")
	} else {
		sumRange := 0
		for i := k1 - 1; i < k2; i++ {
			sumRange += arr[i]
		}
		fmt.Printf("  Сумма элементов с %d-го по %d-й: %d\n",
			k1, k2, sumRange)
	}

	average := float64(sum) / float64(len(arr))
	fmt.Printf("е) Среднее арифметическое всех элементов: %.2f\n", average)

	var s1, s2 int
	fmt.Print("ж) Введите s1 и s2 (s2 > s1): ")
	fmt.Scan(&s1, &s2)

	if s1 < 1 || s2 > len(arr) || s2 <= s1 {
		fmt.Println("  Некорректные индексы")
	} else {
		sumRange := 0
		count := 0
		for i := s1 - 1; i < s2; i++ {
			sumRange += arr[i]
			count++
		}
		avgRange := float64(sumRange) / float64(count)
		fmt.Printf("  Среднее арифметическое с %d-го по %d-й: %.2f\n",
			s1, s2, avgRange)
	}
}

func ch11ex43() {
	fmt.Println("11.43. Знакопеременная сумма")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printArray(arr, "Исходный массив")

	alternatingSum := 0
	sign := 1

	for i := 0; i < len(arr); i++ {
		alternatingSum += sign * arr[i]
		sign = -sign
	}

	fmt.Printf("Знакопеременная сумма: %d\n", alternatingSum)
}

func ch11ex44() {
	fmt.Println("11.44. Осадки за январь")

	precipitation := []int{5, 3, 0, 7, 2, 4, 6, 1, 8, 3,
		0, 5, 2, 0, 9, 4, 1, 6, 3, 7,
		2, 8, 0, 5, 1, 4, 6, 2, 3, 8, 5}

	total := 0
	for _, p := range precipitation {
		total += p
	}

	fmt.Printf("Общее количество осадков за январь: %d мм\n", total)
}

func ch11ex45() {
	fmt.Println("11.45. Стоимость предметов")

	prices := []int{150, 230, 450, 120, 680, 320,
		540, 210, 380, 290, 510, 170}

	total := 0
	for _, price := range prices {
		total += price
	}

	fmt.Printf("Общая стоимость 12 предметов: %d руб.\n", total)
}

func ch11ex46() {
	fmt.Println("11.46. Сопротивление цепи (последовательное соединение)")

	resistances := []float64{10, 15, 20, 25, 30, 35, 40, 45, 50, 55,
		60, 65, 70, 75, 80, 85, 90, 95, 100, 105}

	total := 0.0
	for _, r := range resistances {
		total += r
	}

	fmt.Printf("Общее сопротивление (последовательное): %.2f Ом\n", total)
}

func ch11ex47() {
	fmt.Println("11.47. Сопротивление цепи (параллельное соединение)")

	resistances := []float64{10, 15, 20, 25, 30, 35, 40, 45, 50, 55,
		60, 65, 70, 75, 80, 85, 90, 95, 100, 105}

	inverseTotal := 0.0
	for _, r := range resistances {
		if r > 0 {
			inverseTotal += 1.0 / r
		}
	}

	totalResistance := 0.0
	if inverseTotal > 0 {
		totalResistance = 1.0 / inverseTotal
	}

	fmt.Printf("Общее сопротивление (параллельное): %.4f Ом\n", totalResistance)
}

func ch11ex48() {
	fmt.Println("11.48. Осадки по декадам июня")

	precipitation := []int{5, 3, 0, 7, 2, 4, 6, 1, 8, 3,
		0, 5, 2, 0, 9, 4, 1, 6, 3, 7,
		2, 8, 0, 5, 1, 4, 6, 2, 3, 8}

	firstDecade := 0
	for i := 0; i < 10; i++ {
		firstDecade += precipitation[i]
	}

	secondDecade := 0
	for i := 10; i < 20; i++ {
		secondDecade += precipitation[i]
	}

	thirdDecade := 0
	for i := 20; i < 30; i++ {
		thirdDecade += precipitation[i]
	}

	fmt.Printf("Осадки за 1-ю декаду: %d мм\n", firstDecade)
	fmt.Printf("Осадки за 2-ю декаду: %d мм\n", secondDecade)
	fmt.Printf("Осадки за 3-ю декаду: %d мм\n", thirdDecade)
	fmt.Printf("Всего за месяц: %d мм\n", firstDecade+secondDecade+thirdDecade)
}

func ch11ex49() {
	fmt.Println("11.49. Среднедневные осадки в феврале")

	precipitation := []int{5, 3, 0, 7, 2, 4, 6, 1, 8, 3,
		0, 5, 2, 0, 9, 4, 1, 6, 3, 7,
		2, 8, 0, 5, 1, 4, 6, 2}

	total := 0
	for _, p := range precipitation {
		total += p
	}

	average := float64(total) / float64(len(precipitation))
	fmt.Printf("Среднедневное количество осадков в феврале: %.2f мм\n", average)
}

func ch11ex50() {
	fmt.Println("11.50. Средние осадки по декадам сентября")

	precipitation := []int{5, 3, 0, 7, 2, 4, 6, 1, 8, 3,
		0, 5, 2, 0, 9, 4, 1, 6, 3, 7,
		2, 8, 0, 5, 1, 4, 6, 2, 3, 8}

	firstDecadeTotal := 0
	for i := 0; i < 10; i++ {
		firstDecadeTotal += precipitation[i]
	}
	firstDecadeAvg := float64(firstDecadeTotal) / 10.0

	secondDecadeTotal := 0
	for i := 10; i < 20; i++ {
		secondDecadeTotal += precipitation[i]
	}
	secondDecadeAvg := float64(secondDecadeTotal) / 10.0

	thirdDecadeTotal := 0
	for i := 20; i < 30; i++ {
		thirdDecadeTotal += precipitation[i]
	}
	thirdDecadeAvg := float64(thirdDecadeTotal) / 10.0

	fmt.Printf("Средние осадки за 1-ю декаду: %.2f мм/день\n", firstDecadeAvg)
	fmt.Printf("Средние осадки за 2-ю декаду: %.2f мм/день\n", secondDecadeAvg)
	fmt.Printf("Средние осадки за 3-ю декаду: %.2f мм/день\n", thirdDecadeAvg)
}

func ch11ex51() {
	fmt.Println("11.51. Проверка суммы массива")

	arr := []int{3, -5, 2, -1, 4, -2, 6, -3, 5, -4}
	printArray(arr, "Исходный массив")

	sum := 0
	for _, v := range arr {
		sum += v
	}

	isNonNegative := sum >= 0
	fmt.Printf("Сумма элементов: %d\n", sum)
	fmt.Printf("Сумма является неотрицательным числом: %v\n", isNonNegative)
}

func ch11ex52() {
	fmt.Println("11.52. Проверка свойств суммы")

	arr := []int{23, 45, 67, 89, 12, 34, 56, 78, 90, 11}
	printArray(arr, "Исходный массив")

	sum := 0
	for _, v := range arr {
		sum += v
	}
	isEvenSum := sum%2 == 0
	fmt.Printf("а) Сумма элементов (%d) четная: %v\n", sum, isEvenSum)

	sumSquares := 0
	for _, v := range arr {
		sumSquares += v * v
	}
	isFiveDigit := sumSquares >= 10000 && sumSquares <= 99999
	fmt.Printf("б) Сумма квадратов (%d) пятизначная: %v\n",
		sumSquares, isFiveDigit)
}
