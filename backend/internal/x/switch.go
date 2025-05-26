package x

import (
	"fmt"
	"strings"
)

type RegisteredCases struct {
	cases  []string
	actual string
}

func (r *RegisteredCases) AddCase(c string) bool {
	r.cases = append(r.cases, c)
	return r.actual == c
}

func (r *RegisteredCases) String() string {
	return "[" + strings.Join(r.cases, ", ") + "]"
}

func (r *RegisteredCases) ToUnknownCase() string {
	return fmt.Sprintf("expected one of %s but got %s", r.String(), r.actual)
}

func SwitchExact(actual string) *RegisteredCases {
	return &RegisteredCases{
		actual: actual,
	}
}
