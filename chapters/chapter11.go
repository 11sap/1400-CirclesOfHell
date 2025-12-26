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
		ch11ex49, ch11ex50, ch11ex51, ch11ex52, ch11ex53, ch11ex54, ch11ex55, ch11ex56,
		ch11ex57, ch11ex58, ch11ex59, ch11ex60, ch11ex61, ch11ex62, ch11ex63, ch11ex64,
		ch11ex65, ch11ex66, ch11ex67, ch11ex68, ch11ex69, ch11ex70, ch11ex71, ch11ex72,
		ch11ex73, ch11ex74, ch11ex75, ch11ex76, ch11ex77, ch11ex78, ch11ex79, ch11ex80,
		ch11ex81, ch11ex82, ch11ex83, ch11ex84, ch11ex85, ch11ex86, ch11ex87, ch11ex88,
		ch11ex89, ch11ex90, ch11ex91, ch11ex92, ch11ex92, ch11ex93, ch11ex94, ch11ex95,
		ch11ex96, ch11ex97, ch11ex98, ch11ex99, ch11ex100, ch11ex101, ch11ex102, ch11ex103,
		ch11ex104, ch11ex105, ch11ex106, ch11ex107, ch11ex108, ch11ex109, ch11ex110, ch11ex111,
		ch11ex112, ch11ex113, ch11ex114, ch11ex115, ch11ex116, ch11ex117, ch11ex118, ch11ex119,
		ch11ex120, ch11ex121, ch11ex122, ch11ex123, ch11ex124, ch11ex125, ch11ex126, ch11ex127,
		ch11ex128, ch11ex129, ch11ex130, ch11ex131, ch11ex132, ch11ex133,
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
	fmt.Println("Ex1. Заполнение массива из 8 элементов")

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
	fmt.Println("Ex2. Массив чисел от 1 до 10")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}

	printArray(arr[:], "Массив")
}

func ch11ex03() {
	fmt.Println("Ex3. Массив чисел от 10 до 1")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = 10 - i
	}

	printArray(arr[:], "Массив")
}

func ch11ex04() {
	fmt.Println("Ex4. Первые 10 четных чисел")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = 2 * (i + 1)
	}

	printArray(arr[:], "Массив")
}

func ch11ex05() {
	fmt.Println("Ex5. Первые 10 нечетных чисел")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = 2*i + 1
	}

	printArray(arr[:], "Массив")
}

func ch11ex06() {
	fmt.Println("Ex6. Первые 10 чисел, делящихся на 3")

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
	fmt.Println("Ex7. Первые 10 степеней числа 2")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = int(math.Pow(2, float64(i)))
	}

	printArray(arr[:], "Массив")
}

func ch11ex08() {
	fmt.Println("Ex8. Первые 15 чисел Фибоначчи")

	var arr [15]int
	arr[0] = 1
	arr[1] = 1

	for i := 2; i < 15; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	printArray(arr[:], "Массив")
}

func ch11ex09() {
	fmt.Println("Ex9. 10 членов арифметической прогрессии")

	var arr [10]int
	a := 3
	d := 2

	for i := 0; i < 10; i++ {
		arr[i] = a + i*d
	}

	printArray(arr[:], "Массив")
}

func ch11ex10() {
	fmt.Println("Ex10. 10 членов геометрической прогрессии")

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
	fmt.Println("Ex11. Первые 10 точных квадратов")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = (i + 1) * (i + 1)
	}

	printArray(arr[:], "Массив")
}

func ch11ex12() {
	fmt.Println("Ex12. Первые 10 точных кубов")

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = (i + 1) * (i + 1) * (i + 1)
	}

	printArray(arr[:], "Массив")
}

func ch11ex13() {
	fmt.Println("Ex13. Степени числа 2")

	n := 10
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = int(math.Pow(2, float64(i+1)))
	}

	printArray(arr, "Массив")
}

func ch11ex14() {
	fmt.Println("Ex14. Цифры числа в обратном порядке")

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
	fmt.Println("Ex15. Убывающая и возрастающая последовательности")

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
	fmt.Println("Ex16. Арифметическая и геометрическая прогрессии")

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
	fmt.Println("Ex17. Первые 10 чисел Фибоначчи")

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
	fmt.Println("Ex18. Числа, делящиеся на 13 и 17")

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
	fmt.Println("Ex19. Первые 10 простых чисел, начиная с 100")

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
	fmt.Println("Ex20. Проверка таблицы умножения")

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
	fmt.Println("Ex21. Массив из 20 неповторяющихся элементов")

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
	fmt.Println("Ex22. Вывод элементов массива")

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
	fmt.Println("Ex23. Четные элементы и элементы, оканчивающиеся нулем")

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
	fmt.Println("Ex24. Двузначные и трехзначные элементы")

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
	fmt.Println("Ex25. Элементы с четными и нечетными индексами")

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
	fmt.Println("Ex26. Сначала неотрицательные, затем отрицательные элементы")

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
	fmt.Println("Ex27. Сначала четные, затем нечетные элементы")

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
	fmt.Println("Ex28. Индексы элементов, оканчивающихся цифрой 0")

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
	fmt.Println("Ex29. Дни января без осадков")

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
	fmt.Println("Ex30. Команды с менее чем 3 победами")

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
	fmt.Println("Ex31. Элементы с четных и нечетных позиций")

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
	fmt.Println("Ex32. Преобразование массива действительных чисел")

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
	fmt.Println("Ex33. Преобразование массива действительных чисел")

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
	fmt.Println("Ex34. Преобразование массива действительных чисел")

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
	fmt.Println("Ex35. Преобразование массива действительных чисел")

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
	fmt.Println("Ex36. Преобразование массива действительных чисел")

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
	fmt.Println("Ex37. Преобразование массива действительных чисел")

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
	fmt.Println("Ex38. Преобразование массива целых чисел")

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
	fmt.Println("Ex39. Преобразование массива целых чисел")

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
	fmt.Println("Ex40. Расчеты с элементами массива")

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
	fmt.Println("Ex41. Проверка элементов массива")

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
	fmt.Println("Ex42. Операции с элементами массива")

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
	fmt.Println("Ex43. Знакопеременная сумма")

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
	fmt.Println("Ex44. Осадки за январь")

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
	fmt.Println("Ex45. Стоимость предметов")

	prices := []int{150, 230, 450, 120, 680, 320,
		540, 210, 380, 290, 510, 170}

	total := 0
	for _, price := range prices {
		total += price
	}

	fmt.Printf("Общая стоимость 12 предметов: %d руб.\n", total)
}

func ch11ex46() {
	fmt.Println("Ex46. Сопротивление цепи (последовательное соединение)")

	resistances := []float64{10, 15, 20, 25, 30, 35, 40, 45, 50, 55,
		60, 65, 70, 75, 80, 85, 90, 95, 100, 105}

	total := 0.0
	for _, r := range resistances {
		total += r
	}

	fmt.Printf("Общее сопротивление (последовательное): %.2f Ом\n", total)
}

func ch11ex47() {
	fmt.Println("Ex47. Сопротивление цепи (параллельное соединение)")

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
	fmt.Println("Ex48. Осадки по декадам июня")

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
	fmt.Println("Ex49. Среднедневные осадки в феврале")

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
	fmt.Println("Ex50. Средние осадки по декадам сентября")

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
	fmt.Println("Ex51. Проверка суммы массива")

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
	fmt.Println("Ex52. Проверка свойств массива")

	arr := []int{5, 12, 8, 3, 17, 9, 21, 4, 15, 6}

	printArray(arr, "Массив")

	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Printf("а) Сумма элементов: %d\n", sum)
	if sum%2 == 0 {
		fmt.Println("Сумма элементов - четное число (верно)")
	} else {
		fmt.Println("Сумма элементов - нечетное число (неверно)")
	}

	sumSquares := 0
	for _, v := range arr {
		sumSquares += v * v
	}

	fmt.Printf("б) Сумма квадратов элементов: %d\n", sumSquares)
	if sumSquares >= 10000 && sumSquares <= 99999 {
		fmt.Println("Сумма квадратов - пятизначное число (верно)")
	} else {
		fmt.Println("Сумма квадратов - не пятизначное число (неверно)")
	}
}

func ch11ex53() {
	fmt.Println("Ex53. Количество учеников в школе")

	localRand := rand.New(rand.NewSource(99))
	classes := make([]int, 42)

	total := 0
	for i := 0; i < 42; i++ {
		classes[i] = 15 + localRand.Intn(21)
		total += classes[i]
	}

	fmt.Printf("Общее количество учеников: %d\n", total)

	if total >= 1000 && total <= 9999 {
		fmt.Println("Общее число учеников - четырехзначное число (верно)")
	} else {
		fmt.Println("Общее число учеников - не четырехзначное число (неверно)")
	}
}

func ch11ex54() {
	fmt.Println("Ex54. Количество книг в библиотеке")

	localRand := rand.New(rand.NewSource(100))
	sections := make([]int, 35)

	total := 0
	for i := 0; i < 35; i++ {
		sections[i] = 50 + localRand.Intn(251)
		total += sections[i]
	}

	fmt.Printf("Общее количество книг: %d\n", total)

	if total >= 100000 && total <= 999999 {
		fmt.Println("Общее число книг - шестизначное число (верно)")
	} else {
		fmt.Println("Общее число книг - не шестизначное число (неверно)")
	}
}

func ch11ex55() {
	fmt.Println("Ex55. Проверка грузоподъемности автомобиля")

	localRand := rand.New(rand.NewSource(101))
	weights := make([]int, 50)

	carryingCapacity := 5000

	totalWeight := 0
	for i := 0; i < 50; i++ {
		weights[i] = 50 + localRand.Intn(151)
		totalWeight += weights[i]
	}

	fmt.Printf("Грузоподъемность автомобиля: %d кг\n", carryingCapacity)
	fmt.Printf("Общая масса предметов: %d кг\n", totalWeight)

	if totalWeight <= carryingCapacity {
		fmt.Println("Общая масса не превышает грузоподъемность (верно)")
	} else {
		fmt.Println("Общая масса превышает грузоподъемность (неверно)")
	}
}

func ch11ex56() {
	fmt.Println("Ex56. Спортсмен-десятиборец")

	localRand := rand.New(rand.NewSource(102))
	scores := make([]int, 10)

	requiredScore := 700

	totalScore := 0
	for i := 0; i < 10; i++ {
		scores[i] = 50 + localRand.Intn(51)
		totalScore += scores[i]
	}

	printArray(scores, "Баллы по видам спорта")
	fmt.Printf("Общая сумма баллов: %d\n", totalScore)
	fmt.Printf("Необходимая сумма для выхода: %d\n", requiredScore)

	if totalScore > requiredScore {
		fmt.Println("Спортсмен вышел в следующий этап соревнований")
	} else {
		fmt.Println("Спортсмен не вышел в следующий этас соревнований")
	}
}

func ch11ex57() {
	fmt.Println("Ex57. Осадки за июнь")

	localRand := rand.New(rand.NewSource(103))
	precipitation := make([]int, 30)

	for i := 0; i < 30; i++ {
		precipitation[i] = localRand.Intn(26)
	}

	firstHalf := 0
	secondHalf := 0

	for i := 0; i < 15; i++ {
		firstHalf += precipitation[i]
	}

	for i := 15; i < 30; i++ {
		secondHalf += precipitation[i]
	}

	fmt.Println("а) Сравнение первой и второй половины июня:")
	fmt.Printf("  Первая половина (1-15 июня): %d мм\n", firstHalf)
	fmt.Printf("  Вторая половина (16-30 июня): %d мм\n", secondHalf)

	if firstHalf > secondHalf {
		fmt.Println("  В первую половину выпало больше осадков")
	} else if secondHalf > firstHalf {
		fmt.Println("  Во вторую половину выпало больше осадков")
	} else {
		fmt.Println("  В обе половины выпало одинаковое количество осадков")
	}

	decade1 := 0
	decade2 := 0
	decade3 := 0

	for i := 0; i < 10; i++ {
		decade1 += precipitation[i]
	}

	for i := 10; i < 20; i++ {
		decade2 += precipitation[i]
	}

	for i := 20; i < 30; i++ {
		decade3 += precipitation[i]
	}

	fmt.Println("\nб) Сравнение декад:")
	fmt.Printf("  Первая декада (1-10 июня): %d мм\n", decade1)
	fmt.Printf("  Вторая декада (11-20 июня): %d мм\n", decade2)
	fmt.Printf("  Третья декада (21-30 июня): %d мм\n", decade3)

	maxDecade := decade1
	decadeName := "первая"

	if decade2 > maxDecade {
		maxDecade = decade2
		decadeName = "вторая"
	}
	if decade3 > maxDecade {
		maxDecade = decade3
		decadeName = "третья"
	}

	fmt.Printf("  Больше всего осадков выпало в %s декаде: %d мм\n", decadeName, maxDecade)
}

func ch11ex58() {
	fmt.Println("Ex58. Фигурное катание")

	localRand := rand.New(rand.NewSource(104))
	scores := make([]int, 18)

	for i := 0; i < 18; i++ {
		scores[i] = 50 + localRand.Intn(11)
	}

	obligatorySum := 0
	for i := 0; i < 6; i++ {
		obligatorySum += scores[i]
	}

	freeSum := 0
	for i := 6; i < 18; i++ {
		freeSum += scores[i]
	}

	fmt.Printf("Обязательная программа (первые 6 оценок): %d баллов\n", obligatorySum)
	fmt.Printf("Произвольная программа (остальные 12 оценок): %d баллов\n", freeSum)

	if obligatorySum > freeSum {
		fmt.Println("Лучший результат показан в обязательной программе")
	} else if freeSum > obligatorySum {
		fmt.Println("Лучший результат показан в произвольной программе")
	} else {
		fmt.Println("Результаты в обеих программах одинаковы")
	}
}

func ch11ex59() {
	fmt.Println("Ex59. Сумма элементов с условиями")

	arr := []int{15, 25, 8, 30, 12, 40, 5, 18, 22, 10}

	printArray(arr, "Массив")

	sumLess20 := 0
	for _, v := range arr {
		if v <= 20 {
			sumLess20 += v
		}
	}
	fmt.Printf("а) Сумма элементов не превышающих 20: %d\n", sumLess20)

	var a int
	fmt.Print("б) Введите число a: ")
	fmt.Scan(&a)

	sumGreaterA := 0
	for _, v := range arr {
		if v > a {
			sumGreaterA += v
		}
	}
	fmt.Printf("   Сумма элементов больших %d: %d\n", a, sumGreaterA)
}

func ch11ex60() {
	fmt.Println("Ex60. Сумма элементов с разными условиями")

	arr := []int{3, 8, 15, 21, 12, 9, 6, 24, 18, 27}

	printArray(arr, "Массив")

	sumOdd := 0
	for _, v := range arr {
		if v%2 != 0 {
			sumOdd += v
		}
	}
	fmt.Printf("а) Сумма нечетных элементов: %d\n", sumOdd)

	var k int
	fmt.Print("б) Введите число k: ")
	fmt.Scan(&k)

	sumMultipleK := 0
	for _, v := range arr {
		if v%k == 0 {
			sumMultipleK += v
		}
	}
	fmt.Printf("   Сумма элементов кратных %d: %d\n", k, sumMultipleK)

	var a, b int
	fmt.Print("в) Введите число a: ")
	fmt.Scan(&a)
	fmt.Print("   Введите число b: ")
	fmt.Scan(&b)

	sumMultipleAB := 0
	for _, v := range arr {
		if v%a == 0 || v%b == 0 {
			sumMultipleAB += v
		}
	}
	fmt.Printf("   Сумма элементов кратных %d или %d: %d\n", a, b, sumMultipleAB)
}

func ch11ex61() {
	fmt.Println("Ex61. Сумма элементов с четными позициями (2, 4, 6, ...)")

	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	printArray(arr, "Массив")

	sum := 0

	for i := 1; i < len(arr); i += 2 {
		sum += arr[i]
	}

	fmt.Printf("Сумма второго, четвертого, шестого и т.д. элементов: %d\n", sum)
}

func ch11ex62() {
	fmt.Println("Ex62. Осадки по четным числам февраля")

	localRand := rand.New(rand.NewSource(105))
	precipitation := make([]int, 28)

	for i := 0; i < 28; i++ {
		precipitation[i] = localRand.Intn(16)
	}

	totalEvenDays := 0

	for i := 1; i < 28; i += 2 {
		totalEvenDays += precipitation[i]
	}

	fmt.Printf("Общее число осадков, выпавших по четным числам февраля: %d мм\n", totalEvenDays)
}

func ch11ex63() {
	fmt.Println("Ex63. Осадки по месяцам года")

	localRand := rand.New(rand.NewSource(106))
	precipitation := make([]int, 12)
	months := []string{"январь", "февраль", "март", "апрель", "май", "июнь",
		"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"}

	for i := 0; i < 12; i++ {
		precipitation[i] = 20 + localRand.Intn(81)
	}

	total := precipitation[2] + precipitation[5] + precipitation[8] + precipitation[11]

	fmt.Println("Осадки по месяцам:")
	for i := 0; i < 12; i++ {
		fmt.Printf("  %s: %d мм\n", months[i], precipitation[i])
	}

	fmt.Printf("\nОбщее число осадков в марте, июне, сентябре и декабре: %d мм\n", total)
}

func ch11ex64() {
	fmt.Println("Ex64. Частное от деления сумм")

	arr := []int{-5, 12, -8, 3, -17, 9, 21, -4, 15, -6}

	printArray(arr, "Массив")

	sumPositive := 0
	sumNegative := 0

	for _, v := range arr {
		if v > 0 {
			sumPositive += v
		} else if v < 0 {
			sumNegative += v
		}
	}

	fmt.Printf("Сумма положительных элементов: %d\n", sumPositive)
	fmt.Printf("Сумма отрицательных элементов: %d\n", sumNegative)

	if sumNegative != 0 {

		ratio := float64(sumPositive) / float64(abs(sumNegative))
		fmt.Printf("Частное от деления: %.2f\n", ratio)
	} else {
		fmt.Println("Деление на ноль невозможно (сумма отрицательных элементов равна 0)")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ch11ex65() {
	fmt.Println("Ex65. Проверка сумм с условиями")

	arr := []int{25, 15, 30, 45, 10, 60, 5, 35, 20, 55}

	printArray(arr, "Массив")

	sumGreater20 := 0
	for _, v := range arr {
		if v > 20 {
			sumGreater20 += v
		}
	}

	fmt.Printf("а) Сумма элементов больше 20: %d\n", sumGreater20)
	if sumGreater20 > 100 {
		fmt.Println("   Сумма превышает 100 (верно)")
	} else {
		fmt.Println("   Сумма не превышает 100 (неверно)")
	}

	sumLess50 := 0
	for _, v := range arr {
		if v < 50 {
			sumLess50 += v
		}
	}

	fmt.Printf("б) Сумма элементов меньше 50: %d\n", sumLess50)
	if sumLess50%2 == 0 {
		fmt.Println("   Сумма - четное число (верно)")
	} else {
		fmt.Println("   Сумма - нечетное число (неверно)")
	}
}

func ch11ex66() {
	fmt.Println("Ex66. Осадки по четным и нечетным дням февраля")

	localRand := rand.New(rand.NewSource(107))
	precipitation := make([]int, 28)

	for i := 0; i < 28; i++ {
		precipitation[i] = localRand.Intn(21)
	}

	sumEven := 0
	sumOdd := 0

	for i := 0; i < 28; i++ {

		if (i+1)%2 == 0 {
			sumEven += precipitation[i]
		} else {
			sumOdd += precipitation[i]
		}
	}

	fmt.Printf("Осадки по нечетным дням: %d мм\n", sumOdd)
	fmt.Printf("Осадки по четным дням: %d мм\n", sumEven)

	if sumEven > sumOdd {
		fmt.Println("По четным числам выпало больше осадков (верно)")
	} else {
		fmt.Println("По четным числам не выпало больше осадков (неверно)")
	}
}

func ch11ex67() {
	fmt.Println("Ex67. Жители на разных сторонах улицы")

	localRand := rand.New(rand.NewSource(108))
	houses := make([]int, 20)

	for i := 0; i < 20; i++ {
		houses[i] = 10 + localRand.Intn(41)
	}

	sumOddHouses := 0
	sumEvenHouses := 0

	for i := 0; i < 20; i++ {

		if (i+1)%2 == 0 {
			sumEvenHouses += houses[i]
		} else {
			sumOddHouses += houses[i]
		}
	}

	fmt.Printf("Жителей в домах с нечетными номерами: %d\n", sumOddHouses)
	fmt.Printf("Жителей в домах с четными номерами: %d\n", sumEvenHouses)

	if sumOddHouses > sumEvenHouses {
		fmt.Println("На стороне с нечетными номерами проживает больше жителей")
	} else if sumEvenHouses > sumOddHouses {
		fmt.Println("На стороне с четными номерами проживает больше жителей")
	} else {
		fmt.Println("На обеих сторонах проживает одинаковое количество жителей")
	}
}

func ch11ex68() {
	fmt.Println("Ex68. Количество неотрицательных элементов")

	arr := []int{-3, 5, -2, 0, 8, -1, 4, -7, 6, 0}

	printArray(arr, "Массив")

	countNonNegative := 0
	for _, v := range arr {
		if v >= 0 {
			countNonNegative++
		}
	}

	fmt.Printf("Количество неотрицательных элементов: %d\n", countNonNegative)
}

func ch11ex69() {
	fmt.Println("Ex69. Количество элементов с условиями")

	arr := []int{5, 12, 5, 8, 12, 3, 5, 9, 12, 7}

	printArray(arr, "Массив")

	lastElement := arr[len(arr)-1]
	countDiffFromLast := 0
	for _, v := range arr {
		if v != lastElement {
			countDiffFromLast++
		}
	}

	fmt.Printf("а) Элементов отличных от последнего (%d): %d\n", lastElement, countDiffFromLast)

	var a int
	fmt.Print("б) Введите число a: ")
	fmt.Scan(&a)

	countMultipleA := 0
	for _, v := range arr {
		if v%a == 0 {
			countMultipleA++
		}
	}

	fmt.Printf("   Элементов кратных %d: %d\n", a, countMultipleA)
}

func ch11ex70() {
	fmt.Println("Ex70. Дни без осадков в феврале")

	localRand := rand.New(rand.NewSource(109))
	precipitation := make([]int, 28)

	for i := 0; i < 28; i++ {

		if localRand.Intn(100) < 30 {
			precipitation[i] = 0
		} else {
			precipitation[i] = 1 + localRand.Intn(15)
		}
	}

	countNoPrecipitation := 0
	for i := 0; i < 28; i++ {
		if precipitation[i] == 0 {
			countNoPrecipitation++
		}
	}

	fmt.Printf("Количество дней без осадков в феврале: %d\n", countNoPrecipitation)
}

func ch11ex71() {
	fmt.Println("Ex71. Неуспевающие по химии")

	localRand := rand.New(rand.NewSource(110))
	grades := make([]int, 25)

	for i := 0; i < 25; i++ {
		grades[i] = 2 + localRand.Intn(4)
	}

	countFailing := 0
	for _, grade := range grades {
		if grade == 2 {
			countFailing++
		}
	}

	fmt.Printf("Количество неуспевающих по химии: %d из 25 учеников\n", countFailing)
}

func ch11ex72() {
	fmt.Println("Ex72. Продажи товаров в марте")

	localRand := rand.New(rand.NewSource(111))
	sales := make([]int, 31)

	for i := 0; i < 31; i++ {
		sales[i] = 10000 + localRand.Intn(40001)
	}

	var z int
	fmt.Print("Введите значение z: ")
	fmt.Scan(&z)

	countAboveZ := 0
	for i := 0; i < 31; i++ {
		if sales[i] > z {
			countAboveZ++
		}
	}

	fmt.Printf("Количество дней с продажами выше %d: %d\n", z, countAboveZ)
}

func ch11ex73() {
	fmt.Println("Ex73. Рост учеников")

	localRand := rand.New(rand.NewSource(112))
	heights := make([]int, 22)

	for i := 0; i < 22; i++ {
		heights[i] = 150 + localRand.Intn(41)
	}

	var r int
	fmt.Print("Введите значение r (рост в см): ")
	fmt.Scan(&r)

	countBelowR := 0
	for _, height := range heights {
		if height <= r {
			countBelowR++
		}
	}

	fmt.Printf("Количество учеников с ростом не превышающим %d см: %d из 22\n", r, countBelowR)
}

func ch11ex74() {
	fmt.Println("Ex74. Элементы в промежутке")

	arr := []int{15, 8, 25, 3, 12, 30, 5, 18, 22, 9}

	printArray(arr, "Массив")

	var a, b int
	fmt.Print("Введите a: ")
	fmt.Scan(&a)
	fmt.Print("Введите b (должно быть > a): ")
	fmt.Scan(&b)

	if b <= a {
		fmt.Println("Ошибка: b должно быть больше a")
		return
	}

	countInRange := 0
	for _, v := range arr {
		if v >= a && v <= b {
			countInRange++
		}
	}

	fmt.Printf("Количество элементов в промежутке [%d, %d]: %d\n", a, b, countInRange)
}

func ch11ex75() {
	fmt.Println("Ex75. Результаты футбольных игр")

	localRand := rand.New(rand.NewSource(113))
	results := make([]int, 20)

	for i := 0; i < 20; i++ {

		r := localRand.Intn(100)
		if r < 40 {
			results[i] = 3
		} else if r < 70 {
			results[i] = 1
		} else {
			results[i] = 0
		}
	}

	wins := 0
	draws := 0

	for _, result := range results {
		if result == 3 {
			wins++
		} else if result == 1 {
			draws++
		}
	}

	fmt.Printf("Количество выигрышей: %d\n", wins)
	fmt.Printf("Количество ничьих: %d\n", draws)
	fmt.Printf("Общее количество выигрышей и ничьих: %d\n", wins+draws)
}

func ch11ex76() {
	fmt.Println("Ex76. Оценки ученика")

	localRand := rand.New(rand.NewSource(114))
	grades := make([]int, 10)

	for i := 0; i < 10; i++ {
		grades[i] = 2 + localRand.Intn(4)
	}

	countFours := 0
	countFives := 0

	for _, grade := range grades {
		if grade == 4 {
			countFours++
		} else if grade == 5 {
			countFives++
		}
	}

	printArray(grades, "Оценки")
	fmt.Printf("Количество четверок: %d\n", countFours)
	fmt.Printf("Количество пятерок: %d\n", countFives)
	fmt.Printf("Общее количество четверок и пятерок: %d\n", countFours+countFives)
}

func ch11ex77() {
	fmt.Printf("Ex77. Определить количество положительных и отрицательных элементов массива.\n")
	arr := []int{-2, 5, -8, 10, 3, -1, 0, 7}
	pos, neg := 0, 0
	for _, v := range arr {
		if v > 0 {
			pos++
		} else if v < 0 {
			neg++
		}
	}
	fmt.Printf("Положительных: %d, Отрицательных: %d\n", pos, neg)
}

func ch11ex78() {
	fmt.Printf("Ex78. Дан массив целых чисел. Определить количество четных элементов и элементов, оканчивающихся на цифру 5.\n")
	arr := []int{12, 5, 8, 15, 24, 35, 42, 55}
	even, endsWith5 := 0, 0
	for _, v := range arr {
		if v%2 == 0 {
			even++
		}
		if v%10 == 5 {
			endsWith5++
		}
	}
	fmt.Printf("Четных: %d, Оканчивающихся на 5: %d\n", even, endsWith5)
}

func ch11ex79() {
	fmt.Printf("Ex79. Определить количество выигрышей, ничьих и проигрышей футбольной команды.\n")
	results := []int{3, 1, 3, 2, 1, 3, 2, 1, 3, 3, 2, 1, 3, 1, 2, 3, 1, 2, 3, 1}
	wins, draws, losses := 0, 0, 0
	for _, r := range results {
		switch r {
		case 3:
			wins++
		case 1:
			draws++
		case 2:
			losses++
		}
	}
	fmt.Printf("Выигрышей: %d, Ничьих: %d, Проигрышей: %d\n", wins, draws, losses)
}

func ch11ex80() {
	fmt.Printf("Ex80. Определить количество пятерок, четверок, троек и двоек.\n")
	grades := []int{5, 4, 3, 5, 2, 4, 5, 3, 4, 5, 2, 3, 4, 5, 5, 4, 3, 2, 4, 5, 3, 4}
	count := map[int]int{5: 0, 4: 0, 3: 0, 2: 0}
	for _, g := range grades {
		if g >= 2 && g <= 5 {
			count[g]++
		}
	}
	fmt.Printf("Пятерок: %d, Четверок: %d, Троек: %d, Двоек: %d\n", count[5], count[4], count[3], count[2])
}

func ch11ex81() {
	fmt.Printf("Ex81. Найти число пар соседних элементов массива, являющихся четными числами.\n")
	arr := []int{2, 4, 6, 8, 10, 12, 14}
	pairs := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]%2 == 0 && arr[i+1]%2 == 0 {
			pairs++
		}
	}
	fmt.Printf("Пар четных соседей: %d\n", pairs)
}

func ch11ex82() {
	fmt.Printf("Ex82. Найти число пар соседних элементов массива, оканчивающихся нулем.\n")
	arr := []int{10, 20, 30, 40, 50, 60, 70}
	pairs := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]%10 == 0 && arr[i+1]%10 == 0 {
			pairs++
		}
	}
	fmt.Printf("Пар соседей, оканчивающихся нулем: %d\n", pairs)
}

func ch11ex83() {
	fmt.Printf("Ex83. Найти число элементов массива, которые больше своих «соседей».\n")
	arr := []int{1, 3, 2, 4, 1, 5, 3}
	count := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			count++
		}
	}
	fmt.Printf("Элементов больше соседей: %d\n", count)
}

func ch11ex84() {
	fmt.Printf("Ex84. Проверить условия для массива вещественных чисел.\n")
	arr := []float64{10.5, -3.2, 45.0, 67.8, 12.1, 99.9, 30.2, 50.55, 25.0, 60.0}
	posCount, leCount := 0, 0
	for _, v := range arr {
		if v > 0 {
			posCount++
		}
		if v <= 50.55 {
			leCount++
		}
	}
	a := posCount <= 5
	b := leCount%4 == 0
	fmt.Printf("а) Положительных не более 5: %s\n", map[bool]string{true: "да", false: "нет"}[a])
	fmt.Printf("б) Элементов <=50.55 кратно 4: %s\n", map[bool]string{true: "да", false: "нет"}[b])
}

func ch11ex85() {
	fmt.Printf("Ex85. Определить, можно ли сформировать баскетбольную команду.\n")
	heights := []int{165, 180, 172, 168, 185, 175, 190, 162, 178, 169, 174, 182, 167, 171, 176}
	tallCount := 0
	for _, h := range heights {
		if h > 170 {
			tallCount++
		}
	}
	canForm := tallCount >= 5
	fmt.Printf("Рост >170 см: %d чел. Можно сформировать команду: %s\n",
		tallCount, map[bool]string{true: "да", false: "нет"}[canForm])
}

func ch11ex86() {
	fmt.Printf("Ex86. Проверить, было ли 10 дней без осадков в марте.\n")
	precipitation := []int{0, 5, 0, 0, 3, 0, 0, 8, 0, 0, 2, 0, 0, 4, 0, 0, 1, 0, 0, 6, 0, 0, 7, 0, 0, 9, 0, 0, 0, 0, 0}
	dryDays := 0
	for _, p := range precipitation {
		if p == 0 {
			dryDays++
		}
	}
	has10DryDays := dryDays >= 10
	fmt.Printf("Дней без осадков: %d. Верно ли, что >=10 дней: %s\n",
		dryDays, map[bool]string{true: "да", false: "нет"}[has10DryDays])
}

func ch11ex87() {
	fmt.Printf("Ex87. Найти среднее арифметическое элементов массива, больших числа 10.\n")
	arr := []int{5, 15, 8, 20, 3, 12, 25, 9, 18}
	sum, count := 0, 0
	for _, v := range arr {
		if v > 10 {
			sum += v
			count++
		}
	}
	avg := 0.0
	if count > 0 {
		avg = float64(sum) / float64(count)
	}
	fmt.Printf("Среднее элементов >10: %.2f\n", avg)
}

func ch11ex88() {
	fmt.Printf("Ex88. Найти среднее арифметическое элементов массива, меньших числа m.\n")
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	m := 45
	sum, count := 0, 0
	for _, v := range arr {
		if v < m {
			sum += v
			count++
		}
	}
	avg := 0.0
	if count > 0 {
		avg = float64(sum) / float64(count)
	}
	fmt.Printf("Среднее элементов <%d: %.2f\n", m, avg)
}

func ch11ex89() {
	fmt.Printf("Ex89. Определить среднее количество осадков в дождливые дни.\n")
	august := []float64{0.0, 5.2, 0.0, 3.8, 2.1, 0.0, 7.5, 0.0, 4.3, 6.2, 0.0, 1.5, 0.0, 9.1, 0.0}
	sum, rainyDays := 0.0, 0
	for _, p := range august {
		if p > 0 {
			sum += p
			rainyDays++
		}
	}
	avg := 0.0
	if rainyDays > 0 {
		avg = sum / float64(rainyDays)
	}
	fmt.Printf("Среднее в дождливые дни: %.2f мм\n", avg)
}

func ch11ex90() {
	fmt.Printf("Ex90. Найти средние арифметические положительных и отрицательных элементов.\n")
	arr := []int{-5, 10, -3, 8, -2, 15, -7, 20}
	posSum, posCount := 0, 0
	negSum, negCount := 0, 0
	for _, v := range arr {
		if v > 0 {
			posSum += v
			posCount++
		} else if v < 0 {
			negSum += v
			negCount++
		}
	}
	posAvg, negAvg := 0.0, 0.0
	if posCount > 0 {
		posAvg = float64(posSum) / float64(posCount)
	}
	if negCount > 0 {
		negAvg = float64(negSum) / float64(negCount)
	}
	fmt.Printf("Среднее положительных: %.2f, среднее отрицательных: %.2f\n", posAvg, negAvg)
}

func ch11ex91() {
	fmt.Printf("Ex91. Определить среднюю массу полных людей и остальных.\n")
	weights := []int{65, 110, 85, 120, 70, 105, 90, 115, 80, 102, 75, 108, 95, 130, 78}
	fatSum, fatCount := 0, 0
	otherSum, otherCount := 0, 0
	for _, w := range weights {
		if w > 100 {
			fatSum += w
			fatCount++
		} else {
			otherSum += w
			otherCount++
		}
	}
	fatAvg, otherAvg := 0.0, 0.0
	if fatCount > 0 {
		fatAvg = float64(fatSum) / float64(fatCount)
	}
	if otherCount > 0 {
		otherAvg = float64(otherSum) / float64(otherCount)
	}
	fmt.Printf("Средняя масса полных: %.2f кг, остальных: %.2f кг\n", fatAvg, otherAvg)
}

func ch11ex92() {
	fmt.Printf("Ex92. Определить средний рост мальчиков и девочек.\n")
	heights := []int{150, -165, 155, -170, 160, -175, 158, -180, 162, -168, 152, -172}
	boysSum, boysCount := 0, 0
	girlsSum, girlsCount := 0, 0
	for _, h := range heights {
		if h < 0 {
			boysSum += -h
			boysCount++
		} else {
			girlsSum += h
			girlsCount++
		}
	}
	boysAvg, girlsAvg := 0.0, 0.0
	if boysCount > 0 {
		boysAvg = float64(boysSum) / float64(boysCount)
	}
	if girlsCount > 0 {
		girlsAvg = float64(girlsSum) / float64(girlsCount)
	}
	fmt.Printf("Средний рост мальчиков: %.1f см, девочек: %.1f см\n", boysAvg, girlsAvg)
}

func ch11ex93() {
	fmt.Printf("Ex93. Проверить, превышает ли средняя стоимость автомобилей среднюю стоимость мотоциклов более чем в 3 раза.\n")
	cars := []float64{15000, 20000, 18000, 25000, 22000}
	bikes := []float64{5000, 6000, 4000, 5500, 4500, 3500}
	carsSum, bikesSum := 0.0, 0.0
	for _, v := range cars {
		carsSum += v
	}
	for _, v := range bikes {
		bikesSum += v
	}
	carsAvg := carsSum / float64(len(cars))
	bikesAvg := bikesSum / float64(len(bikes))
	condition := bikesAvg > 0 && carsAvg/bikesAvg > 3
	fmt.Printf("Средняя авто: %.2f$, мото: %.2f$. Превышает более чем в 3 раза: %s\n",
		carsAvg, bikesAvg, map[bool]string{true: "да", false: "нет"}[condition])
}

func ch11ex94() {
	fmt.Printf("Ex94. Проверить, превышает ли средний рост мальчиков средний рост девочек более чем на 10 см.\n")
	heights := []int{-175, 160, -180, 155, -170, 165, -185, 158, -172, 162}
	boysSum, boysCount := 0, 0
	girlsSum, girlsCount := 0, 0
	for _, h := range heights {
		if h < 0 {
			boysSum += -h
			boysCount++
		} else {
			girlsSum += h
			girlsCount++
		}
	}
	boysAvg := float64(boysSum) / float64(boysCount)
	girlsAvg := float64(girlsSum) / float64(girlsCount)
	condition := boysAvg-girlsAvg > 10
	fmt.Printf("Средний рост мальчиков: %.1f см, девочек: %.1f см. Разница >10 см: %s\n",
		boysAvg, girlsAvg, map[bool]string{true: "да", false: "нет"}[condition])
}

func ch11ex95() {
	fmt.Printf("Ex95. Определить количество элементов, больших суммы всех элементов массива, и их номера.\n")
	arr := []int{1, 2, 3, 10, 4, 5}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	var indices []int
	for i, v := range arr {
		if v > sum {
			indices = append(indices, i)
		}
	}
	fmt.Printf("Элементов больше суммы (%d): %d. Индексы: %v\n", sum, len(indices), indices)
}

func ch11ex96() {
	fmt.Printf("Ex96. Найти элементы, большие среднего арифметического min и max.\n")
	arr := []int{10, 5, 8, 15, 3, 12, 7}
	if len(arr) == 0 {
		fmt.Println("Массив пуст")
		return
	}
	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	avgMinMax := float64(minVal+maxVal) / 2
	var indices []int
	for i, v := range arr {
		if float64(v) > avgMinMax {
			indices = append(indices, i)
		}
	}
	fmt.Printf("Среднее min(%d) и max(%d): %.1f. Элементов больше: %d. Индексы: %v\n",
		minVal, maxVal, avgMinMax, len(indices), indices)
}

func ch11ex97() {
	fmt.Printf("Ex97. Определить количество учеников с ростом больше среднего.\n")
	heights := []int{165, 170, 175, 160, 180, 172, 168, 178, 182, 167}
	sum := 0
	for _, h := range heights {
		sum += h
	}
	avg := float64(sum) / float64(len(heights))
	count := 0
	for _, h := range heights {
		if float64(h) > avg {
			count++
		}
	}
	fmt.Printf("Средний рост: %.1f см. Выше среднего: %d учеников\n", avg, count)
}

func ch11ex98() {
	fmt.Printf("Ex98. Определить, сколько видов товара имеют стоимость меньше средней.\n")
	prices := []float64{100, 200, 150, 300, 250, 180, 220, 170, 280, 120}
	sum := 0.0
	for _, p := range prices {
		sum += p
	}
	avg := sum / float64(len(prices))
	count := 0
	for _, p := range prices {
		if p < avg {
			count++
		}
	}
	fmt.Printf("Средняя цена: %.2f. Товаров дешевле средней: %d\n", avg, count)
}

func ch11ex99() {
	fmt.Printf("Ex99. Определить дни с осадками выше среднего.\n")
	precipitation := []int{5, 0, 8, 3, 10, 2, 7, 4, 9, 1, 6}
	sum := 0
	for _, p := range precipitation {
		sum += p
	}
	avg := float64(sum) / float64(len(precipitation))
	days := []int{}
	for i, p := range precipitation {
		if float64(p) > avg {
			days = append(days, i+1)
		}
	}
	fmt.Printf("Среднее за день: %.1f мм. Дней выше среднего: %d. Даты: %v\n", avg, len(days), days)
}

func ch11ex100() {
	fmt.Printf("Ex100. Определить учеников с оценкой ниже средней и их индексы.\n")
	grades := []int{5, 4, 3, 5, 2, 4, 3, 5, 4, 3, 2, 4, 5, 3, 4}
	sum := 0
	for _, g := range grades {
		sum += g
	}
	avg := float64(sum) / float64(len(grades))
	indices := []int{}
	for i, g := range grades {
		if float64(g) < avg {
			indices = append(indices, i)
		}
	}
	fmt.Printf("Средняя оценка: %.2f. Ниже средней: %d учеников. Индексы: %v\n", avg, len(indices), indices)
}

func ch11ex101() {
	fmt.Printf("Ex101. Вычислить среднее количество осадков и отклонение от среднего для каждого года.\n")
	precipitation := []float64{600, 550, 700, 650, 580, 720, 680, 630, 590, 710, 670, 640, 690, 610, 660}
	sum := 0.0
	for _, p := range precipitation {
		sum += p
	}
	avg := sum / float64(len(precipitation))
	fmt.Printf("Среднее за 15 лет: %.1f мм\n", avg)
	fmt.Println("Отклонения от среднего:")
	for i, p := range precipitation {
		deviation := p - avg
		fmt.Printf("  Год %d: %.1f мм\n", i+1, deviation)
	}
}

func ch11ex102() {
	fmt.Printf("Ex102. Найти элемент, наиболее близкий к среднему значению.\n")
	arr := []float64{10.5, 12.1, 8.9, 15.2, 11.8, 9.5, 14.3}
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	avg := sum / float64(len(arr))
	closest := arr[0]
	minDiff := math.Abs(arr[0] - avg)
	for _, v := range arr[1:] {
		diff := math.Abs(v - avg)
		if diff < minDiff {
			minDiff = diff
			closest = v
		}
	}
	fmt.Printf("Среднее: %.2f. Наиболее близкий элемент: %.2f\n", avg, closest)
}

func ch11ex103() {
	fmt.Printf("Ex103. Найти пять соседних элементов с максимальной суммой.\n")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	n := 5
	maxSum := 0
	startIndex := 0
	for i := 0; i <= len(arr)-n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += arr[i+j]
		}
		if sum > maxSum {
			maxSum = sum
			startIndex = i
		}
	}
	fmt.Printf("Максимальная сумма %d элементов: %d (индексы %d-%d)\n",
		n, maxSum, startIndex, startIndex+n-1)
}

func ch11ex104() {
	fmt.Printf("Ex104. Определить самые теплые 7 дней подряд.\n")
	temps := []float64{20.5, 22.1, 21.8, 23.5, 24.2, 22.9, 23.8, 24.5, 25.1, 24.8, 23.9, 22.5}
	n := 7
	maxSum := 0.0
	startIndex := 0
	for i := 0; i <= len(temps)-n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += temps[i+j]
		}
		if sum > maxSum {
			maxSum = sum
			startIndex = i
		}
	}
	fmt.Printf("Самые теплые %d дней: с %d по %d число (средняя: %.1f°C)\n",
		n, startIndex+1, startIndex+n, maxSum/float64(n))
}

func ch11ex105() {
	fmt.Printf("Ex105. Определить, имеются ли в массиве одинаковые элементы.\n")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 11}
	seen := make(map[int]bool)
	hasDuplicates := false
	for _, v := range arr {
		if seen[v] {
			hasDuplicates = true
			break
		}
		seen[v] = true
	}
	fmt.Printf("Есть одинаковые элементы: %s\n",
		map[bool]string{true: "да", false: "нет"}[hasDuplicates])
}

func ch11ex106() {
	fmt.Printf("Ex106. Определить, имеются ли в массиве только два одинаковых элемента.\n")
	arr := []int{1, 2, 3, 4, 5, 3, 6, 7}
	count := make(map[int]int)
	for _, v := range arr {
		count[v]++
	}
	duplicates := 0
	for _, c := range count {
		if c == 2 {
			duplicates++
		}
	}
	hasExactlyTwo := duplicates == 1
	fmt.Printf("Только два одинаковых элемента: %s\n",
		map[bool]string{true: "да", false: "нет"}[hasExactlyTwo])
}

func ch11ex107() {
	fmt.Printf("Ex107. Найти два одинаковых элемента в массиве.\n")
	arr := []int{10, 20, 30, 40, 50, 30, 60, 70}
	seen := make(map[int]bool)
	duplicate := 0
	for _, v := range arr {
		if seen[v] {
			duplicate = v
			break
		}
		seen[v] = true
	}
	fmt.Printf("Одинаковый элемент: %d\n", duplicate)
}

func ch11ex108() {
	fmt.Printf("Ex108. Определить, сколько раз элементы меняют знак.\n")
	arr := []int{10, -4, 12, 56, -4, -89}
	signChanges := 0
	for i := 1; i < len(arr); i++ {
		if (arr[i-1] > 0 && arr[i] < 0) || (arr[i-1] < 0 && arr[i] > 0) {
			signChanges++
		}
	}
	fmt.Printf("Смена знака происходит %d раз\n", signChanges)
}

func ch11ex109() {
	fmt.Printf("Ex109. Найти количество одинаковых элементов подряд и различных чисел.\n")
	arr := []int{1, 1, 2, 2, 2, 3, 4, 4, 4, 4, 5, 6, 6}
	identicalPairs := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			identicalPairs++
		}
	}
	unique := make(map[int]bool)
	for _, v := range arr {
		unique[v] = true
	}
	fmt.Printf("Одинаковых пар подряд: %d, различных чисел: %d\n", identicalPairs, len(unique))
}

func ch11ex110() {
	fmt.Printf("Ex110. Найти количество различных чисел в неубывающей последовательности.\n")
	arr := []int{1, 1, 2, 2, 3, 3, 3, 4, 5, 5, 6, 6, 6, 6, 7}
	distinctCount := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			distinctCount++
		}
	}
	fmt.Printf("Различных чисел: %d\n", distinctCount)
}

func ch11ex111() {
	fmt.Printf("Ex111. Найти наибольшую длину отрезка из нечетных чисел.\n")
	arr := []int{1, 3, 5, 2, 7, 9, 11, 4, 13, 15, 17, 6, 19}
	maxLen := 0
	currentLen := 0
	for _, v := range arr {
		if v%2 != 0 {
			currentLen++
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 0
		}
	}
	fmt.Printf("Наибольшая длина отрезка из нечетных чисел: %d\n", maxLen)
}

func ch11ex112() {
	fmt.Printf("Ex112. Определить min, max, их индексы и сравнение.\n")
	arr := []int{12, 45, 3, 67, 23, 89, 34, 2, 56, 78}
	if len(arr) == 0 {
		fmt.Println("Массив пуст")
		return
	}
	minVal, maxVal := arr[0], arr[0]
	minIdx, maxIdx := 0, 0
	for i, v := range arr {
		if v < minVal {
			minVal = v
			minIdx = i
		}
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}
	fmt.Printf("Минимальный: arr[%d]=%d\n", minIdx, minVal)
	fmt.Printf("Максимальный: arr[%d]=%d\n", maxIdx, maxVal)
	fmt.Printf("Максимальный больше минимального: %s\n",
		map[bool]string{true: "да", false: "нет"}[maxVal > minVal])
}

func ch11ex113() {
	fmt.Printf("Ex113. Определить количество страниц в самой толстой книге.\n")
	pages := []int{300, 450, 200, 600, 350, 500, 250, 400}
	maxPages := 0
	for _, p := range pages {
		if p > maxPages {
			maxPages = p
		}
	}
	fmt.Printf("Самая толстая книга: %d страниц\n", maxPages)
}

func ch11ex114() {
	fmt.Printf("Ex114. Определить стоимость самого дорогого автомобиля.\n")
	prices := []float64{15000, 25000, 18000, 30000, 22000, 35000, 28000}
	maxPrice := 0.0
	for _, p := range prices {
		if p > maxPrice {
			maxPrice = p
		}
	}
	fmt.Printf("Самый дорогой автомобиль: %.2f $\n", maxPrice)
}

func ch11ex115() {
	fmt.Printf("Ex115. Определить стоимость самых дешевых конфет.\n")
	candyPrices := []float64{5.50, 4.20, 6.80, 3.90, 5.00, 7.20, 4.50}
	minPrice := candyPrices[0]
	for _, p := range candyPrices[1:] {
		if p < minPrice {
			minPrice = p
		}
	}
	fmt.Printf("Самые дешевые конфеты: %.2f $/кг\n", minPrice)
}

func ch11ex116() {
	fmt.Printf("Ex116. Определить результат спортсмена-победителя лыжной гонки.\n")
	results := []float64{58.3, 59.1, 57.8, 60.2, 58.9, 57.5, 59.8, 58.0}
	bestTime := results[0]
	for _, t := range results[1:] {
		if t < bestTime {
			bestTime = t
		}
	}
	fmt.Printf("Результат победителя: %.1f мин\n", bestTime)
}

func ch11ex117() {
	fmt.Printf("Ex117. Определить разницу между ростом самого высокого и самого низкого.\n")
	heights := []int{165, 180, 172, 155, 185, 168, 178, 160}
	min, max := heights[0], heights[0]
	for _, h := range heights[1:] {
		if h < min {
			min = h
		}
		if h > max {
			max = h
		}
	}
	diff := max - min
	fmt.Printf("Рост самого высокого: %d см, самого низкого: %d см, разница: %d см\n", max, min, diff)
}

func ch11ex118() {
	fmt.Printf("Ex118. Определить разницу в возрасте самого старого и самого молодого.\n")
	birthYears := []int{1990, 1985, 2000, 1995, 1988, 1992, 1980, 1998}
	currentYear := 2024
	minAge, maxAge := currentYear-birthYears[0], currentYear-birthYears[0]
	for _, year := range birthYears[1:] {
		age := currentYear - year
		if age < minAge {
			minAge = age
		}
		if age > maxAge {
			maxAge = age
		}
	}
	ageDiff := maxAge - minAge
	fmt.Printf("Самому молодому: %d лет, самому старому: %d лет, разница: %d лет\n", minAge, maxAge, ageDiff)
}

func ch11ex119() {
	fmt.Printf("Ex119. Рассчитать зачетную оценку фигуриста.\n")
	scores := []float64{5.8, 5.9, 5.7, 6.0, 5.6, 5.8, 5.9, 5.7}
	if len(scores) < 3 {
		fmt.Println("Недостаточно оценок")
		return
	}
	minScore, maxScore := scores[0], scores[0]
	minIdx, maxIdx := 0, 0
	for i, s := range scores {
		if s < minScore {
			minScore = s
			minIdx = i
		}
		if s > maxScore {
			maxScore = s
			maxIdx = i
		}
	}
	sum := 0.0
	count := 0
	for i, s := range scores {
		if i != minIdx && i != maxIdx {
			sum += s
			count++
		}
	}
	finalScore := sum / float64(count)
	fmt.Printf("Зачетная оценка: %.2f\n", finalScore)
}

func ch11ex120() {
	fmt.Printf("Ex120. Определить порядковый номер самого быстрого автомобиля.\n")
	speeds := []int{180, 200, 190, 200, 210, 190, 200, 195}
	maxSpeed := speeds[0]
	for _, s := range speeds {
		if s > maxSpeed {
			maxSpeed = s
		}
	}
	firstIdx, lastIdx := -1, -1
	for i, s := range speeds {
		if s == maxSpeed {
			if firstIdx == -1 {
				firstIdx = i
			}
			lastIdx = i
		}
	}
	fmt.Printf("Максимальная скорость: %d км/ч\n", maxSpeed)
	fmt.Printf("Первый с такой скоростью: индекс %d\n", firstIdx)
	fmt.Printf("Последний с такой скоростью: индекс %d\n", lastIdx)
}

func ch11ex121() {
	fmt.Printf("Ex121. Определить дату самого дождливого дня.\n")
	precipitation := []int{5, 8, 3, 12, 7, 12, 4, 9, 6, 10}
	maxPrecip := precipitation[0]
	for _, p := range precipitation {
		if p > maxPrecip {
			maxPrecip = p
		}
	}
	firstDay, lastDay := -1, -1
	for i, p := range precipitation {
		if p == maxPrecip {
			if firstDay == -1 {
				firstDay = i + 1
			}
			lastDay = i + 1
		}
	}
	fmt.Printf("Максимальные осадки: %d мм\n", maxPrecip)
	fmt.Printf("Первый самый дождливый день: %d июля\n", firstDay)
	fmt.Printf("Последний самый дождливый день: %d июля\n", lastDay)
}

func ch11ex122() {
	fmt.Printf("Ex122. Определить номер самого дешевого вида конфет.\n")
	prices := []float64{5.50, 4.20, 3.90, 4.50, 3.90, 6.00, 4.80}
	minPrice := prices[0]
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
	}
	firstIdx, lastIdx := -1, -1
	for i, p := range prices {
		if p == minPrice {
			if firstIdx == -1 {
				firstIdx = i
			}
			lastIdx = i
		}
	}
	fmt.Printf("Минимальная цена: %.2f $/кг\n", minPrice)
	fmt.Printf("Первый самый дешевый: индекс %d\n", firstIdx)
	fmt.Printf("Последний самый дешевый: индекс %d\n", lastIdx)
}

func ch11ex123() {
	fmt.Printf("Ex123. Определить номер самого старшего человека.\n")
	birthYears := []int{1990, 1985, 1980, 1995, 1980, 1992, 1980, 1998}
	currentYear := 2024
	maxAge := 0
	for _, year := range birthYears {
		age := currentYear - year
		if age > maxAge {
			maxAge = age
		}
	}
	firstIdx, lastIdx := -1, -1
	for i, year := range birthYears {
		if currentYear-year == maxAge {
			if firstIdx == -1 {
				firstIdx = i
			}
			lastIdx = i
		}
	}
	fmt.Printf("Максимальный возраст: %d лет\n", maxAge)
	fmt.Printf("Первый самый старший: индекс %d\n", firstIdx)
	fmt.Printf("Последний самый старший: индекс %d\n", lastIdx)
}

func ch11ex124() {
	fmt.Printf("Ex124. Определить количество максимальных и минимальных элементов.\n")
	arr := []int{5, 8, 3, 8, 2, 8, 3, 2, 8, 1}
	if len(arr) == 0 {
		fmt.Println("Массив пуст")
		return
	}
	minVal, maxVal := arr[0], arr[0]
	minCount, maxCount := 0, 0
	for _, v := range arr {
		if v < minVal {
			minVal = v
			minCount = 1
		} else if v == minVal {
			minCount++
		}
		if v > maxVal {
			maxVal = v
			maxCount = 1
		} else if v == maxVal {
			maxCount++
		}
	}
	fmt.Printf("Минимальных элементов: %d (значение %d)\n", minCount, minVal)
	fmt.Printf("Максимальных элементов: %d (значение %d)\n", maxCount, maxVal)
}

func ch11ex125() {
	fmt.Printf("Ex125. Определить количество самых высоких людей.\n")
	heights := []int{180, 175, 185, 182, 185, 178, 185, 180}
	maxHeight := heights[0]
	maxCount := 0
	for _, h := range heights {
		if h > maxHeight {
			maxHeight = h
			maxCount = 1
		} else if h == maxHeight {
			maxCount++
		}
	}
	fmt.Printf("Самый высокий рост: %d см. Таких людей: %d\n", maxHeight, maxCount)
}

func ch11ex126() {
	fmt.Printf("Ex126. Определить количество дней с осадками более 5 мм.\n")
	precipitation := []int{3, 8, 0, 12, 4, 6, 2, 9, 7, 10}
	threshold := 5
	count := 0
	for _, p := range precipitation {
		if p > threshold {
			count++
		}
	}
	fmt.Printf("Дней с осадками > %d мм: %d\n", threshold, count)
}

func ch11ex127() {
	fmt.Printf("Ex127. Определить количество товаров с ценой выше 100.\n")
	prices := []float64{50, 120, 80, 150, 90, 200, 110, 70}
	threshold := 100.0
	count := 0
	for _, p := range prices {
		if p > threshold {
			count++
		}
	}
	fmt.Printf("Товаров дороже %.2f: %d\n", threshold, count)
}

func ch11ex128() {
	fmt.Printf("Ex128. Определить количество дней с температурой выше 20°C.\n")
	temperatures := []float64{18.5, 22.1, 19.8, 25.3, 20.5, 23.7, 19.2, 21.8}
	threshold := 20.0
	count := 0
	for _, t := range temperatures {
		if t > threshold {
			count++
		}
	}
	fmt.Printf("Дней теплее %.1f°C: %d\n", threshold, count)
}

func ch11ex129() {
	fmt.Printf("Ex129. Вывести номера всех минимальных и максимальных элементов.\n")
	arr := []int{5, 8, 3, 8, 2, 8, 3, 2, 8, 1}
	if len(arr) == 0 {
		fmt.Println("Массив пуст")
		return
	}
	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	minIndices := []int{}
	maxIndices := []int{}
	for i, v := range arr {
		if v == minVal {
			minIndices = append(minIndices, i)
		}
		if v == maxVal {
			maxIndices = append(maxIndices, i)
		}
	}
	fmt.Printf("Индексы минимальных элементов (%d): %v\n", minVal, minIndices)
	fmt.Printf("Индексы максимальных элементов (%d): %v\n", maxVal, maxIndices)
}

func ch11ex130() {
	fmt.Printf("Ex130. Проверить условия сравнения min и max.\n")
	arr := []float64{10.5, 25.8, 15.3, 30.2, 12.7, 28.9}
	if len(arr) == 0 {
		fmt.Println("Массив пуст")
		return
	}
	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	a := maxVal > minVal*1.2
	b := minVal < maxVal/1.2
	fmt.Printf("Максимальный (%.2f) > минимального (%.2f) в 1.2 раза: %s\n",
		maxVal, minVal, map[bool]string{true: "да", false: "нет"}[a])
	fmt.Printf("Минимальный (%.2f) < максимального (%.2f) в 1.2 раза: %s\n",
		minVal, maxVal, map[bool]string{true: "да", false: "нет"}[b])
}

func ch11ex131() {
	fmt.Printf("Ex131. Проверить, превышает ли масса самого тяжелого массу самого легкого более чем в 2 раза.\n")
	weights := []float64{50, 110, 65, 120, 70, 105, 55, 130}
	if len(weights) == 0 {
		fmt.Println("Массив пуст")
		return
	}
	minWeight, maxWeight := weights[0], weights[0]
	for _, w := range weights {
		if w < minWeight {
			minWeight = w
		}
		if w > maxWeight {
			maxWeight = w
		}
	}
	condition := maxWeight > minWeight*2
	fmt.Printf("Самый тяжелый: %.1f кг, самый легкий: %.1f кг. Превышение >2 раз: %s\n",
		maxWeight, minWeight, map[bool]string{true: "да", false: "нет"}[condition])
}

func ch11ex132() {
	fmt.Printf("Ex132. Определить скорость двух самых быстрых автомобилей.\n")
	speeds := []int{180, 200, 190, 210, 195, 205, 185, 195}
	if len(speeds) < 2 {
		fmt.Println("Недостаточно данных")
		return
	}
	first, second := 0, 0
	for _, s := range speeds {
		if s > first {
			second = first
			first = s
		} else if s > second && s != first {
			second = s
		}
	}
	fmt.Printf("Две самые высокие скорости: %d км/ч и %d км/ч\n", first, second)
}

func ch11ex133() {
	fmt.Printf("Ex133. Найти второй по минимальности элемент.\n")
	arr := []int{5, 8, 3, 8, 2, 8, 3, 2, 8, 1}
	if len(arr) < 2 {
		fmt.Println("Недостаточно элементов")
		return
	}
	firstMin, secondMin := arr[0], arr[1]
	if firstMin > secondMin {
		firstMin, secondMin = secondMin, firstMin
	}
	for i := 2; i < len(arr); i++ {
		v := arr[i]
		if v < firstMin {
			secondMin = firstMin
			firstMin = v
		} else if v < secondMin && v != firstMin {
			secondMin = v
		}
	}
	fmt.Printf("Второй минимальный элемент: %d\n", secondMin)
}
