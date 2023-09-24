package application

type UnitUseCase interface {
	execute(input interface{}) error
}
