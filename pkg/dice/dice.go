package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Die struct {
	Sides int
	Top   int
}

func (die *Die) Roll() {
	roll := rand.Intn(die.Sides-1) + 1
	die.Top = roll
}

func (die *Die) Set(n int) error {
	if n > die.Sides {
		return fmt.Errorf("Die cannot be set to a number greater than it's number of sides")
	} else if n < 1 {
		return fmt.Errorf("Die cannot be set to a number less than one")
	} else {
		die.Top = n
		return nil
	}

}

func CreatePool(size int, sides int) *Pool {
	// Create a pool of {size} dice, each with {sides}, sides. Returns a pointer to that Pool
	dice := make([]*Die, size)
	for d := 0; d < size; d++ {
		dice[d] = &Die{Sides: sides, Top: 1}
	}
	return &Pool{Dice: dice, Sides: sides, Description: ""}
}

func ParseDiceString(pool_string string) (*Pool, error) {
	// Reads a string of XdY format and returns a pool with those parameters.
	// Returns an error if string is in the wrong format.
	var pool *Pool
	var err error
	pool_int := make([]int, 2)

	// Check if the string contains a d. If not return an error that the correct format was not used.
	if strings.Contains(pool_string, "d") {
		pool_str := strings.Split(pool_string, "d")

		// If the string does not split into two then there was something wrong with the formatting
		if len(pool_str) != 2 {
			err = fmt.Errorf("dice pools need to be in the format XdY")
			return pool, err
		}

		// Convert the two sides of the pool into ints. If they cannot be converted return an error
		for n, d := range pool_str {
			pool_int[n], err = strconv.Atoi(d)
			if err != nil {
				return pool, err
			}
		}
	} else {
		err = fmt.Errorf("dice pools need to be in the format XdY")
		return pool, err
	}
	pool = CreatePool(pool_int[0], pool_int[1])
	return pool, err
}

type Pool struct {
	Dice        []*Die
	Sides       int
	Description string
}

func (pool *Pool) Roll() {
	// Roll each die in the pool
	for _, die := range pool.Dice {
		die.Roll()
	}
}

func (pool *Pool) List() []int {
	// Return a slice of integers of the top of each die in the pool
	list := make([]int, len(pool.Dice))
	for n, die := range pool.Dice {
		list[n] = die.Top
	}
	return list
}

func (pool *Pool) Total() int {
	// Return the total sum of the top of each dice in the pool
	total := 0
	for _, die := range pool.Dice {
		total += die.Top
	}
	return total
}

func (pool *Pool) Add() {
	// Add a die to the pool
	pool.Dice = append(pool.Dice, &Die{Sides: pool.Sides, Top: 1})
}

func (pool *Pool) Subtract() error {
	// subtract a die from the pool. If there are no dice left in the pool return an error
	var err error
	if len(pool.Dice) < 1 {
		err = fmt.Errorf("cannot subtract from a pool with no dice")
		return err
	}
	pool.Dice = pool.Dice[:len(pool.Dice)-1]
	return err
}

func (pool *Pool) Describe() string {
	// Return a human readable description of the dice in the string
	list_dice := pool.List()
	first_dice := list_dice[0 : len(list_dice)-1]

	var first_str string
	for _, i := range first_dice {
		s := fmt.Sprintf("%d, ", i)
		first_str += s
	}

	last_die := list_dice[len(list_dice)-1]
	desc := fmt.Sprintf("A pool of %d d%ds. The dice are facing %sand %d.", len(pool.Dice), pool.Sides, first_str, last_die)

	// If the pool has a description add it onto the end of the normal description
	if pool.Description != "" {
		desc = desc + " " + pool.Description
	}
	return desc
}

func CreateTable(pool_list []*Pool, names []string) (Table, error) {
	var err error
	pools := make(map[string]*Pool)
	if len(pool_list) != len(names) {
		err = fmt.Errorf("the number of pools and the number of names do not match up")
		return Table{Pools: pools, Name: ""}, err
	} else {
		for d := 0; d < len(pool_list); d++ {
			pools[names[d]] = pool_list[d]
		}
		return Table{Pools: pools, Name: ""}, err
	}
}

func ParseTableString(pool_list []string, names []string) (Table, error) {
	// Parse through the command line arguments given and return a Table with those dice in it
	// If dice are not in the right format return an error
	var err error
	var table Table
	var pools []*Pool

	for _, p := range pool_list {
		pool, err := ParseDiceString(p)
		if err != nil {
			return table, err
		}

		// Create a new pool using the parameters and append it to the existing pools
		pools = append(pools, pool)
	}

	table, err = CreateTable(pools, names)
	return table, err
}

type Table struct {
	Pools map[string]*Pool
	Name  string
}

func (table *Table) Roll() {
	// Roll all dice on the table
	for _, pool := range table.Pools {
		pool.Roll()
	}
}

func (table *Table) Clear() {
	// Delete each pool in the table
	for name := range table.Pools {
		delete(table.Pools, name)
	}
}

func (table *Table) Remove(name string) error {
	// Delete the pool of the name provided. If the name doesn't exist in the table return an error
	var err error
	if _, ok := table.Pools[name]; ok {
		delete(table.Pools, name)
	} else {
		err = fmt.Errorf("%s is not the name of a pool in this table", name)
	}
	return err
}

func (table *Table) Add(name string, size int, sides int) {
	// Add a pool to a table
	table.Pools[name] = CreatePool(size, sides)
}
