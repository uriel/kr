package main

import "time"


type EntKind uint8

const (
	EntCreature = iota
	EntPlayer
	EntItem
)

const (
	DefaultHealth = 99
)

type Location struct {
	x, y uint
}

type Direction uint

type Ent struct {
	kind   EntKind
	loc    Location
	dir    Direction
	status uint // Can be used for several things, hp, ttl (for drops), ...
}

type Health uint16

type Player struct {
	name string
	Ent
	// TODO perks, weapons, etc
}

type PlayerUpdate struct {
	move      bool
	moveDir   Direction
	fire      bool
	fireDir   Direction
	reloading bool
	quit      bool
}

type State struct {
	players []Player
	ents    []Ent
}

var World = struct {
	state State
	time  time.Time
}{}

