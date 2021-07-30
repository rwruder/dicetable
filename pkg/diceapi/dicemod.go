package diceapi

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
	ID          int       `json:"ID"`
	Table       int       `json:"Table"`
	Description string    `json:"Description"`
	DiceSize    int       `json:"DiceSize"`
	Dice        []int     `json:"Dice"`
	DateCreated time.Time `json:"DateCreated"`
}

type Character struct {
	ID          int       `json:"ID"`
	User        int       `json:"User"`
	Table       int       `json:"Table"`
	DateCreated time.Time `json:"DateCreated"`
}

type User struct {
	ID          int       `json:"ID"`
	Username    string    `json:"Username"`
	Password    string    `json:"Password"`
	DateCreated time.Time `json:"DateCreated"`
}

type Roll struct {
	ID           int    `json:"ID"`
	Table        int    `json:"Table"`
	User         int    `json:"User"`
	Character    int    `json:"Character"`
	Pool         int    `json:"Pool"`
	TypeOfChange string `json:"TypeOfChange"`
	DiceBefore   []int  `json:"DiceBefore"`
	DiceAfter    []int  `json:"DiceAfter"`
}
