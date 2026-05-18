package main

import "errors"

func CheckIfLifelineHasBeenUsed(input bool) error {
	if input {
		return errors.New(Color("THIS LIFELINE HAS ALREADY BEEN USED\nPICK ANOTHER", "\033[1;31m"))

	}
	return nil
}