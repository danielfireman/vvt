package todo

import "testing"

func TestAdd(t *testing.T) {
	s := NewStore()
	if err := s.Add(&item{""}); err == nil {
		t.Errorf("must error, got:nil")
	}
	i := &item{"foo"}
	if err := s.Add(i); err != nil {
		t.Errorf("want:nil, got:%q", err)
	}
	if s.content[0] != i {
		t.Errorf("want:%v, got:%v", i, s.content[0])
	}
}

func TestList(t *testing.T) {
	content := []*item{&item{"foo"}, &item{"bar"}}
	s := store{content}
	for i, v := range s.List() {
		if s.content[i] != v {
			t.Errorf("want:%v, got:%v", content[i], v)
		}
	}
}
