package base

import (
	"fmt"
	"strings"
)

type Strs []string

func (v *Strs) String() string {
	r := []string{}
	for _, s := range *v {
		r = append(r, fmt.Sprintf("%q", s))
	}
	return strings.Join(r, ", ")
	// return (*v)[0]
}

func (v *Strs) Set(s string) error {
	*v = append(*v, s)
	return nil
}
