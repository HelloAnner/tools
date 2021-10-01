package settools

import "sort"

type Empty struct{}

type String map[string]Empty

func NewString(items ...string) String {
	ss := String{}
	ss.Insert(items...)
	return ss
}

func (s String) Insert(items ...string) String {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

func (s String) Delete(items ...string) String {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

func (s String) Has(item string) bool {
	_, contained := s[item]
	return contained
}

func (s String) HasAll(items ...string) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

func (s String) HasAny(items ...string) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

func (s String) Union(s2 String) String {
	result := NewString()
	for key := range s {
		result.Insert(key)
	}
	for key := range s2 {
		result.Insert(key)
	}
	return result
}

func (s String) IsSuperset(s2 String) bool {
	for item := range s2 {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

type sortableSliceOfString []string

func (s sortableSliceOfString) Len() int           { return len(s) }
func (s sortableSliceOfString) Less(i, j int) bool { return lessString(s[i], s[j]) }
func (s sortableSliceOfString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s String) List() []string {
	res := make(sortableSliceOfString, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return []string(res)
}

func (s String) UnsortedList() []string {
	res := make([]string, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

func (s1 String) Equal(s2 String) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

func (s String) Len() int {
	return len(s)
}

func lessString(lhs, rhs string) bool {
	return lhs < rhs
}
