package types

import (
	"math"
)

//Incrementor test implementation of Incrementorinterface
type Incrementor struct {
	maxValue uint64
	counter  uint64
}

//NewIncrementor Constructor for incrementor object
func NewIncrementor() *Incrementor {
	return &Incrementor{
		maxValue: math.MaxUint64,
		counter:  0,
	}
}

//GetNumber Returns current incrementor internal counter value
func (i *Incrementor) GetNumber() uint64 {
	return i.counter
}

//IncrementNumber Increments internal counter value by 1
func (i *Incrementor) IncrementNumber() {
	if i.counter >= i.maxValue {
		i.counter = 0
		return
	}
	i.counter++
}

//SetMaxValue Set maximum internal counter value. If counter reached that value,
// It will be reseted to 0
func (i *Incrementor) SetMaxValue(value uint64) error {
	i.maxValue = value
	return nil
}
