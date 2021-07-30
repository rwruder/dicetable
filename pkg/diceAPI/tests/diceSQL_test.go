package diceAPI_test

import (
	"testing"

	"dicetable/pkg/diceAPI"
)

func TestHello(t *testing.T) {
	greeting, err := diceAPI.Hello()
	if err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("%v", greeting)
	}
}
