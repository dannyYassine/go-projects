package app

import (
	"reflect"
	"sync"

	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
	innerMap  map[string]func() interface{}
	once      sync.Once
}

func NewContainer() *Container {
	return &Container{container: dig.New(), innerMap: make(map[string]func() interface{})}
}

func (c *Container) Bind(constructor interface{}) {
	funcType := reflect.TypeOf(constructor)
	name := funcType.String()

	c.innerMap[name] = func() interface{} {
		return constructor
	}
}

func (c *Container) PartialMock(constructor interface{}) {
	funcType := reflect.TypeOf(constructor)
	name := funcType.String()

	c.innerMap[name] = func() interface{} {
		return constructor
	}
}

func (c *Container) EnsureBuilt() {
	c.once.Do(func() {
		c.Build()
	})
}

func (c *Container) Build() {
	for _, value := range c.innerMap {
		constructor := value()
		err := c.container.Provide(constructor)

		if err != nil {
			panic(err)
		}
	}
}
