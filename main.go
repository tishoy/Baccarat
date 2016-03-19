package main

import (
"fmt"
"time"
"math/rand"
"strconv"
)

type card struct {
	point string
	color, value int
}

var resultCard []*card

func main() {
	cardList := []*card{}
	for i := 1; i <= 52 * 8; i++ {
		card := generateCard(i)
		cardList = append(cardList, card)
	}
	
	for i := 0; i < 4; i++ {
		if i == 0 {
			fmt.Print("Bank get card")
		} else if i == 2 {
			fmt.Print(" ===== Play get card")
		}
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(52 * 8 - i)
		fmt.Printf("%c", cardList[index].color)
		fmt.Print(cardList[index].point)
    	cardList = append(cardList[:index], cardList[index+1:]...)
    	resultCard = append(resultCard, cardList[index])
	}

	checkWin()
}

func generateCard(i int) (*card) {
	num := i / 32 + 1
	color := i % 4 + 3
	value := 0
	point := ""
	switch num {
		case 1:
			value = 1
			point = "A"
		case 11:
			value = 1
			point = "J"
		case 12:
			value = 1
			point = "Q"
		case 13:
			value = 1
			point = "K"
		default:
			value = num
			point = strconv.Itoa(num)
	}
	return &card{point: point, color: color, value: value}
}

func checkWin() {
	if (resultCard[0].point == resultCard[1].point || resultCard[2].point == resultCard[3].point) {
		fmt.Print("\n\rPair Win")
		return
	}
	zhuangResult := (resultCard[0].value + resultCard[1].value) % 10
	xianResult := (resultCard[2].value + resultCard[3].value) % 10
	if zhuangResult > xianResult {
		fmt.Print("\n\rBank Win")
	} else if zhuangResult < xianResult {
		fmt.Print("\n\rPlay Win")
	} else if zhuangResult == xianResult {
		fmt.Print("\n\rhe Win")
	}
	
}
