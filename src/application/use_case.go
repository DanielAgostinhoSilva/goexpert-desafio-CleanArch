package application

type UseCase interface {
	execute(input interface{}) (interface{}, error)
}
