package main

import "fmt"

/* Patterns */

func main() {
	/*
		Pattern - 1

		*****
		*****
		*****
		*****
		*****

		var stars int = 5

		for i:=0; i<stars; i++ {
			for j:=0; j<stars; j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 2

		*
		**
		***
		****
		*****

		var stars int = 5

		for i:=0; i<stars; i++ {
			for j:=0; j<=i; j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 3

		1
		12
		123
		1234
		12345

		var nums int = 5

		for i:=1; i<=nums; i++ {
			for j:=1; j<=i; j++ {
				fmt.Printf("%d", j)
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 4

		1
		22
		333
		4444
		55555

		var nums int = 5

		for i:=1; i<=nums; i++ {
			for j:=1; j<=i; j++ {
				fmt.Printf("%d", i)
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 5

		*****
		****
		***
		**
		*

		var stars int = 5

		for i:=0; i<stars; i++ {
			for j:=stars; j>i; j-- {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 6

		12345
		1234
		123
		12
		1

		var nums int = 5

		for i:=nums; i>0; i-- {
			for j:=1; j<=i; j++ {
				fmt.Printf("%d", j)
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 7

		    *
		   ***
		  *****
		 *******
		*********

		var stars int = 5

		for i:=1; i<=stars; i++ {
			for k:=stars; k>i; k-- {
				fmt.Printf(" ")
			}
			for j:=1; j<=((2*i)-1); j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	*/

	/*

		Pattern - 8

		*********
		 *******
		  *****
		   ***
		    *

		var stars int = 5

		for i:=stars; i>0; i-- {
			for k:=stars; k>i; k-- {
				fmt.Printf(" ")
			}
			for j:=1; j<=((2*i)-1); j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 9

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

		var stars int = 5

		for i:=1; i<=stars; i++ {
			for k:=stars; k>i; k-- {
				fmt.Printf(" ")
			}
			for j:=1; j<=((2*i)-1); j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
		for i:=stars; i>0; i-- {
			for k:=stars; k>i; k-- {
				fmt.Printf(" ")
			}
			for j:=1; j<=((2*i)-1); j++ {
				fmt.Printf("*")
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 10

		*
		**
		***
		****
		*****
		****
		***
		**
		*

		var stars int = 5

		for i:=1; i<(2*stars); i++ {
			if i<=5 {
				for j:=1; j<=i; j++ {
					fmt.Printf("*")
				}
			} else {
				for j:=1; j<=((2*stars)-i); j++ {
					fmt.Printf("*")
				}
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 11

		1
		01
		101
		0101
		10101

		var zero1 int = 5
		for i:=1; i<=zero1; i++ {
			for j:=1; j<=i; j++ {
				if (i+j)%2 == 0 {
					fmt.Printf("1")
				} else {
					fmt.Printf("0")
				}
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 12

		1        1
		12      21
		123    321
		1234  4321
		1234554321

		var nums int = 5

		for i:=1; i<= nums; i++ {
			for j:=1; j<=i; j++ {
				fmt.Printf("%d", j)
			}
			for spaces:=1; spaces<=(2*(nums-i)); spaces++ {
				fmt.Printf(" ")
			}
			for k:=i; k>=1; k-- {
				fmt.Printf("%d", k)
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 13

		01
		02 03
		04 05 06
		07 08 09 10
		11 12 13 14 15

		var nums int = 5
		var count int = 1

		for i:=1; i<=nums; i++ {
			for j:=1; j<=i; j++ {
				fmt.Printf("%02d ", count)
				count++
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 14

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

		var stars int = 5

		for i:=1; i<=stars; i++ {
			for j:=1; j<=(2*stars); j++ {
				if j<=(stars-(i-1)) || j>(stars+(i-1)) {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
		for i:=stars; i>0; i-- {
			for j:=1; j<=(2*stars); j++ {
				if j<=(stars-(i-1)) || j>(stars+(i-1)){
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 15

		*        *
		**      **
		***    ***
		****  ****
		**********
		****  ****
		***    ***
		**      **
		*        *

		var stars int = 5

		for i:=stars; i>0; i-- {
			for j:=1; j<=(2*stars); j++ {
				if j<=(stars-(i-1)) || j>(stars+(i-1)){
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
		for i:=2; i<=stars; i++ {
			for j:=1; j<=(2*stars); j++ {
				if j<=(stars-(i-1)) || j>(stars+(i-1)) {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
	*/

	/*
		Pattern - 16

		*****
		*   *
		*   *
		*   *
		*****
	*/

	var stars int = 5

	for i := 1; i <= stars; i++ {
		for j := 1; j <= stars; j++ {
			if i == 1 || j == 1 || i == stars || j == stars {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}

}
