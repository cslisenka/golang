package main

import (
	"fmt"
)

type Student struct {
	name    string
	surname string
	age     int
}

type Examination struct {
	// This struct has a reference to other struct "Student"
	student *Student
	mark    int
}

type Specialist struct {
	// Composition, this struct inerhits values from student
	// We may not use inheritance in go
	*Student
	speciality string
}

func incrementAge(st *Student) {
	// We pass a pointer to the student structure declared outside of the function
	fmt.Println(st)
	(*st).age++
}

func sayGreetings(st Student) {
	// Student structure is copied (full copy)
	fmt.Println("Greetings ", st.name, " !")
}

// Assiciate function with struct
func (s *Student) sayHi() {
	fmt.Println("Hi! ", s.name)
}

func (s *Student) sayGreeting(greeting string) {
	fmt.Println(greeting, " ", s.name)
}

// Instead of constructor we create function that returns intance of structure
func NewStudent(n string, s string, ag int) Student {
	st := Student{name: n, surname: s, age: ag}
	return st
}

func main() {
	// Create instance of struct
	var st1 Student = Student{"Anna", "Bond", 28}

	// Change variables
	st1.age = 22

	// This is copy of the structure
	st2 := st1

	// This is pointer, not a copy
	st3 := &st1

	fmt.Println(st1, st2, st3)

	incrementAge(&st1)
	sayGreetings(st1)

	fmt.Println(st1, st2, st3)

	// We may pass attributes by names, if we skip - attribute doesn't get initialized
	var st4 Student = Student{surname: "Bond", age: 40}
	fmt.Println(st4)

	// Testing of methods associated with struct
	st1.sayHi()
	st1.sayGreeting("Gomarjoba!")

	// Using constructor function
	st5 := NewStudent("aaa", "bbb", 111)
	fmt.Println(st5)

	// Create nested structs
	ex := Examination{
		mark:    10,
		student: &Student{name: "Test"},
	}
	fmt.Println(ex)

	sp := Specialist{speciality: "it", Student: &Student{}}
	// We assign to sp.Student.surname, but keep short syntax here
	sp.surname = "test"
	sp.name = "test name"
	fmt.Println(sp.name, sp.surname, sp.speciality)
	fmt.Println(sp.Student.name, sp.Student.surname, sp.speciality)
	fmt.Println(sp)

	// New structure also has all functions associated with Student
	sp.sayHi()
}
