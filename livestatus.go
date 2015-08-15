// Package livestatus provides a binding to MK Livestatus sockets.
package livestatus

// Livestatus is a binding instance.
type Livestatus struct {
	network string
	address string
}

// Query creates a new query instance on a spacific table.
func (l *Livestatus) Query(table string) *Query {
	return newQuery(table, l)
}

// NewLivestatus creates a new binding instance.
func NewLivestatus(network, address string) *Livestatus {
	return &Livestatus{
		network: network,
		address: address,
	}
}
