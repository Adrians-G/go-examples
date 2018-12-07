package main

// Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Value receiver
func (p Person) greet() string {
	return "Hello my name is " + p.firstName
}

// Pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

// Pointer receiver (marry method)
func (p *Person) marry(someone Person) {
	p.lastName = someone.lastName
}
