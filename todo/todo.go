// Package todos handles todo simple string todo list stores.
package todo

import "errors"

// Creates a new todo item store.
func NewStore() *store {
	return &store{}
}

// todo item store.
type store struct {
	content []string
}

// Adds a todo item to the store.
func (s *store) Add(i string) error {
	if len(i) == 0 {
		return errors.New("Invalid item.")
	}
	s.content = append(s.content, i)
	return nil
}

// Returns a copy of the todo items stored.
func (s *store) List() []string {
	c := make([]string, len(s.content))
	copy(c, s.content)
	return c
}
