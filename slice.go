package intro

import "bytes"

type Path []byte

func (p Path) Equal(other Path) bool {
	return bytes.Equal(p, other)
}

func (p *Path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p *Path) ToUpper() {
	for i, e := range *p {
		if 'a' <= e && e <= 'z' {
			(*p)[i] = e + 'A' - 'a'
		}
	}
}
