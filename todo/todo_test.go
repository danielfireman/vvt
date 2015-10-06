package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s := InMemoryStore()
	assert.NotNil(t, s.Add(&item{""}))
	i := &item{"foo"}
	if assert.Nil(t, s.Add(i)) {
		assert.Equal(t, i, s.content[0], "item must be in the list")
	}
}

func TestList(t *testing.T) {
	content := []*item{&item{"foo"}, &item{"bar"}}
	s := store{content}
	for i, v := range s.List() {
		assert.Equal(t, s.content[i], v, "item must be in the list")
	}
}
