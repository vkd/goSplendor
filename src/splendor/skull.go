package splendor

import (
	"strconv"
	"strings"
)

type SkullType int

const (
	SBlack SkullType = iota
	SBlue
	SGreen
	SRed
	SWhite
	SGold
)

var (
	AllSkulls = []SkullType{
		SWhite,
		SBlue,
		SGreen,
		SRed,
		SBlack,
		SGold,
	}
	AllVSkulls = AllSkulls[:len(AllSkulls)-1]
)

type Skulls map[SkullType]int

func (s Skulls) String() string {
	ss := make([]string, 6)
	for i, st := range AllSkulls {
		ss[i] = strconv.Itoa(s[st])
	}
	return strings.Join(ss, "|")
}
func (s Skulls) Clear() {
	for k := range s {
		s[k] = 0
	}
}
func (s Skulls) ToArray() []SkullType {
	res := make([]SkullType, 0)
	for st, count := range s {
		for i := 0; i < count; i++ {
			res = append(res, st)
		}
	}
	return res
}

func (s Skulls) Count() int {
	var c int
	for _, v := range s {
		c += v
	}
	return c
}
