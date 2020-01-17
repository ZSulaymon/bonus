package main

func main() {

}
func bonus(sales []int) int {
	const (
		percentBound = 5
		saleBound    = 10_000
	)
	sumBonus := 0
	for _, sale := range sales{
		if sale > saleBound {
			sumBonus += (sale - saleBound) * percentBound / 100
		}
	}
	return sumBonus
}