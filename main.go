package main

import (
	"fmt"
)


func Server(wup chan<- State, pup <-chan PlayerUpdate) {
	s := State{
		[]Player{Player{"Glenda", Ent{EntPlayer, Location{33, 33}, 0, DefaultHealth}}},
		[]Ent{Ent{EntCreature, Location{22, 22}, 90, 666}}}

	for up := range pup {
		if up.quit {
			return
		}

		s.players[0].dir = up.fireDir
		wup <- s
	}
}




func Client(wup <-chan State, in chan<- PlayerUpdate) {
	go Display(wup)
	Input(in)
}

func Input(in chan<- PlayerUpdate) {
	pu := PlayerUpdate{}
	for i := 0; i < 3; i++ {
		pu.fireDir = Direction(i * 10)
		in <- pu
	}
	pu.quit = true
	in <- pu
}

func Display(wup <-chan State) {
	for w := range wup {
		fmt.Printf("%v\n", w)
	}
}

func main() {
	wup := make(chan State)
	pup := make(chan PlayerUpdate)
	go Client(wup, pup)
	Server(wup, pup)
}
