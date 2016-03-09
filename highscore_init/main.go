package main

import (
	"allthethings/highscore/score"
	"fmt"
)

func main() {
	score.DeleteTable()
	score.CreateTable()
	score.Seed("TRS", 1001)
	score.Seed("AAA", 999)
	fmt.Println(score.GetAll())
}
