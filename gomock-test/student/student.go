package student

import "gomock-test/person"

type Student struct {
	Name string
	p    person.Person
}

func (s *Student) Eat() string {
	return s.p.Eat()
}

func (s *Student) Sleep() string {
	return s.p.Sleep(s.Name)
}
