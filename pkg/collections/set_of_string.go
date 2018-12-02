package collections

import (
	"strings"
)

// NewSetOfString creates a new SetOfString.
func NewSetOfString(values ...string) SetOfString {
	set := SetOfString{}
	for _, v := range values {
		set.Add(v)
	}
	return set
}

// SetOfString is a set of strings
type SetOfString map[string]bool

// Add adds an element.
func (ss SetOfString) Add(entry string) {
	if _, hasEntry := ss[entry]; !hasEntry {
		ss[entry] = true
	}
}

// Remove deletes an element, returns if the element was in the set.
func (ss SetOfString) Remove(entry string) bool {
	if _, hasEntry := ss[entry]; hasEntry {
		delete(ss, entry)
		return true
	}
	return false
}

// Contains returns if an element is in the set.
func (ss SetOfString) Contains(entry string) bool {
	_, hasEntry := ss[entry]
	return hasEntry
}

// Len returns the length of the set.
func (ss SetOfString) Len() int {
	return len(ss)
}

// Copy returns a new copy of the set.
func (ss SetOfString) Copy() SetOfString {
	newSet := SetOfString{}
	for key := range ss {
		newSet.Add(key)
	}
	return newSet
}

// Union joins two sets together without dupes.
func (ss SetOfString) Union(other SetOfString) SetOfString {
	union := NewSetOfString()
	for k := range ss {
		union.Add(k)
	}

	for k := range other {
		union.Add(k)
	}
	return union
}

// Intersect returns shared elements between two sets.
func (ss SetOfString) Intersect(other SetOfString) SetOfString {
	intersection := NewSetOfString()
	for k := range ss {
		if other.Contains(k) {
			intersection.Add(k)
		}
	}
	return intersection
}

// Difference returns non-shared elements between two sets.
func (ss SetOfString) Difference(other SetOfString) SetOfString {
	difference := NewSetOfString()
	for k := range ss {
		if !other.Contains(k) {
			difference.Add(k)
		}
	}
	for k := range other {
		if !ss.Contains(k) {
			difference.Add(k)
		}
	}
	return difference
}

// IsSubsetOf returns if a given set is a complete subset of another set,
// i.e. all elements in target set are in other set.
func (ss SetOfString) IsSubsetOf(other SetOfString) bool {
	for k := range ss {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

// Values returns the set as a slice.
func (ss SetOfString) Values() (output []string) {
	output = make([]string, len(ss))
	var index int
	for key := range ss {
		output[index] = key
		index++
	}
	return
}

// String returns the set as a csv string.
func (ss SetOfString) String() string {
	return strings.Join(ss.Values(), ", ")
}
