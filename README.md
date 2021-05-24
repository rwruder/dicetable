# Dicetable
dicetable is a command line tool that can be used to quickly roll a number of pools of dice, or it can create a interactive prompt with those dice pools.

## Flags:
-i - sets the mode to the interactive prompt
-names - strings separated by a comma this will change the name of each dice pool
-tablename - a string that changes the name of the table. Right now this only changes the interactive prompt but in the future I plan on using this to save and access different sets of dice

## Examples

> dicetable 3d6 4d8
This will roll a pool of 3d6s and 4d8s and output them to the command line.

> dicetable -i 3d6 4d8
This will create a interactive prompt with 2 dice pools named 3d6 and 4d8

> dicetable -i -names=strength,agility 4d6 2d6
This will create a prompt with two dice pools one named strength with 4d6 and another with 2d6 named agility

> dicetable -i -tablename=MyTable
Will make a prompt with no dice pools named MyTable

## Interactive Prompt:
Running dicetable with the -i flag will open up a prompt for the table with the name given by the -tablename flag. The interactive table is supposed to be like a real table where dice can be divided up into pools, rolled, and the dice will stay in a persistant state. This is to help with games where dice are rolled and then the numbers the dice display are saved and used over the course of the game as opposed to a system that uses the results of the roll immediately. 

##### Commands:
    help -	Display this list of help commands
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
		format: set [die] [die position] [set to]/[pool] [pool name] [set to]/[table] [set to]

###### Improvements:
- Remove function to remove specific dice from pools based on position or what number they are facing
- function to prompt for user input to add a check to make sure the user wants to clear a pool or table

###### Future Updates:
- Logging
- Database
- Server
- Discord Bot