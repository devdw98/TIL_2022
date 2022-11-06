package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 시간 값을 랜덤 시드로 설정
	resultNum := rand.Intn(100)      // 0 ~ 99 까지 수를 랜덤으로 생성
	var count int8

	for {
		var num int
		fmt.Print("숫자값을 입력하세요: ")
		_, err := fmt.Scanf("%d\n", &num)

		count++
		if err != nil {
			fmt.Println("잘못된 값입니다. error : " + err.Error())
			continue
		}

		if num < resultNum {
			fmt.Println("입력하신 숫자가 더 작습니다")
		} else if num > resultNum {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else {
			break
		}
	}
	fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: %d\n", count)

}
