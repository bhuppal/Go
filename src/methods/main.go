package main

import "fmt"

type user struct {
	name string
	email string
}

// The parameter between the keyword func and the function name is called a "receiver"
// and binds the function to a specified type. when a function has a receiver,
// the funciton is called a method.

// Value receiver operate on a copy of the value used to make the method call
// Pointer receiver operate on the actual value
// You can also call methods that are declared with a pointer receiver using a value

// notify() implements a method with a value receiver
func (u user) notify(newName string) {
	u.name = newName
	fmt.Println("Sending User Email To %s<%s>",
		u.name,
		u.email)

}

// changeEmail() implements a method with a poitner receiver
func (u *user) changeEmail(newEmail string) {
	u.email = newEmail
}


func main() {
	person1 := &user{name: "Bhuppal", email:"bhuppal@gmail.com"}

	fmt.Printf("Memory address of Person1 %v\n",&person1);
	person2 := &person1 //assigning memory address of person1 to person2
	fmt.Printf("Memory address of Person2 %v\n", person2) // both person1 and person2 are pointing to same memory address location

	fmt.Printf("Person2 values are %s %s\n",(*person2).name, (*person2).email)

	person1.name = "Rhonda"
	fmt.Printf("Name change %s\n",(*person2).name)

	(*person2).name = "Kumar"
	fmt.Printf("Name change %s\n",(*person1).name)


	// Values of type user can be used to call methods
	// declared with a value of receiver.
	bill := user{"Bill", "bill@gmail.com"}
	bill.notify("Temp")

	fmt.Println("Sending User Email To %s<%s>",
		bill.name,
		bill.email)

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	lisa := &user{"Lisa", "lisa@gmail.com"}
	(*lisa).notify("Temp")

	fmt.Println("Sending User Email To %s<%s>",
		lisa.name,
		lisa.email)

	// Values of type user can be used to call methods
	// declared with a pointer of receiver
	bill.changeEmail("newEmail@email.com")
	bill.notify("Temp")

	fmt.Println("Sending User Email To %s<%s>",
		bill.name,
		bill.email)

	// Pointer of type user can be used to call methods
	// declared with a pointer of receiver
	lisa.changeEmail("lisaNewEmail@email.com")
	lisa.notify("Temp")

	fmt.Println("Sending User Email To %s<%s>",
		lisa.name,
		lisa.email)
}
