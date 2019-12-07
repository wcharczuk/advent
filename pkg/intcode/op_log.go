package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// OpLog is a log item.
type OpLog struct {
	PC         int
	Op         OpCode
	Parameters []OpLogParameter
	Store      OpLogParameter
}

// String implements fmt.Stringer.
func (ol OpLog) String() string {
	var pieces []string
	pieces = append(pieces, fmt.Sprintf("(%d)", ol.PC))
	pieces = append(pieces, ol.Op.String())
	for _, param := range ol.Parameters {
		pieces = append(pieces, param.String())
	}
	if ol.Store.IsReference {
		pieces = append(pieces, "=> "+ol.Store.String())
	}
	return strings.Join(pieces, " ")
}

// OpLogParameter is a parameter to the operation.
type OpLogParameter struct {
	IsReference bool
	Addr        int
	Value       int
}

// String implements fmt.Stringer.
func (olp OpLogParameter) String() string {
	if olp.IsReference {
		return fmt.Sprintf("&%d(%d)", olp.Addr, olp.Value)
	}
	return strconv.Itoa(olp.Value)
}
