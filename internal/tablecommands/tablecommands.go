package tablecommands

import (
	"bufio"
	"dicetable/pkg/dice"
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
	if _, ok := commands[command]; ok {
		return commands[command](table, args)
	} else {
		return fmt.Sprintf("%s is not a valid command. Maybe try help for a list of valid commands.", input)
	}
	
}

func help(table dice.Table, args []string) string {
	return `Commands:
	help -	Display this prompt
	roll - Roll the dice in any number of pools or roll all dice in the table
		format: roll [pool/table] {if pool} [pool names]
		examples: roll pool strength, roll table
	add - Add a die to a pool or a new pool to the table
		format: add [die/pool] {if die} [pool names...] {if pool} [pool names:XdY...]
		examples: add die strength agility, add pool power:4d6
	subtract - Subtract a dice from any number of pools, or pools from the table
		format: subtract [die/pool] {if die} [pool names:number of dice] {if pool} [pool names]
		examples: subtract die strngth:3 agility:1, subtract pool strength agility
	view - prints a discription of the pool or the whole table
		format: view [pool/table] {if pool} [pool names]
		examples: view pool strength, view table
	clear - removes all dice from a pool or all pools from the table
		format - clear [pool/table] {if pool} [pool names]
		examples: clear pool strength, clear table
	set - set a die to a number, all dice in a pool to the same number, or all dice on the table to the same number
		format: set die [pool name] [die position] [set to]/pool [pool name] [set to]/table [set to]`
}

func roll(table dice.Table, args []string) string {
	return_str := "Your Rolls:\n"
	var str string

	// Make sure that the user provides arguments
	if len(args) < 1 {
		return "Use roll pool [pool name] or roll table."
	}

	if args[0] == "pool" {
		// Pool argument rolls a list of pools in the table

		// Make sure that the pool name is provided with the pool argument
		if len(args[1:]) < 1 {
			return "Not the right ammount of arguments for roll pool [pool name]."
		}

		// Roll the dice for each pool name provided
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
		// Roll each pool in the table if the table argument is provided

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

	// Make sure at least two arguments are provided
	if len(args) < 2 {
		return "Not enough arguments provided. add [die/pool] [pool name/pool name:dice]"
	}

	if args[0] == "die" {
		// add a die to a pool

		// loop through each name provided
		for _, name := range args[1:] {

			// Make sure each pool name provided exisits
			if pool, ok := table.Pools[name]; ok {
				pool.Add()
				str = fmt.Sprintf("Successfully added die to pool %s. Now there are %dd%ds\n", name, pool.Sides, len(pool.Dice))
			} else {
				str = fmt.Sprintf("Pool %s does not exist.\n", name)
			}
			return_str = return_str + str
		}
	} else if args[0] == "pool" {
		// add a pool to the table

		// loop through each name provided
		for n, arg := range args[1:] {

			// make sure that the name and XdY are separated by a colon
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

	// Make sure at least two arguments are provided
	if len(args) < 2 {
		return "Not enough arguments provided. subtract [die/pool] [pool name:number of dice/pool name]"
	}

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
	// Print a descriptions of specific pools view pool [pool names...]
	// or all pools. view table
	return_str := "Pool Descriptions:\n"
	var str string

	// Make sure at least one argument is provided
	if len(args) < 1 {
		return "Not enough arguments provided. view [pool/table] [pool name]"
	}

	if args[0] == "pool" {
		// Make sure at least one additonal argument is provided
		if len(args[1:]) < 1 {
			return "Not enough arguments provided. view [pool/table] [pool name]"
		}

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
	// Clears all dice from a number of pools. clear pool [pool names...]
	// or clears all pools from the table. clear table
	return_str := "Cleared:\n"
	var str string

	// Make sure at least one argument is provided
	if len(args) < 1 {
		return "Not enough arguments provided. clear [pool/table] [pool name]"
	}

	if args[0] == "pool" {
		// Make sure at least one additonal argument is provided
		if len(args[1:]) < 1 {
			return "Not enough arguments provided. clear [pool/table] [pool name]"
		}

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
	// Set a specific die in a table to a number. set die [pool name] [die position] [set to]
	// or set all dice in a pool to a number. set pool [pool names...] [set to]
	// or set all dice in the table to a specific number
	return_str := "Set:"
	var str string

	// Make sure at least one argument is provided
	if len(args) < 1 {
		return "Not enough arguments provided. set [die/pool/table] [pool names...] [die position] [set to]"
	}

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
