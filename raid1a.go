package piscine

import "github.com/01-edu/z01"

	//func DrawBase(x int) {
	//for i := 0; i < x; i++ {
	//	if i == 0 || i == x-1 {
	//		z01.PrintRune(111)
	//	} else if i < x-1 {
	//		z01.PrintRune(45)
	//	}
//	}
//}

func Raid1a(x, y int) {
	side := 0
	if x > 1 {
		side = 2
	} else {
		side = 1
	}
	//DrawBase(x)
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune(111)
		} else if i < x-1 {
			z01.PrintRune(45)
		}
	}
	z01.PrintRune(rune(10))
	if y > 1 {
		for l := 0; l < y-2; l++ {
			for j := 0; j < side; j++ {
				z01.PrintRune(124)
				for k := 0; k < x-2; k++ {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune(rune(10))
		}
		//DrawBase(x)
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune(111)
			} else if i < x-1 {
				z01.PrintRune(45)
			}
		}
		z01.PrintRune(rune(10))
	}
}
