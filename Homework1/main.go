package main

import (
	"fmt"
)

func main() {
	var columns int
	var lines int

	fmt.Println("Шахматная доска 8х8")
	draw(8, 8)

	fmt.Println("Сделайте свою доску :)")
	fmt.Print("Введите количество строк решетки: ")
	fmt.Scan(&lines)

	fmt.Print("Введите количество стоблцов решетки: ")
	fmt.Scan(&columns)

	draw(columns, lines)

}

func draw(c int, l int) {

	r := "#"
	s := " "

	for i := 0; i < l; i++ {
		for j := 0; j < c; j++ {
			//если строка четная и столбец нечетный или строка нечетная и столбец четный, то пробел
			if (i%2 == 0 && j%2 != 0) || (i%2 != 0 && j%2 == 0) {
				fmt.Print(s)
			} else if (i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0) { //если строка четная столбец четный или строка нечетная и столбец нечетный, то символ решетка
				fmt.Print(r)
			}
		}
		fmt.Println()
	}

}
