package main

import (
	"dates"
	"fmt"
)

func main() {
	days := 5
	fmt.Println("예약일은 ", days, "일 입니다.")
	fmt.Println("추가적으로 7일이 소요될 수 있습니다. 총 ", days + dates.DaysInWeek, "일 안에는 처리가 됩니다.")
}
