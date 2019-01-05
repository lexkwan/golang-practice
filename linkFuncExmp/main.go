package main

import "fmt"

//Student Define student stuct
type Student struct {
	Name string
	Age  int
}

//SetName set the name for the student
func (s *Student) SetName(name string) *Student {
	s.Name = name
	return s
}

//SetAge set student age
func (s *Student) SetAge(age int) *Student {
	s.Age = age
	return s
}

//Print output the name and age for the student.
func (s *Student) Print() {
	fmt.Printf("name:%s, age:%d\n", s.Name, s.Age)
}

func main() {
	xiaoming := &Student{}
	xiaoming.SetName("xiaoming").SetAge(20).Print()
}
