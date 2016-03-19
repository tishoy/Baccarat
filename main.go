package main

import (
"fmt"
"time"
"math/rand"
)

var resultCard []int

func main() {
	cardList := []int{}
	for i := 1; i <= 52; i++ {
		cardList = append(cardList, i)
		// fmt.Print(cardList[i])
	}
	
	for i := 0; i < 4; i++ {
		if i == 0 {
			fmt.Print("zhuang")
		} else if i == 2 {
			fmt.Print("xian")
		}
		rand.Seed(time.Now().UnixNano())
		card := rand.Intn(52 - i)
    	cardList = append(cardList[:card],cardList[card+1:]...)
		outPutCard(3 + card % 4, 1 + card / 4)
	}

	checkWin()
}

func outPutCard(color int, num int) {
	fmt.Printf("%c", color)
	switch num {
		case 1:
			fmt.Print("A")
			resultCard = append(resultCard, 1)
		case 11:
			fmt.Print("J")
			resultCard = append(resultCard, 1)
		case 12:
			fmt.Print("Q")
			resultCard = append(resultCard, 1)
		case 13:
			fmt.Print("K")
			resultCard = append(resultCard, 1)
		default:
			fmt.Print(num)
			resultCard = append(resultCard, num)
	}
}

func checkWin() {
	zhuangResult := (resultCard[0] + resultCard[1]) % 10
	xianResult := (resultCard[2] + resultCard[3]) % 10
	if zhuangResult > xianResult {
		fmt.Print("\n\rzhuang Win")
	} else if zhuangResult < xianResult {
		fmt.Print("\n\rxian Win")
	} else if zhuangResult == xianResult {
		fmt.Print("\n\rhe Win")
	}
	
}
