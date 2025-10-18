package main

import "fmt"

// For outer loop, count the no. of rows
// For inner loop, count the no. of cols & connect them somehow to the rows
// Print them inside the inner for loop
// Observe symmetry in case of some pattern
func main() {
	PrintLastPattern(5)
}

/*
Print n*n star
****
****
****
****
*/
func printBox(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

/*
Print n*n Right Triangle
*
* *
* * *
* * * *
* * * * *
*/
func printRightTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

/*
Print n*n Reverse Right Triangle
* * * * *
* * * *
* * *
* *
*
*/
func printReverseRightTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

/*
Print n*n nums in Right Triangle
1
1 2
1 2 3
1 2 3 4
1 2 3 4 5
*/
func printNumsRightTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println("")
	}
}

/*
Print n*n nums in Reverse Right Triangle
1 2 3 4 5
1 2 3 4
1 2 3
1 2
1
*/
func printNumsInReverseRightTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			fmt.Printf("%d ", n-j+1)
		}
		fmt.Println("")
	}
}

/*
Print n*n Right Triangle
1
2 2
3 3 3
4 4 4 4
5 5 5 5 5
*/
func printNumsForiInRightTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println("")
	}
}

/*
Print n*n Triangle

		    *
		   ***
		  *****
	     *******
	    *********
*/
func printIsoscelesTriangle(n int) {
	for i := 0; i < n; i++ {
		// n-i-1 spaces
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		// (2*i)+1 stars
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		// n-i-1 spaces
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

/*
Print n*n Triangle

	*********
	 *******
	  *****
	   ***
	    *
*/
func printReverseIsoscelesTriangle(n int) {
	for i := n; i >= 0; i-- {
		// n-i-1 spaces
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		// (2*i)+1 stars
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		// n-i-1 spaces
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

/*
Print following like diamond

		    *
		   ***
		  *****
	     *******
	    *********
		*********
		 *******
		  *****
		   ***
		    *
*/
func printDiamond(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	for i := n; i > 0; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

/*
Print following like horizontal triangle

	*
	**
	***
	****
	*****
	****
	***
	**
	*
*/
func printHorizontalTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

/*
Print 1 0 pattern
1
0 1
1 0 1
0 1 0 1
1 0 1 0 1
*/
func printZeroOnePattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", (i+j+1)%2)
		}
		fmt.Println("")
	}
}

/*
Print nums pattern
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
*/
func printNumsPattern(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println("")
	}
}

/*
Print Valley of numbers like following
1        1
12      21
123    321
1234  4321
1234554321
*/
func printValleyOfNums(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		for j := 1; j <= 2*(n-i); j++ {
			fmt.Print(" ")
		}

		for j := i; j > 0; j-- {
			fmt.Print(j)
		}

		fmt.Println()
	}
}

/*
A
A B
A B C
A B C D
A B C D E
*/
func printCapitalAlphabetsRightTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%c ", 65+j-1)
		}
		fmt.Println("")
	}
}

/*
A B C D E
A B C D
A B C
A B
A
*/
func printCapitalAlphabetsInReverseRightTriangle(n int) {
	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%c ", 65+j-1)
		}
		fmt.Println("")
	}
}

/*
E E E E E
D D D D
C C C
B B
A
*/
func printCapitalAlphaRowsInReverseRightTriangle(n int) {
	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%c ", 65+i-1)
		}
		fmt.Println("")
	}
}

/*
A
B B
C C C
D D D D
E E E E E
*/
func printCapitalAlphaRowsInRightTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%c ", 65+i-1)
		}
		fmt.Println("")
	}
}

/*
Print the following alphabets in triangle!

		   A
		  ABA
		 ABCBA
		ABCDCBA
	   ABCDEDCBA
*/
func PrintAlphaTriangle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (2*i+1)/2; j++ {
			fmt.Printf("%c", 65+j)
		}
		for j := (2*i + 1) / 2; j >= 0; j-- {
			fmt.Printf("%c", 65+j)
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

/*
Print the following alphabets in triangle!
E
DE
CDE
BCDE
ABCDE
*/
func PrintRightAlphaTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			fmt.Printf("%c", 65+n-j)
		}
		fmt.Println("")
	}
}

/*
Print the following diamond pattern
**********
****  ****
***    ***
**      **
*        *
*        *
**      **
***    ***
****  ****
**********
*/
func PrintEmptyDiamond(n int) {
	for i := 0; i < n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		if i != n-1 {
			fmt.Println("")
		}
	}
	for i := n; i >= 0; i-- {
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

/*
Print the following diamond pattern
*        *
**      **
***    ***
****  ****
**********
****  ****
***    ***
**      **
*        *
*/
func PrintReverseEmptyDiamond(n int) {
	for i := n; i >= 0; i-- {
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for i := 1; i < n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		for j := 0; j < 2*i; j++ {
			fmt.Print(" ")
		}
		for j := n - i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

/*
Print square
*****
*   *
*   *
*   *
*****
*/
func PrintSquare(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

/*
4444444
4333334
4322234
4321234
4322234
4333334
4444444
*/
func PrintLastPattern(n int) {
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			top := i
			left := j
			right := 2*n - 2 - j
			bottom := 2*n - 2 - i
			fmt.Printf("%d", n-min(min(top, bottom), min(left, right)))
		}
		fmt.Println("")
	}
}
