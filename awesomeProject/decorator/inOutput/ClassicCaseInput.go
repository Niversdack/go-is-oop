package inOutput

import (
	"strings"
)

type ClassicCase string

func NewClassicCase(string string) Stringer {
	return ClassicCase(string)
}
func (c ClassicCase) String() string {
	return string(c)
}

type lowerCase struct {
	Stringer Stringer
}

func LowerCase(s Stringer) Stringer {
	return lowerCase{Stringer: s}
}

func (l lowerCase) String() string {
	return strings.ToLower(l.Stringer.String())
}

type upCase struct {
	Stringer Stringer
}

func UpCase(s Stringer) Stringer {
	return upCase{Stringer: s}
}

func (l upCase) String() string {
	return strings.ToUpper(l.Stringer.String())
}
