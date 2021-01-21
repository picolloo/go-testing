package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(1,2)

	if got == 3 {
		t.Logf("Expected 3, Got %d", got)
	}

	if got != 3 {
		t.Errorf("Expected 3, Got %d", got)
	}
}

func TestUser(t *testing.T) {
	user := user{1,"jon", 20}

	if user.id !=  1 {
		t.Errorf("Excpected %d, Got %d", 1,user.id)
	}
}
