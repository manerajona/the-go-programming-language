package main

import (
	"fmt"
)

// Model
type person struct {
	firstName, lastName string
}

// Repository layer
type personRepository interface {
	save(id int, p person)
	findById(id int) person
}

type personMongoRepository map[int]person
type personPostgresRepository map[int]person

func (r personMongoRepository) save(id int, p person) {
	r[id] = p
}

func (r personMongoRepository) findById(id int) person {
	return r[id]
}

func (r personPostgresRepository) save(id int, p person) {
	r[id] = p
}

func (r personPostgresRepository) findById(id int) person {
	return r[id]
}

// service layer
type personService struct {
	repo personRepository
}

func (s personService) create(id int, p person) (err error) {

	if s.repo.findById(id) == (person{}) {
		s.repo.save(id, p)
	} else {
		err = fmt.Errorf("person with id=%d already exists", id)
	}
	return
}

func (s personService) get(id int) (p person, err error) {
	p = s.repo.findById(id)
	if p == (person{}) {
		err = fmt.Errorf("person with id=%d not found", id)
	}
	return
}

func main() {
	//service := personService{personMongoRepository{}}
	service := personService{personPostgresRepository{}}

	err0 := service.create(1, person{"John", "Doe"})
	if err0 == nil {
		fmt.Println("person created")
	}

	err1 := service.create(1, person{"Jane", "Smith"})
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	p, _ := service.get(1)
	fmt.Println(p.firstName + " " + p.lastName)

	_, err2 := service.get(99)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
}
