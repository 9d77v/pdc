package sensors

import "gobot.io/x/gobot/drivers/i2c"

// Adaptor is the interface that describes an adaptor in gobot
type Adaptor interface {
	// Name returns the label for the Adaptor
	Name() string
	// SetName sets the label for the Adaptor
	SetName(n string)
	// Connect initiates the Adaptor
	Connect() error
	// Finalize terminates the Adaptor
	Finalize() error
	GetConnection(address int, bus int) (device i2c.Connection, err error)
	GetDefaultBus() int
}
