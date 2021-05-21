package dice_test

import (
	"dicetable/dice"
	"testing"
)

//Tests for Die struct methods
func TestRollDie(t *testing.T) {
	// Test to see if the top side of the die changes and is within the proper range when the die is rolled.
	die := dice.Die{Sides: 10000, Top: 1}
	starting := die.Top
	die.Roll()
	ending := die.Top

	// Roll the die and check to see if the top is a different number, if not, try again before giving an error
	if starting == ending {
		die.Roll()
		ending = die.Top
		if starting == ending {
			t.Errorf("Rolled die twice and got the same number. Either function isn't working or you're very lucky")
		}
	}

	// Give an error if the die rolls a number above the number of sides it has
	if ending > 10000 {
		t.Errorf("The die rolled %d, which is greater than the max roll on the 1d10000", ending)
	}

	// Give an error if the die rolls a number below one
	if ending < 1 {
		t.Errorf("The die rolled %d, which is lower than the minimum value of 1", ending)
	}
}

func TestSetDie(t *testing.T) {
	// Test to see if the Set method successfully sets the die to a specific side and will throw an error if trying to set it to a side greater than it's Sides
	die := dice.Die{Sides: 6, Top: 1}
	die.Set(5)

	// A d6 should be able to be set to 5
	if die.Top != 5 {
		t.Errorf("Die shows %d as it's top, but it should have been set to 5.", die.Top)
	}

	// A d6 should not be able to be set to 7
	err := die.Set(7)
	if die.Top == 7 {
		t.Errorf("Die could be set to a side higher than it's number of sides.")
	}
	if err == nil {
		t.Errorf("Set method did not detect an error when die was set to a number higher than it's number of sides")
	}

	// A d6 should not be able to be set to 0
	err = die.Set(0)
	if die.Top == 0 {
		t.Errorf("Die could be set to a side lower than it's number of sides")
	}
	if err == nil {
		t.Errorf("Set method did not detect an error when die was set to a number lower than it's number of sides")
	}
}

// Tests for the CreatePool function
func TestCreatePool(t *testing.T) {
	// CreatePool function should create a Pool using the first argument for the number of dice,
	// and the second for the number of sides to each die.
	pool := dice.CreatePool(5, 6)

	// A pool with 5 dice should be created
	if len(pool.Dice) != 5 {
		t.Errorf("CreatePool should have made a pool with 5 dice in it. Instead it has %d", len(pool.Dice))
	}

	// The pool should have six sides.
	// In addition checks each individual die in the pool to make sure it also has six sides
	if pool.Sides != 6 {
		t.Errorf("CreatePool should have created a pool of d6s instead it created a pool of d%ds", pool.Sides)
	} else {
		for n, die := range pool.Dice {
			if die.Sides != 6 {
				t.Errorf("CreatePool created a pool with the correct number of sides, but die%d had %d sides instead", n, die.Sides)
			}
		}
	}

}

func TestParseDiceString(t *testing.T) {
	// ParseDiceString function should parse through a string in the format of XdY and return a *Pool
	// If the string is in the wrong format it should return an error
	good_str := "3d6"
	test, err := dice.ParseDiceString(good_str)

	// ParseDiceString should not throw an error here
	if err != nil {
		t.Errorf("ParseDiceString returned an error when it was not supposed to.")
	}

	// ParseDiceString should return a pool with 3 dice
	if len(test.Dice) != 3 {
		t.Errorf("ParseDiceString should have created a pool with 3 dice. Instead it has %d", len(test.Dice))
	}

	// ParseDiceString should return a pool of d6s
	if test.Sides != 6 {
		t.Errorf("ParseDiceString should have created a pool of d6s but instead it created a pool of d%ds", test.Sides)
	}

	// Test that the proper errors are returned
	wrong_letter := "3c6"
	test, err = dice.ParseDiceString(wrong_letter)
	if err == nil {
		t.Errorf("ParseDiceString did not return an error when the wrong separator letter was used.")
	}
	if test != nil {
		t.Errorf("ParseDiceString returned a Pool when the wrong letter was used. Instead it should have returned nil")
	}

	too_many_splits := "3d6d7"
	test, err = dice.ParseDiceString(too_many_splits)
	if err == nil {
		t.Errorf("ParseDiceString did not return an error when too many ds were in the string")
	}
	if test != nil {
		t.Errorf("ParseDiceString returned a Pool when too many ds were in the string. Instead it should have returned nil")
	}

	no_numbers := "adc"
	test, err = dice.ParseDiceString(no_numbers)
	if err == nil {
		t.Errorf("ParseDiceString did not return an error when there were no numbers in the string")
	}
	if test != nil {
		t.Errorf("ParseDiceString returned a Pool when there were no numbers in the string. Instead it should have returned nil")
	}
}

// Tests for Pool struct methods
func TestRollPool(t *testing.T) {
	// Test to see if each die in a pool is rolled when the Roll function is called on the pool
	pool := dice.CreatePool(3, 10000)
	pool.Roll()

	// Check each die in the pool
	for n, die := range pool.Dice {

		// Make sure the top number on each die changes when Roll() is called
		if die.Top == 1 {
			pool.Roll()
			if die.Top == 1 {
				t.Errorf("Rolled die twice and got 1 both times on die%d", n)
			}
		}

		// Make sure no die is over it's maximum number of sides
		if die.Top > 10000 {
			t.Errorf("The die%d rolled %d, which is greater than the max roll on the 1d10000", n, die.Top)
		}

		// Make sure no die rolls under 1
		if die.Top < 1 {
			t.Errorf("The die%d rolled %d, which is lower than the minimum value of 1", n, die.Top)
		}
	}
}

func TestListPool(t *testing.T) {
	// The List method of pool should return a slice containing the Top of each die.
	pool := dice.CreatePool(3, 4)
	list := pool.List()
	default_list := []int{1, 1, 1}

	// The default pool created should have all ones as it's top so List() should return [1,1,1]
	for n, l := range list {
		if l != default_list[n] {
			t.Errorf("List should have returned [1, 1, 1], but instead returned %d", list)
		}
	}

	// Check to make sure the function displays the proper values after the dice are changed
	changed_list := []int{2, 3, 4}
	for n, d := range changed_list {
		pool.Dice[n].Set(d)
	}
	list = pool.List()
	for n, l := range list {
		if l != changed_list[n] {
			t.Errorf("List should have returned [2, 3, 4], but instead returned %d", list)
		}
	}
}

func TestTotalPool(t *testing.T) {
	// The Total method should return the int sum of each Top of each die in the pool
	pool := dice.CreatePool(3, 6)
	sum := pool.Total()
	if sum != 3 {
		t.Errorf("Total should have been 3 but instead was %d", sum)
	}
}

func TestAddPool(t *testing.T) {
	// The Add method should add a dice to the pool set with 1 as the top
	pool := dice.CreatePool(3, 6)
	pool.Add()
	// Make sure one die was added to the pool
	if len(pool.Dice) != 4 {
		if len(pool.Dice) == 3 {
			t.Errorf("Add() did not add a die to the pool")
		} else {
			t.Errorf("After adding a die to the pool there should have been 4 dice, instead there were %d", len(pool.Dice))
		}
	}

	added_die := pool.Dice[3]
	// Make sure the die has the proper number of sides
	if added_die.Sides != 6 {
		t.Errorf("The die added to the pool should have 6 sides, but instead it has %d", added_die.Sides)
	}

	// Make sure the die is set with 1 as it's Top
	if added_die.Top != 1 {
		t.Errorf("The die added should have 1 on the top, but instead has %d", added_die.Top)
	}
}

func TestSubtractPool(t *testing.T) {
	// The Subtract function should remove a die from the pool.
	// it should throw an error if there are no dice in the pool
	pool := dice.CreatePool(1, 6)
	pool.Subtract()
	if len(pool.Dice) != 0 {
		t.Errorf("After subtracting a die from the pool there should be no dice left, instead there are %d dice left", len(pool.Dice))
	}
	err := pool.Subtract()
	if err == nil {
		t.Errorf("When trying to subtract from an empty pool Subtract should have returned an error. It did not.")
	}
}

func TestDescribePool(t *testing.T) {
	// The Description method should return a string that describes the dice in the pool
	pool := dice.CreatePool(3, 6)
	pool.Dice[1].Set(5)
	description := "A pool of 3 d6s. The dice are facing 1, 5, and 1."
	if pool.Describe() != description {
		t.Errorf("The Describe function returned %s rather than the proper description.", pool.Describe())
	}
}

// Tests for CreateTable function
func TestCreateTable(t *testing.T) {
	// The CreateTable function should create a Table with a list of pools and a list of names for those pools.
	// It should return the Table struct and an error for if there is a different number of names and pools.
	pools := []*dice.Pool{dice.CreatePool(2, 6), dice.CreatePool(3, 10)}
	names := []string{"d6s", "d10s"}
	table, err := dice.CreateTable(pools, names)

	// Nil should be returned for err in this instance
	if err != nil {
		t.Errorf("CreateTable function returned an error when it should not have.")
	}

	// There should be two pools in the table
	if len(table.Pools) != 2 {
		t.Errorf("The CreateTable fuction should have created a table with 2 pools, but instead it has %d", len(table.Pools))
	}

	// Check each pool is matched to the right name
	if len(table.Pools["d6s"].Dice) != 2 {
		t.Errorf("The first name should have been assigned to a pool of 2d6s. Instead it had %dd%ds.", len(table.Pools["d6s"].Dice), table.Pools["d6s"].Sides)
	}
	if len(table.Pools["d10s"].Dice) != 3 {
		t.Errorf("The first name should have been assigned to a pool of 3d10s. Instead it had %dd%ds.", len(table.Pools["d10s"].Dice), table.Pools["d10s"].Sides)
	}

	// The fuction should return an error if a mismatched number of pools and names is entered
	names = append(names, "d29s")
	_, err = dice.CreateTable(pools, names)
	if err == nil {
		t.Errorf("CreateTable did not return an error when it should have.")
	}

}

func TestParseTableString(t *testing.T) {
	// The ParseTableString function should take in a list of strings in XdY format and return a Table
	// with those dice pools, and an error if any of the strings are not in the corredt format
	table_list := []string{"2d6", "3d10"}
	names := []string{"pool1", "pool2"}
	table, err := dice.ParseTableString(table_list, names)

	// ParseTableString should not throw an error
	if err != nil {
		t.Errorf("ParseTableString returned an error when it should not have")
	}

	// Table should have 2 pools
	if len(table.Pools) != 2 {
		t.Errorf("ParseTableString should have returned a table with 2 pools. Instead it has %d", len(table.Pools))
	}

	// Make sure each name is in the pool
	for _, name := range names {
		if _, ok := table.Pools[name]; !ok {
			t.Errorf("ParseTableString did not add %s under the correct name", name)
		}
	}

}

// Tests for Table struct methods

func TestRollTable(t *testing.T) {
	// The Roll method for table should aplly the roll method to each pool in the table
	pools := []*dice.Pool{dice.CreatePool(2, 10000), dice.CreatePool(3, 1000)}
	names := []string{"pool1", "pool2"}
	table, _ := dice.CreateTable(pools, names)
	table.Roll()

	// Test each pool in the table
	for name, pool := range table.Pools {
		for n, die := range pool.Dice {
			if die.Top == 1 {
				pool.Roll()
				if die.Top == 1 {
					t.Errorf("In %s rolled die twice and got 1 both times on die%d", name, n)
				}
			}

			// Make sure no die rolls under 1
			if die.Top < 1 {
				t.Errorf("In %s the die%d rolled %d, which is lower than the minimum value of 1", name, n, die.Top)
			}
		}
	}
}

func TestClearTable(t *testing.T) {
	// The Clear method should remove all pools from the table
	pools := []*dice.Pool{dice.CreatePool(2, 6), dice.CreatePool(3, 10)}
	names := []string{"d6s", "d10s"}
	table, _ := dice.CreateTable(pools, names)
	table.Clear()

	// The table should now have no pools in it
	if len(table.Pools) != 0 {
		t.Errorf("There are still %d pools left in the table after the Clear function was called", len(table.Pools))
	}
}

func TestRemoveTable(t *testing.T) {
	// The Remove method should take in a string as an argument and remove the pool with that name from the table.
	// An error should be returned if there is no pool with that name in the table.
	pools := []*dice.Pool{dice.CreatePool(2, 6), dice.CreatePool(3, 10)}
	names := []string{"d6s", "d10s"}
	table, _ := dice.CreateTable(pools, names)
	table.Remove("d6s")

	// Pool called d6s should no longer be in the table
	if _, ok := table.Pools["d6s"]; ok {
		t.Errorf("Remove function did not successfully remove d6s pool from the table.")
	}

	// Remove should return an error if trying to remove a pool that doesn't exist in the table
	err := table.Remove("d20s")
	if err == nil {
		t.Errorf("Remove function should have returned an error if a name that doesn't exist was entered.")
	}
}

func TestAddTable(t *testing.T) {
	// The Add function for a table should take in a name, size and sides as arguments and create a pool in the table
	// with those parameters
	pools := []*dice.Pool{dice.CreatePool(2, 6), dice.CreatePool(3, 10)}
	names := []string{"d6s", "d10s"}
	table, _ := dice.CreateTable(pools, names)
	table.Add("d20s", 2, 20)

	// Make sure a pool was added to the table
	if len(table.Pools) != 3 {
		t.Errorf("Add function did not add a pool to the table. There should be 3 pools but there are %d instead", len(table.Pools))
	} else {

		// If the table was added check if the name d20s exists
		if pool, ok := table.Pools["d20s"]; ok {

			// The pool should be made up of 20 sided dice
			if pool.Sides != 20 {
				t.Errorf("Add function should have created a pool of d20s. Instead it created a pool of d%ds", pool.Sides)
			}

			// There should be two dice in the pool
			if len(pool.Dice) != 2 {
				t.Errorf("Add function should have created a pool of 2 dice. Instead it created a pool of %d dice", len(pool.Dice))
			}
		} else {

			// If there is no pool with the name d20s check to see what names do exist in the table
			n := make([]string, 4)
			for name := range table.Pools {
				n = append(n, name)
			}
			t.Errorf("Add function should have created a pool with the name d20s. Instead the names in the table are %s", n)
		}
	}

}
