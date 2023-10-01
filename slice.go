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

// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func Insert(slice []int, index, value int) []int {
	// Grow the slice by one element
	slice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}
