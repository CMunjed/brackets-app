package main

import (
	"fmt"
	"math"

	"example.com/api/app"
)

type Member struct {
	name string
}

type Match struct {
	member1     Member
	member2     Member
	member1Wins uint
	member2Wins uint
}

func main() {
	fmt.Print("Type a number of rounds: ")
	var rounds float64

	fmt.Print("Type a number of rounds: ")
	fmt.Scan(&rounds)
	rounds = math.Pow(2, math.Ceil(math.Log(rounds)/math.Log(2)))

	var r int = int(rounds)
	var matches = make([]Match, r)

	UNINITIALIZED_MEMBER := Member{name: "UNINITIALIZED"}
	UNINITIALIZED_MATCH := Match{member1: UNINITIALIZED_MEMBER, member2: UNINITIALIZED_MEMBER, member1Wins: 0, member2Wins: 0}

	for i := 0; i < r; i++ {
		matches = append(matches, UNINITIALIZED_MATCH)
	}

	for i := 0; i < r; i++ {
		fmt.Println(matches[i].member1.name + " " + matches[i].member2.name)
	}

	fmt.Println("Hello!")
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
