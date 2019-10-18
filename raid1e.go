package piscine

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 {
					if j == 0 {
						z01.PrintRune(65)
					} else if j == x-1 {
						z01.PrintRune(67)
					} else {
						z01.PrintRune(66)
					}
				} else if i == y-1 {
					if j == 0 {
						z01.PrintRune(67)
					} else if j == x-1 {
						z01.PrintRune(65)
					} else {
						z01.PrintRune(66)
					}
				} else {
					if j == 0 || j == x-1 {
						z01.PrintRune(66)
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune(rune(10))
		}
	}
}
