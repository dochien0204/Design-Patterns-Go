package decorator

type VeggieMania struct{}

func (v *VeggieMania) GetPrice() int {
	return 10
}
