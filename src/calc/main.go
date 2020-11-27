package main

import (
	"mathcalculator"
	"fmt"
	"os"
)

func main() {
	for {	//무한 루프
		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5) 사칙연산\n")
		fmt.Print("	      6) 제곱근 7) 순열 8)소수판정 9)조합 10)종료 : ")
		//사용할 번호 입력
		var in int = mathcalculator.Input()

		switch in {
			//1 입력 시 절대값 계산 함수 호출 및 출력
			case 1 :
				fmt.Print("정수 입력(절대값 계산) : ")
				fmt.Println(mathcalculator.Absolute(mathcalculator.Input()))
			//2 입력 시 팩토리얼 계산 함수 호출 및 출력
			case 2 :
				fmt.Print("정수 입력(팩토리얼 계산) : ")
				fmt.Println(mathcalculator.Factorial(mathcalculator.Input()))
			//3 입력 시 피보나치 수열 계산 함수 호출 및 출력
			case 3 : 
				fmt.Print("정수 입력(피보나치 계산) : ")
				fmt.Println(mathcalculator.Fibonacci(mathcalculator.Input()))
			//4 입력 시 거듭 제곱 계산 함수 호출 및 출력
			//Base : 밑,    Exponent : 지수
			case 4 :
				var base, exp int
				fmt.Print("Base 입력 : ")
				base = mathcalculator.Input()
				fmt.Print("Exponent 입력 : ")
				exp = mathcalculator.Input()
				fmt.Println(mathcalculator.Power(base, exp))
			//5 입력 시 사칙 연산 
			case 5 :
				var f1, f2, op int
				fmt.Print("1) [+] 2) [-] 3) [/] 4) [*] : ")
				op = mathcalculator.Input()
				fmt.Print("정수 입력(f1) = ")
				f1 = mathcalculator.Input()
				fmt.Print("정수 입력(f2) = ")
				f2 = mathcalculator.Input()
				fmt.Println(mathcalculator.Operation(f1, f2, op))
			//6 입력 시 제곱근 계산
			case 6 : 
				fmt.Print("정수 입력(제곱근 계산) : ")
				fmt.Println(mathcalculator.Root(mathcalculator.Input()))
			//7 입력 시 순열 계산
			case 7 :
				fmt.Print("정수 입력(n) : ")
				n := mathcalculator.Input()
				for {
					fmt.Print("정수 입력(r, 0 <= r <= n) : ")
					r := mathcalculator.Input()
					if 0 > r || r > n {
						fmt.Println("0 이상의 수를 혹은 n 이하의 수만 입력해주세요.")
						continue
					}
					fmt.Println(mathcalculator.Permutation(n, r))
					break
				}
			//8 입력 시 입력 된 수가 소수 인지 판별
			case 8 :
				fmt.Print("정수 입력(소수 판별) : ")
				n := mathcalculator.Input()
				if mathcalculator.Prime(n) {
					fmt.Println(n, " is not prime number")
				} else {
					fmt.Println(n, " is prime number")
				}
			//9 입력 시 조합 계산
			case 9 :
				fmt.Print("정수 입력(n) : ")
				n := mathcalculator.Input()
				for {
					fmt.Print("정수 입력(r, 0 <= r <= n) : ")
					r := mathcalculator.Input()
					if 0 > r || r > n {
						fmt.Println("0 이상의 수를 혹은 n 이하의 수만 입력해주세요.")
						continue
					}
					fmt.Println(mathcalculator.Combination(n, r))
					break
				}
			//10 입력 시 프로그램 종료
			case 10 :
				os.Exit(3)
			default :
				fmt.Println("1 ~ 10 까지의 정수만을 입력 해주십시오.")
		}
	}
}
