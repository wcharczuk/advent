package collections

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestSetOfString(t *testing.T) {
	assert := assert.New(t)

	set := SetOfString{}
	assert.Equal(0, set.Len())

	set.Add("test")
	assert.Equal(1, set.Len())
	assert.True(set.Contains("test"))

	set.Add("test")
	assert.Equal(1, set.Len())
	assert.True(set.Contains("test"))

	set.Add("not test")
	assert.Equal(2, set.Len())
	assert.True(set.Contains("not test"))

	set.Remove("test")
	assert.Equal(1, set.Len())
	assert.False(set.Contains("test"))
	assert.True(set.Contains("not test"))

	set.Remove("not test")
	assert.Equal(0, set.Len())
	assert.False(set.Contains("test"))
	assert.False(set.Contains("not test"))
}

func TestSetOfStringOperations(t *testing.T) {
	assert := assert.New(t)

	a := NewSetOfString("a", "b", "c", "d")
	b := NewSetOfString("a", "b")
	c := NewSetOfString("c", "d", "e", "f")

	union := a.Union(c)
	assert.Len(union, 6)
	intersect := a.Intersect(b)
	assert.Len(intersect, 2)
	diff := a.Difference(c)
	assert.Len(diff, 4)
	diff = c.Difference(a)
	assert.Len(diff, 4)
	assert.True(b.IsSubsetOf(a))
	assert.False(a.IsSubsetOf(b))
}
