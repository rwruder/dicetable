package main

import (
	"bufio"
	"dicetable/dice"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InteractiveLoop(table dice.Table) {
	// A looping function meant to simulate rolling dice at a table
	for {
		// Display the prompt. If the table has a name add it to the prompt
		prompt := table.Name + ":> "
		fmt.Print(prompt)

		// Read from stdin until there's a newline
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		// send the command to the tablecommands.ParseCommand for parsing
		if command == "exit" {
			fmt.Println("Goodbye...")
			break
		} else {
			fmt.Println(ParseCommand(command, table))
		}
	}
}

func ParseCommand(input string, table dice.Table) string {
	// Create a map with command names as keys and their functions as the values
	// Return a string as an answer to the command

	command := strings.Split(input, " ")[0]
	args := strings.Split(input, " ")[1:]

	commands := map[string]func(dice.Table, []string) string{
		"help":     help,
		"roll":     roll,
		"add":      add,
		"subtract": subtract,
		"view":     view,
		"clear":    clear,
		"set":      set,
	}
	return commands[command](table, args)
}

func help(table dice.Table, args []string) string {
	return ""
}

func roll(table dice.Table, args []string) string {
	return_str := "Your Rolls:\n"
	var str string

	if args[0] == "pool" {
		for _, name := range args[1:] {
			if pool, ok := table.Pools[name]; ok {
				pool.Roll()
				str = fmt.Sprintf("Pool %s: %d Total: %d\n", name, pool.List(), pool.Total())
			} else {
				str = fmt.Sprintf("Pool %s does not exist.\n", name)
			}
			return_str = return_str + str
		}
	} else if args[0] == "table" {
		table.Roll()
		for name, pool := range table.Pools {
			str = fmt.Sprintf("Pool %s: %d Total: %d\n", name, pool.List(), pool.Total())
			return_str = return_str + str
		}
	} else {
		return_str = "roll command format is roll [table or pool] [pool names if pool]"
	}
	return return_str
}

func add(table dice.Table, args []string) string {
	return_str := "Added:\n"
	var str string

	if args[0] == "die" {
		for _, name := range args[1:] {
			if pool, ok := table.Pools[name]; ok {
				pool.Add()
				str = fmt.Sprintf("Successfully added die to pool %s. Now there are %dd%ds\n", name, pool.Sides, len(pool.Dice))
			} else {
				str = fmt.Sprintf("Pool %s does not exist.\n", name)
			}
			return_str = return_str + str
		}
	} else if args[0] == "pool" {
		for n, arg := range args[1:] {
			if !strings.Contains(arg, ":") {
				str = fmt.Sprintf("Arg %d failed. Format is [name]:[XdY]", n)
				return_str = return_str + str
				continue
			}
			a := strings.Split(arg, ":")
			name := a[0]
			pool, err := dice.ParseDiceString(a[1])
			if err != nil {
				str = fmt.Sprintf("%s\n", err)
				return_str = return_str + str
				continue
			}
			table.Pools[name] = pool
			str = fmt.Sprintf("Successfully added pool %s of %dd%ds to the table.\n", name, len(pool.Dice), pool.Sides)
			return_str = return_str + str
		}
	} else {
		return_str = "add command format is add [die or pool] [pool name:XdY(if add pool)]"
	}
	return return_str
}

func subtract(table dice.Table, args []string) string {
	return_str := "Subtracted:"
	var str string
	var err error

	if args[0] == "die" {
		for _, name := range args[1:] {
			i := 1
			plural := ""

			// Check for a colon in argument. If so subtract multiple dice
			if strings.Contains(name, ":") {
				plural = "e"
				s := strings.Split(name, ":")
				name = s[0]
				i, err = strconv.Atoi(s[1])
				if err != nil {
					str = fmt.Sprintf("%s\n", err)
					return_str = return_str + str
					continue
				}
			}
			if pool, ok := table.Pools[name]; ok {
				for c := 0; c < i; c++ {
					err = pool.Subtract()
					if err != nil {
						str = fmt.Sprintf("%s\n", err)
						return_str = return_str + str
						break
					}
				}
				str = fmt.Sprintf("Successfully subtracted %d di%se from pool %s. Now there are %dd%ds\n", i, plural, name, pool.Sides, len(pool.Dice))
			} else {
				str = fmt.Sprintf("Pool %s does not exist.\n", name)
			}
			return_str = return_str + str
		}
	} else if args[0] == "pool" {
		for _, arg := range args[1:] {
			err = table.Remove(arg)
			if err != nil {
				str = fmt.Sprintf("%s\n", err)
				return_str = return_str + str
				continue
			}
			str = fmt.Sprintf("Successfully removed %s pool from table.", arg)
			return_str = return_str + str
		}
	} else {
		return_str = "subtract command format is add [die or pool] [pool name:number of dice if dice]"
	}
	return return_str
}

func view(table dice.Table, args []string) string {
	return_str := "Pool Descriptions:\n"
	var str string

	if args[0] == "pool" {
		for _, name := range args[1:] {
			if _, ok := table.Pools[name]; !ok {
				str = fmt.Sprintf("%s is not the name of a pool on the table.\n", name)
				return_str = return_str + str
				continue
			}
			str = fmt.Sprintf("%s: %s\n", name, table.Pools[name].Describe())
			return_str = return_str + str
		}
	} else if args[0] == "table" {
		for name, pool := range table.Pools {
			str = fmt.Sprintf("%s: %s\n", name, pool.Describe())
			return_str = return_str + str
		}
	} else {
		return_str = "view command format is view [pool or table] if pool [pool name...]"
	}
	return return_str
}

func clear(table dice.Table, args []string) string {
	return_str := "Cleared:\n"
	var str string

	if args[0] == "pool" {
		for _, name := range args[1:] {
			if _, ok := table.Pools[name]; !ok {
				str = fmt.Sprintf("%s is not the name of a pool on the table.\n", name)
				return_str = return_str + str
				continue
			}
			pool := table.Pools[name]
			for x := 0; x < len(pool.Dice); x++ {
				err := pool.Subtract()
				if err != nil {
					break
				}
			}
			str = fmt.Sprintf("Cleared pool %s\n", name)
			return_str = return_str + str
		}
	} else if args[0] == "table" {
		table.Clear()
		str = "Cleared table\n"
		return_str = return_str + str
	}
	return return_str
}

func set(table dice.Table, args []string) string {
	return_str := "Set:"
	var str string

	if args[0] == "die" {
		if len(args) != 4 {
			return_str = "Not enough arguments to set a die. die [pool] [die] [set to].\n"
			return return_str
		}
		pool_name := args[1]
		if _, ok := table.Pools[pool_name]; !ok {
			str = fmt.Sprintf("%s is not the name of a pool on the table.\n", pool_name)
			return_str = return_str + str
			return return_str
		}
		die, err := strconv.Atoi(args[2])
		if err != nil {
			str = fmt.Sprintf("%s", err)
			return_str = return_str + str
			return return_str
		}
		set_to, err := strconv.Atoi(args[3])
		if err != nil {
			str = fmt.Sprintf("%s", err)
			return_str = return_str + str
			return return_str
		}
		err = table.Pools[pool_name].Dice[die].Set(set_to)
		if err != nil {
			str = fmt.Sprintf("%s", err)
			return_str = return_str + str
			return return_str
		}
		str = fmt.Sprintf("Die %d in pool %s successfully set to %d.\n", die, pool_name, set_to)
		return_str = return_str + str
	} else if args[0] == "pool" {
		if len(args) != 3 {
			return_str = "Not enough arguments to set a pool. pool [pool] [set to].\n"
			return return_str
		}
		pool_name := args[1]
		if _, ok := table.Pools[pool_name]; !ok {
			str = fmt.Sprintf("%s is not the name of a pool on the table.\n", pool_name)
			return_str = return_str + str
			return return_str
		}
		set_to, err := strconv.Atoi(args[3])
		if err != nil {
			str = fmt.Sprintf("%s", err)
			return_str = return_str + str
			return return_str
		}
		pool := table.Pools[pool_name]
		for _, die := range pool.Dice {
			err = die.Set(set_to)
			if err != nil {
				str = fmt.Sprintf("%s", err)
				return_str = return_str + str
				return return_str
			}
		}
		str = fmt.Sprintf("Successfully set all dice in %s pool to %d", pool_name, set_to)
		return_str = return_str + str
	} else if args[0] == "table" {
		if len(args) != 2 {
			return_str = "Not enough arguments to set a table. table [set to].\n"
			return return_str
		}
		set_to, err := strconv.Atoi(args[3])
		if err != nil {
			str = fmt.Sprintf("%s", err)
			return_str = return_str + str
			return return_str
		}
		for _, pool := range table.Pools {
			for _, die := range pool.Dice {
				err = die.Set(set_to)
				if err != nil {
					str = fmt.Sprintf("%s", err)
					return_str = return_str + str
					return return_str
				}
			}
		}
		str = fmt.Sprintf("Successfully set all dice in on the table to %d", set_to)
		return_str = return_str + str
	} else {
		return_str = `set command format set [die, pool, or table]: 
		die [pool name] [die] [set_to]
		pool [pool name] [set to]
		table [set to]`
	}
	return return_str
}
