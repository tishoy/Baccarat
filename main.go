package main

import (
"fmt"
"time"
"math/rand"
)

// var slc []byte

func main() {
	cardList := []int{1}
	for i := 1; i <= 52; i++ {
		cardList = append(cardList, i)
		// fmt.Print(cardList[i])
	}

	
	for i := 0; i < 4; i++ {
		if i == 0 {
			fmt.Print("zhuangjia")
		} else if i == 2 {
			fmt.Print("xianjia")
		}
		rand.Seed(time.Now().UnixNano())
		card := rand.Intn(54 - i)
    	cardList = append(cardList[:card],cardList[card+1:]...)
		outPutCard(3 + card % 4, 1 + card / 4)
	}

	

	
	// getCard()
	// getCard()

	// outPutCard()
}

func outPutCard(color int, num int) {
	fmt.Printf("%c,", color)
	switch num {
		case 1:
			fmt.Print("A")
		case 11:
			fmt.Print("J")
		case 12:
			fmt.Print("Q")
		case 13:
			fmt.Print("K")
		default:
			fmt.Print(num)
	}
}