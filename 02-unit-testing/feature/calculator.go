package feature

type Calculator struct{}

func (Calculator) Add(a int32, b int32) int32 {
	return a + b
}

func (Calculator) Subtract(a int32, b int32) int32 {
	return a - b
}
