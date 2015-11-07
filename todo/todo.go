// Package todos handles todo simple string todo list stores.
package todo

import "errors"

// Creates a new todo item store.
func InMemoryStore() *store {
	return &store{}
}

// todo item.
type item struct {
	Desc string `desc`
}

// todo item store.
type store struct {
	content []*item
}

// Adds a todo item to the store.
func (s *store) Add(i *item) error {
	return errors.New("Invalid item.")
	/*if len(i.Desc) == 0 {
		return errors.New("Invalid item.")
	} else {
		s.content = append(s.content, i)
		return nil
	}*/
}

// Returns a copy of the todo items stored.
func (s *store) List() []*item {
	c := make([]*item, len(s.content))
	copy(c, s.content)
	return c
}
