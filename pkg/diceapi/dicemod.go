package diceapi

import (
	"time"
)

type Table struct {
	ID          int       `json:"ID"`
	Name        string    `json:"Name"`
	Owner       int       `json:"Owner"`
	Description string    `json:"Description"`
	DateCreated time.Time `json:"DateCreated"`
}

type Pool struct {
	ID          int       `json:"ID"`
	TableID     int       `json:"TableID"`
	Description string    `json:"Description"`
	DiceSize    int       `json:"DiceSize"`
	Dice        []int     `json:"Dice"`
	DateCreated time.Time `json:"DateCreated"`
}

type Character struct {
	ID          int       `json:"ID"`
	Name        string    `json:"Name"`
	UserID      int       `json:"UserID"`
	TableID     int       `json:"TableID"`
	DateCreated time.Time `json:"DateCreated"`
}

type User struct {
	ID          int       `json:"ID"`
	Username    string    `json:"Username"`
	Password    string    `json:"Password"`
	DateCreated time.Time `json:"DateCreated"`
}

type Roll struct {
	ID           int       `json:"ID"`
	TableID      int       `json:"TableID"`
	UserID       int       `json:"UserID"`
	CharacterID  int       `json:"CharacterID"`
	PoolID       int       `json:"PoolID"`
	TypeOfChange string    `json:"TypeOfChange"`
	DiceBefore   []int     `json:"DiceBefore"`
	DiceAfter    []int     `json:"DiceAfter"`
	Date         time.Time `json:"Date"`
}
