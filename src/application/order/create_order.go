package application

type CreateOrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
}

func (props *CreateOrderUseCase) execute(input CreateOrderInput) (CreateOrderOutput, error) {
	//TODO implement me
	panic("implement me")
}
