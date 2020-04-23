package main

import (
	"fmt"
)

type errUserNameExist struct {
	UserName string
}

func (e errUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.UserName)
}
func isErrUserNameExist(err error) bool {
	_, ok := err.(errUserNameExist)
	return ok
}
func checkUserNameExist(username string) (bool, error) {
	if username == "bar" {
		return true, errUserNameExist{UserName: username}
	}
	if username == "foo" {
		return true, errUserNameExist{UserName: username}
	}
	return false, nil
}
func main() {
	if _, err := checkUserNameExist("foo"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		}
	}
	if _, err := checkUserNameExist("bar"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("isErrUserNameExist is false")
		}
	}
}
