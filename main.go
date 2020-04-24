package main

import (
	"fmt"
)

func add(a,b int) int {
	return a + b
}

func multiply(c,d int) int{
	return c * d
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(	n-2)
}

func primeNumbers(n int)   {
	p := 2
	i := 1
	for i <= n{
		flag := 1
		for count := 2; count<=p-1; count++{
			if p%count==0{
				flag = 0
				break
			}
		}
		if flag == 1{
			fmt.Printf("%v ", p)
			i++
		}
		p++
	}
	return
}


func main(){

	fmt.Println(add(5,6)) //addition
	fmt.Println(multiply(4,112)) //multiplication

	for n := 0; n < 3; n++ {
		result := fibonacci(n)
		fmt.Printf("%v ",result)
	} // fibonnaci, change the 3 into desired number

	primeNumbers(4) //list prime numbers
}
