package decorator

type TomatoTopping struct {
	Pizza IPizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.Pizza.GetPrice() + 10
}
