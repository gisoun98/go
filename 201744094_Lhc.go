//201744094 C반 이해찬
package main

//사용되는 패키지
import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
)

//절댓값
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

//팩토리얼
func Factorial(n int) int {
	//0의 팩토리얼은 1이다.
	if n == 0 {
		return 1
	}
	//해당 함수를 최종 리턴 값이 1이 될 때 까지 재귀호출한다.
	return n * Factorial(n - 1)
}

//피보나치수열
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

//거듭제곱
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

func main() {
	for {
		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5) 종료 : ")
		//사용할 번호 입력
		var in int = Input()

		switch in {
			//1 입력 시 절대값 계산 함수 호출 및 출력
			case 1 :
				fmt.Print("정수 입력(절대값 계산) : ")
				fmt.Println(Absolute(Input()))
			//2 입력 시 팩토리얼 계산 함수 호출 및 출력
			case 2 :
				fmt.Print("정수 입력(팩토리얼 계산) : ")
				fmt.Println(Factorial(Input()))
			//3 입력 시 피보나치 수열 계산 함수 호출 및 출력
			case 3 :
				fmt.Print("정수 입력(피보나치 계산) : ")
				fmt.Println(Fibonacci(Input()))
			//4 입력 시 거듭 제곱 계산 함수 호출 및 출력
			//Base : 밑,	Exponent : 지수
			case 4 : 
				var base, exp int
				fmt.Print("Base 입력 : ")
				base = Input()
				fmt.Print("Exponent 입력 : ")
				exp = Input()
				fmt.Println(Power(base, exp))
			//5 입력 시 프로그램 종료
			case 5 :
				os.Exit(3)
			default :
				fmt.Println("1 ~ 5 까지의 정수만을 입력 해주십시오.")
		}
	}
}
