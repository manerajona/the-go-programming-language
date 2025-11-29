package godoc

import (
	"fmt"
	"sync"
)

var nextID = 5

type Usr struct {
	ID       int
	Username string
}

var users = []Usr{
	Usr{
		ID:       1,
		Username: "adent",
	},
	Usr{
		ID:       2,
		Username: "tmacmillan",
	},
	Usr{
		ID:       3,
		Username: "fprefect",
	},
	Usr{
		ID:       4,
		Username: "zbeeblebrox",
	},
}

var m sync.RWMutex

func getAll() []Usr {
	m.RLock()
	defer m.RUnlock()
	return users
}

func add(u Usr) Usr {
	m.Lock()
	defer m.Unlock()
	u.ID = nextID
	nextID++
	users = append(users, u)
	return u
}

func getOne(id int) (Usr, error) {
	m.RLock()
	defer m.RUnlock()
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return Usr{}, fmt.Errorf("user not found with id %v", id)
}

func GetOne(id int) (Usr, error) {
	return getOne(id)
}

func update(u Usr, id int) (Usr, error) {
	m.Lock()
	defer m.Unlock()
	for i := range users {
		if users[i].ID == id {
			users[i] = u
			return u, nil
		}
	}
	return Usr{}, fmt.Errorf("user not found with id %v", id)
}

func delete(id int) bool {
	m.Lock()
	defer m.Unlock()
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
