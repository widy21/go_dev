package balance

type Balancer interface {
	Dobalance([]*Instance) (*Instance, error)
}
