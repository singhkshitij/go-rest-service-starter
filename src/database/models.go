package database

import "fmt"

type User struct {
	FirstName string `pg:",notnull"`
	LastName  string `pg:",notnull"`
	Age       uint8
	Email     string
}

func (u User) String() string {
	return fmt.Sprintf("User<%s %s %d %s>", u.FirstName, u.LastName, u.Age, u.Email)
}
