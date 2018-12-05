package types

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestRectIntersects(t *testing.T) {
	assert := assert.New(t)

	assert.True(Rect{0, 0, 10, 10}.Intersects(Rect{5, 5, 15, 15}))
}
