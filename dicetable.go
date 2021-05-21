package main

import (
	"dicetable/dice"
	"dicetable/tablecommands"
	"flag"
	"fmt"
	"strings"
)

func main() {
	interactivePtr := flag.Bool("i", false, "Start interactive table prompt")
	namesPtr := flag.String("names", "", "Names for the dice pools entered. Seperate each by a coma with no space")
	tablenamePtr := flag.String("tablename", "", "Names the table. Changes the prompt.")
	flag.Parse()
	dice_args := flag.Args()
	var names []string

	// if no names were entered name each pool XdY
	if *namesPtr == "" {
		names = append(names, dice_args...)
	} else {
		// Otherwise split names by a coma separator
		names = strings.Split(*namesPtr, ",")
	}
	table, err := dice.ParseTableString(dice_args, names)
	table.Name = *tablenamePtr
	if err != nil {
		fmt.Println(err)
	} else if *interactivePtr {
		tablecommands.InteractiveLoop(table)
	} else {
		table.Roll()
		for _, pool := range table.Pools {
			fmt.Printf("%s Their total is %d.\n", pool.Describe(), pool.Total())
		}
	}
}
