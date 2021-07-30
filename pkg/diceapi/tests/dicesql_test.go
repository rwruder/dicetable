package diceapi_test

import (
	"testing"

	"dicetable/pkg/diceapi"
)

func TestHello(t *testing.T) {
	greeting, err := diceapi.Hello()
	if err != nil {
		t.Errorf("%v", err)
	}
	if greeting != "Hello, world!" {
		t.Errorf("Greeting contained %v", greeting)
	}
}
