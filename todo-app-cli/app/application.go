package app

type Application struct {
	Container *Container
}

func NewApplication() *Application {
	return &Application{NewContainer()}
}

func (application *Application) Bootstrap() *Application {
	application.Register()
	application.Boot()

	return application
}

func (application *Application) Register() *Application {
	application.Container.Bind(NewCreateTodoUseCase)
	application.Container.Bind(NewUpdateTodoUseCase)
	application.Container.Bind(NewListTodosUseCase)
	application.Container.Bind(NewDeleteTodoUseCase)
	application.Container.Bind(NewConsoleRenderer)
	application.Container.Bind(NewTodoRepositoryInterface)

	return application
}

func (application *Application) Boot() *Application {
	return application
}

func (application *Application) Shutdown() *Application {
	application.Container = nil

	return application
}

func Get[T any](application *Application) *T {
	if application == nil {
		panic("app.Get: application is nil")
	}
	if application.Container == nil {
		panic("app.Get: application container is nil; was Application.Shutdown() called?")
	}

	application.Container.EnsureBuilt()

	var u *T
	err := application.Container.container.Invoke(func(t *T) error {
		u = t
		return nil
	})

	if err != nil {
		panic(err)
	}

	return u
}
