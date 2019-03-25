package mkp

type InstanceType interface {
	GetName() string
	GenerateInstance(r int, m int) ([]int, []int)
}

type BagType interface {
	GetName() string
	GenerateCapacity(weights []int, profits []int, n int) []int
}