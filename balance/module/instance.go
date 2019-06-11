package module

import "fmt"

type Instance struct {
	Name string
	Port int
}

func (this *Instance) String() string {
	return fmt.Sprintf("%s : %d", this.Name, this.Port)
}

func CreateInstance(name string, port int) *Instance {
	return &Instance{
		Name: name,
		Port: port,
	}
}
