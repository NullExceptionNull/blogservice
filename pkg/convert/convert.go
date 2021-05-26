package convert

import "strconv"

type StrTo string

func (s StrTo) string() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	return strconv.Atoi(s.string())
}

func (s StrTo) MustInt() int {
	i, _ := s.Int()
	return i
}
