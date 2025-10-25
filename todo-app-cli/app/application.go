package app

type application struct {
	Container *container
}

var Application = &application{Container: Container}

func (application *application) Bootstrap() *application {
	application.Register()
	application.Boot()

	return application
}

func (application *application) Register() *application {
	application.Container.Bind(NewCreateTodoUseCase)
	application.Container.Bind(NewUpdateTodoUseCase)
	application.Container.Bind(NewListTodosUseCase)
	application.Container.Bind(NewDeleteTodoUseCase)
	application.Container.Bind(NewConsoleRenderer)
	application.Container.Bind(NewTodoRepositoryInterface)

	return application
}

func (application *application) Boot() *application {
	return application
}
