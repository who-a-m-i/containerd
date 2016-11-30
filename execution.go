package containerd

import "errors"

var ErrProcessSet = errors.New("container process is already set")

type Runtime interface {
	Create(*Container) (ProcessDelegate, error)
	Start(*Container) error
	Delete(*Container) error
	Exec(*Container, *Process) (ProcessDelegate, error)
	Load(id string) (ProcessDelegate, error)
}
