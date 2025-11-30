package main

import (
	chapters "1400-CirclesOfHell/chapters"
	"fmt"
)

func main() {
	chaptersList := []func(){
		chapters.Run01, chapters.Run02, chapters.Run03, chapters.Run04, chapters.Run05,
		chapters.Run06,
	}

	for {
		fmt.Print("\n Выберите главу (1-14): ")
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}
		if n >= 1 && n <= 14 {
			chaptersList[n-1]()
		} else {
			fmt.Println("От 1 до 14! Для выхода введите 0.")
		}
	}
}
