package diceapi_test

import (
	"testing"
	"time"

	"dicetable/pkg/diceapi"
)

var database_url string

func init() {
	database_url = "postgres://robert:theanswer42@localhost:5432/test"
}

func TestHello(t *testing.T) {
	greeting, err := diceapi.Hello()
	if err != nil {
		t.Errorf("%v", err)
	}
	if greeting != "Hello, world!" {
		t.Errorf("Greeting contained %v", greeting)
	}
}

func TestInsert(t *testing.T) {
	conn := diceapi.Connect(database_url)
	err := diceapi.UserInsert(conn, "testuser2", "password1", time.Now())
	if err != nil {
		t.Errorf("%v", err)
	}
	defer conn.Close()
}
