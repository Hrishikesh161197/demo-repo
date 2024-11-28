package main

import "fmt"

func main(){
			var num int
			fmt.print("Enter the number :")
			fmt.scanln(&num)

			if num%3 == 0 && num%5 == 0{
					fmt.println("Fizzbuzz")
				 }else if num%3 == 0{
					fmt.println("Fizz")
				 }else if num%5 == 0{
					fmt.println("Buzz")
				 }else{
					fmt.println(num)
				 }	 
		}