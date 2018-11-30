package seq

// Int returns an int sequence.
func Int(end int) *IntSeq {
	return &IntSeq{
		End: &end,
	}
}

// IntSeq is an integer sequence
type IntSeq struct {
	Start, End, Step *int
	Current          *int
}

// WithStart sets the start.
func (i *IntSeq) WithStart(start int) *IntSeq {
	i.Start = &start
	return i
}

// WithEnd sets the end.
func (i *IntSeq) WithEnd(end int) *IntSeq {
	i.End = &end
	return i
}

// WithStep sets the end.
func (i *IntSeq) WithStep(step int) *IntSeq {
	i.Step = &step
	return i
}

// StepOrDefault returns the step or a default.
func (i IntSeq) StepOrDefault() int {
	if i.Step != nil {
		return *i.Step
	}
	return 1
}

// StartOrDefault returns the start or a default.
func (i IntSeq) StartOrDefault() int {
	if i.Start != nil {
		return *i.Start
	}
	return 0
}

// Ensure ensures the sequence.
func (i *IntSeq) Ensure() {
	if i.Current == nil {
		current := i.StartOrDefault()
		i.Current = &current
	}
}

// HasNext returns if there is a possible next value.
func (i IntSeq) HasNext() bool {
	i.Ensure()
	if i.End == nil {
		return true
	}
	return *i.Current < *i.End
}

// Next returns the next value and if there is a valid next value.
func (i *IntSeq) Next() (next int, ok bool) {
	if !i.HasNext() {
		return
	}
	*i.Current = *i.Current + i.StepOrDefault()
	next = *i.Current
	ok = true
	return
}

// Values returns the values for the sequence.
func (i *IntSeq) Values() []int {
	if i.End == nil {
		panic("unbounded sequence; cannot realize values")
	}
	var output []int
	value, ok := i.Next()
	for ok {
		output = append(output, value)
		value, ok = i.Next()
	}
	return output
}
