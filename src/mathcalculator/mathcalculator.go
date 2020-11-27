package mathcalculator

import (
	"fmt"
	"bufio"
	"log"
	"strconv"
	"strings"
	"os"
)

const Repeat int = 20 //Root()의 for문 반복 횟수

//1.절댓값
func Absolute(n int) int {
	//입력 받은 정수가 음수라면
	if n < 0 {
		//-1을 곱해 양수로 만들어 리턴한다.
		return n * -1
	} else {
		//입력 받은 정수가 양수라면 그대로 리턴한다.
		return n
	}
}

//2.팩토리얼
func Factorial(n int) int {
	//0의 팩토리얼은 1이다.
	if n == 0 {
		return 1
	}
	//해당 함수를 최종 리턴 값이 1이 될 때 까지 재귀호출한다.
	return n * Factorial(n - 1)
}

//3.피보나치수열
func Fibonacci(n int) int {
	//a1, a2 : 피보나치 수열의 값이 되는 항.
	//result : 사용자가 원하는 위치에 있는 항의 값을 갖는 변수.
	var a1, a2, result int

	//피보나치 수열의 첫째, 둘째 항은 값이 1 이다.
	if n <= 2 {
		return 1
	}
	//a1, a2의 초기 값은 첫째, 둘째 항으로 한다.
	a1 = 1
	a2 = 1

	//피보나치 수열 공식을 반복문으로 구현
	for i := 2; i < n; i++ {
		result = a1 + a2
		a1 = a2
		a2 = result
	}
	return result
}

//4.거듭제곱
func Power(base int, exp int) int {
	//거듭제곱 결과 값을 저장하는 변수 result에 밑 값 base 삽입
	var result int = base

	//지수 exp 만큼 반복문을 돌려 result에 base 제곱값 저장
	for i := 0; i < exp; i++ {
		result *= base
	}

	//초기 result에 이미 base의 값이 한 번 들어갔었으므로 base로 한 번 나눈 뒤에 리턴
	return result / base
}

//5.사칙연산
func Operation(f1 int, f2 int, op int) int {
	var result int

	switch op {
		case 1 :
			result = f1 + f2
			fmt.Printf("%d + %d = ", f1, f2)
		case 2 :
			result = f1 - f2
			fmt.Printf("%d - %d = ", f1, f2)
		case 3 :
			result = f1 / f2
			fmt.Printf("%d / %d = ", f1, f2)
		case 4 :
			result = f1 * f2
			fmt.Printf("%d * %d = ", f1, f2)
	}
	return result
}

//6.제곱근
func Root(n int) float64 {
	var k int
	var result, tmp float64 = 0, float64(n)

	for k, result = 0, tmp; k < Repeat; k++ {
		if result < 1.0 {
			break
		}
		result = (result * result + tmp) / (2.0 * result)
	}
	return result
}

//7.순열
func Permutation(n int, r int) int {
	/*if r == 1 {
		return n
	} else {
		return n * Permutation(n - 1, r - 1)
	}*/
	return Factorial(n) / Factorial(n - r)
}

//8.소수 판정
func Prime(n int) bool {
	isPrime := false

	for i := 2; i < n; i++ {
		if n % i == 0 {
			isPrime = true
			break;
		}
	}
	return isPrime
}

//9.조합
func Combination(n int, r int) int {
	/*if n == r || r == 0 {
		return 1
	} else {
		return Combination(n - 1, r - 1) + Combination(n - 1, r)
	}*/
	return Permutation(n, r) / Factorial(r)
}

//입력한 문자열 정수형으로 변환 함수
func Input() int {
	//문자열 입력
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	//오류 검출
	if err != nil {
		log.Fatal(err)
	}
	//줄 바꿈 제거
	input = strings.TrimSpace(input)
	//Atoi 함수를 이용한 문자열 -> 정수 변환
	num, err := strconv.Atoi(input)
	//오류 검출
	if err != nil {
		log.Fatal(err)
	}
	return num
}
