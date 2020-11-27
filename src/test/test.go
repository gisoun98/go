package test

import "fmt"

func Aaa() { //외부 참조를 위해 함수의 첫 글자는 대문자로 한다.
	fmt.Println("A")
}

func Bbb() {
	fmt.Println("B")
}

func Alphabet() {
	Aaa()
	Bbb()
}
