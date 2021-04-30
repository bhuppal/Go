package main

import "fmt"

//Calculate returns x + 2.
func Calculate(x int) int {
	result := x + 2
	return result
}

type Duration int64

func main() {

	// user defines a user in the program
	type user struct {
	 	name 		string
	 	email 		string
	 	ext 		int
	 	privileged 	bool
	}

	// Declare a variable of type user
	var me user
	me.name = "Bhuppal Kumar"
	me.email = "bhuppal@gmail.com"
	me.ext = 902
	me.privileged = true

	fmt.Println(me)

	rhonda := user{
		name: "Rhonda",
		email: "rhonda@gmail.com",
		ext: 203,
		privileged: true,
	}
	fmt.Println(rhonda)

	person := user{"person", "person@gmail.com", 993, true}

	fmt.Println(person)

	// Declaring fields based on other struct types
	// admin represents an admin user with privileges
	type admin struct {
		person user
		level string
	}

	//using struct literals to create values for fields
	fred := admin{
		person: user{
			name: "Kumar",
			email: "kumar@gmail.com",
			ext: 393,
			privileged: true,
		},
		level: "administrator",
	}

	fmt.Println(fred)

    var dur Duration
	dur = Duration(1000)
	fmt.Println(dur)

}
