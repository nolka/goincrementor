package interfaces

//IncrementorInterface. Implement this interface to
// get ability for use Incrementor functionality
type IncrementorInterface interface {
	GetNumber() uint64
	IncrementNumber()
	SetMaxValue(uint64) error
}
