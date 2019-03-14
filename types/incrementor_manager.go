package types

import (
	"log"

	"github.com/nolka/goincrementor/interfaces"
)

//IncrementorManager A manager for control IncrementorInterface types
type IncrementorManager struct {
	incrementor interfaces.IncrementorInterface
}

//NewManager Constructor for IncrementorManager
func NewManager(impl interfaces.IncrementorInterface) *IncrementorManager {
	return &IncrementorManager{
		incrementor: impl,
	}
}

//IncrementNumber wrapper method for number incrementing
func (m *IncrementorManager) IncrementNumber() {
	m.incrementor.IncrementNumber()
}

//GetValue wrapper method for obtaining current counter value from incrementor
func (m *IncrementorManager) GetNumber() uint64 {
	return m.incrementor.GetNumber()
}

//SetMaxValue wrapper method for setting max value in incrementor
func (m *IncrementorManager) SetMaxValue(value uint64) {
	err := m.incrementor.SetMaxValue(value)
	if err != nil {
		log.Printf("Error: %s", err)
	}
}
