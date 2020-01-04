package main

import "fmt"

func main() {
	score1 := 10
	score2 := 10
	score3 := 10

	promoters :=0
	detractors :=0

	if score1 >=9 {
		promoters := promoters+1
		if score1<=6 {
			detractors=detractors+1
		}
		if score2 >=9 {
			promoters := promoters+1
			if score2<=6 {
				detractors=detractors+1
			}
			if score3 >=9 {
				promoters := promoters+1
				if score3<=6 {
					detractors=detractors+1
				}
			}
			//2/3*100 -> 0*100 -> 0
			//2*100/3 -> 2*100/3 -> 66.67
			//nps:=(promoters-detractors) /3*100
			nps:=(promoters-detractors) *100/3
			fmt.Println(nps)
}