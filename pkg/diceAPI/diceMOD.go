package diceAPI

import (
	"time"
)

type Table struct {
	ID          int       `json:"ID"`
	Name        string    `json:"Name"`
	Creator     int       `json:"Creator"`
	Description string    `json:"Description"`
	DateCreated time.Time `json:"DateCreated"`
}

type Pool struct {
	ID          int    `json:"ID"`
	Table       int    `json:"Table"`
	Description string `json:"Description"`
	DiceSize    int
	Dice        []int
	DateCreated time.Time
}

type Character struct {
	ID          int
	User        int
	Table       int
	DateCreated time.Time
}

type User struct {
	ID          int
	Username    string
	Password    string
	DateCreated time.Time
}

type Roll struct {
	ID           int
	Table        int
	User         int
	Character    int
	Pool         int
	TypeOfChange string
	DiceBefore   []int
	DiceAfter    []int
}
