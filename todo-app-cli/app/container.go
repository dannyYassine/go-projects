package app

import (
	"reflect"

	"go.uber.org/dig"
)

type container struct {
	container   *dig.Container
	innerMap    map[string]func() interface{}
	initialized bool
}

var digContainer *dig.Container = dig.New()

func NewContainer() *container {
	return &container{container: digContainer, innerMap: make(map[string]func() interface{})}
}

var Container = NewContainer()

func (c *container) Bind(constructor interface{}) {
	funcType := reflect.TypeOf(constructor)
	name := funcType.String()

	c.innerMap[name] = func() interface{} {
		return constructor
	}
}

func (c *container) PartialMock(constructor interface{}) {
	funcType := reflect.TypeOf(constructor)
	name := funcType.String()

	c.innerMap[name] = func() interface{} {
		return constructor
	}
}

func (c *container) Build() {
	for _, value := range c.innerMap {
		constructor := value()
		c.container.Provide(constructor)
	}
}

func Get[T any]() *T {
	if !Container.initialized {
		Container.Build()
		Container.initialized = true
	}

	var u *T
	Container.container.Invoke(func(t *T) error {
		u = t
		return nil
	})
	return u
}
