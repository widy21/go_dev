package module

type Balancer interface {
	Dobalance([]*Instance) (*Instance, error)
}
