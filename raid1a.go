package piscine

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune(111)
		} else if i < x-1 {
			z01.PrintRune(45)
		}
	}
	z01.PrintRune(rune(10))

	for j := 0; j < y-1; j++ {
		z01.PrintRune(124)
		if j == 0 {
			for k := 0; k < x-2; k++ {
				z01.PrintRune(' ')
			}
		}

	}
	z01.PrintRune(rune(10))
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune(111)
		} else if i < x-1 {
			z01.PrintRune(45)
		}
	}
	z01.PrintRune(rune(10))
}
