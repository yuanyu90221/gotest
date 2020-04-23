package main

import (
	"fmt"
	"strings"

	_ "github.com/yuanyu90221/gotest/foo"
)

func getUserListSQL(username, email string) string {
	sql := "select * from user"
	where := []string{}
	if username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", username))
	}
	if email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", email))
	}

	return sql + " where " + strings.Join(where, " or ")
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}
	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}
	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}

	if opts.sexy != 0 {
		where = append(where, fmt.Sprintf("sexy = %d", opts.sexy))
	}
	return sql + " where " + strings.Join(where, " or ")
}

type searchOpts struct {
	username string
	email    string
	sexy     int
}

func main() {
	// foo := func() string {
	// 	return "Hello World"
	// }
	// fmt.Println(foo())
	// bar := func() {
	// 	fmt.Println("Hello World 2")
	// }
	// bar()

	// func() {
	// 	fmt.Println("Hello World3")
	// }()
	fmt.Println(getUserListSQL("appleboy", ""))
	fmt.Println(getUserListSQL("appleboy", "test@gmail.com"))
	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "appleboy",
		email:    "test@gmail.com",
	}))
	fmt.Println(getUserListOptsSQL(searchOpts{
		sexy: 2,
	}))
}
