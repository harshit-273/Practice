package main

import (
	"fmt"
)

/************************************************************************

Pattern - 1 

*****
*****
*****
*****
*****

Pattern - 2

*
**
***
****
*****

Pattern - 3

1
12
123
1234
12345

Pattern - 4

1
22
333
4444
55555

Pattern - 5

*****
****
***
**
*

Pattern - 6

12345
1234
123
12
1

Pattern - 7

    *
   ***
  *****
 *******
*********

Pattern - 8

*********
 *******
  *****
   ***
    *

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

************************************************************************/

func main() {
	//Pattern - 1
	/*
	var stars int = 5
	
	for i:=0; i<stars; i++ {
		for j:=0; j<stars; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	*/
	
	//Pattern - 2
	/*
	var stars int = 5
	
	for i:=0; i<stars; i++ {
		for j:=0; j<=i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	*/
	
	// Pattern - 3
	/*
	var nums int = 5
	
	for i:=1; i<=nums; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
	*/
	
	//Pattern - 4
	/*
	var nums int = 5
	
	for i:=1; i<=nums; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
	*/
	
	//Pattern - 5
	/*
	var stars int = 5
	
	for i:=0; i<stars; i++ {
		for j:=stars; j>i; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	*/
	
	//Pattern - 6
	/*
	var nums int = 5
	
	for i:=nums; i>0; i-- {
		for j:=1; j<=i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
	*/
	
	//Pattern - 7
	/*
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
	
	//Pattern - 8
	/*
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
	
	// Pattern - 9
	/*
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
	
	//Pattern - 10
	/*
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
}