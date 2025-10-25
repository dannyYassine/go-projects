package app

import "go.uber.org/dig"

type container struct {
	container *dig.Container
}

var digContainer *dig.Container = dig.New()

var Container = &container{container: digContainer}

func (c *container) Bind(constructor interface{}) {
	c.container.Provide(constructor)
}

func Get[T any]() *T {
	var u *T
	Container.container.Invoke(func(t *T) error {
		u = t
		return nil
	})
	return u
}
