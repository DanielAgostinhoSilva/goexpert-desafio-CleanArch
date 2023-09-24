package application

type NullaryUseCase interface {
	execute() (interface{}, error)
}
