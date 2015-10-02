package todo

import "testing"

func TestAdd(t *testing.T) {
	s := NewStore()
	if err := s.Add(""); err == nil {
		t.Errorf("must error, got:nil")
	}
	v := "foo"
	if err := s.Add(v); err != nil {
		t.Errorf("want:nil, got:%q", err)
	}
	if s.content[0] != v {
		t.Errorf("want:%v, got:%v", v, s.content[0])
	}
}

func TestList(t *testing.T) {
	content := []string{"foo", "bar"}
	s := store{content}
	for i, v := range s.List() {
		if s.content[i] != v {
			t.Errorf("want:%v, got:%v", content[i], v)
		}
	}
}
