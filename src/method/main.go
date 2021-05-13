package main

import "fmt"

// user defines a user in the program
type user struct {
	name  string
	email string
}

//******
// The parameter between the keyword func and the function name
// is called a receiver and binds the function to the specified type.
// When a function has a receiver, that function is called a method.

// notify implements a method with a value receiver
/*
When we declare a method using a value receiver, the method will always be operating
against a copy of the value used to make the method call.
*/
func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver
func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}

func main() {
	//fmt.Println("Methods....")

	bhuppal := user{"Bhuppal", "bhuppal@gmail.com"}
	bhuppal.notify()

	rhonda := &user{"Rhonda", "rhonda@gmail.com"}
	rhonda.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	bhuppal.changeEmail("newEmail@email.com")
	bhuppal.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	rhonda.changeEmail("newEmailAddress@email.com")
	rhonda.notify()

	fmt.Println(bhuppal)
	fmt.Println(rhonda)
}
