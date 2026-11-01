package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			if j == 0 {
				// верхняя строка
				if i == 0 {
					z01.PrintRune('/')
				} else if i == x-1 {
					z01.PrintRune('\\')
				} else {
					z01.PrintRune('*')
				}
			} else if j == y-1 {
				// нижняя строка
				if i == 0 {
					z01.PrintRune('\\')
				} else if i == x-1 {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			} else {
				// средние строки
				if i == 0 || i == x-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}