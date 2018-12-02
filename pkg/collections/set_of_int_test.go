package collections

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestSetOfInt(t *testing.T) {
	assert := assert.New(t)

	set := SetOfInt{}
	set.Add(1)
	assert.True(set.Contains(1))
	assert.Equal(1, set.Len())
	assert.False(set.Contains(2))
	set.Remove(1)
	assert.False(set.Contains(1))
	assert.Zero(set.Len())
}

func TestSetOfIntOperations(t *testing.T) {
	assert := assert.New(t)

	a := NewSetOfInt(1, 2, 3, 4)
	b := NewSetOfInt(1, 2)
	c := NewSetOfInt(3, 4, 5, 6)

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
