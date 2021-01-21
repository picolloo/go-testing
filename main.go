package main

import (
	"errors"
	"fmt"
)


func sum(a,b int)int {
	return a + b
}

type user struct {
	id int
	name string
	balance float32
}

func (u *user) withdraw(amount float32) error {
	if u.balance - amount < 0 {
		return errors.New("Insufficient funds")
	}

	u.balance -= amount
	return nil
}

func main() {
	fmt.Printf("1 + 2 = %d", sum(1,2))
}