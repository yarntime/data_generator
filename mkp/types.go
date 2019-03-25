package mkp

import (
	"math/rand"
	"time"
	"math"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rnd.Intn(max-min) + min
}

type UncorrelatedInstances struct {}

func (ui *UncorrelatedInstances) GetName() string{
	return "UncorrelatedInstances"
}

func (ui *UncorrelatedInstances) GenerateInstance(r int, m int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profits = append(profits, rnd.Intn(r) + 1)
	}
	return weights, profits
}

type WeaklyCorrelatedInstances struct {}

func (wci *WeaklyCorrelatedInstances) GetName() string{
	return "WeaklyCorrelatedInstances"
}

func (wci *WeaklyCorrelatedInstances) GenerateInstance(r int, m int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profits = append(profits, RandInt(weights[i] - (r / 10), weights[i] + (r / 10)))
	}
	return weights, profits
}

type StronglyCorrelatedInstances struct {}

func (sci *StronglyCorrelatedInstances) GetName() string{
	return "StronglyCorrelatedInstances"
}

func (sci *StronglyCorrelatedInstances)  GenerateInstance(r int, m int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profits = append(profits, weights[i] + (r / 10))
	}
	return weights, profits
}

type InverseStronglyCorrelatedInstances struct {}

func (isc *InverseStronglyCorrelatedInstances) GetName() string{
	return "InverseStronglyCorrelatedInstances"
}

func (isc *InverseStronglyCorrelatedInstances) GenerateInstance(r int, m int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		profits = append(profits, rnd.Intn(r) + 1)
		weights = append(weights, profits[i] + r / 10)
	}
	return weights, profits
}

type AlmostSctronglyCorrelatedInstances struct {}

func (asci *AlmostSctronglyCorrelatedInstances) GetName() string{
	return "AlmostSctronglyCorrelatedInstances"
}

func (asci *AlmostSctronglyCorrelatedInstances) GenerateInstance(r int, m int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profits = append(profits, RandInt(weights[i] + r/10 - r/500, weights[i] + r/10 + r/500))
	}
	return weights, profits
}

type SubsetSumInstances struct {}

func (ssi *SubsetSumInstances) GetName() string {
	return "SubsetSumInstances"
}

func (ssi *SubsetSumInstances) GenerateInstance(r int, m int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profits = append(profits, weights[i])
	}
	return weights, profits
}

type UncorrelatedInstancesWithSimilarWeights struct {}

func (uiws *UncorrelatedInstancesWithSimilarWeights) GetName() string{
	return "UncorrelatedInstancesWithSimilarWeights"
}

func (uiws *UncorrelatedInstancesWithSimilarWeights) GenerateInstance(r int, m int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, RandInt(100000, 100100))
		profits = append(profits, RandInt(1, 1000))
	}
	return weights, profits
}

type SpannerInstances struct {}

func (si *SpannerInstances) GetName() string {
	return "SpannerInstances"
}

func (si *SpannerInstances) GenerateInstance(r int, m int, v, mt int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)

	return weights, profits
}

type MultipleStronglyCorrelatedInstances struct{}

func (msci *MultipleStronglyCorrelatedInstances) GetName() string {
	return "MultipleStronglyCorrelatedInstances"
}

// we may choose k1 = 3R / 10, k2 - 2R / 10, d = 6 to generate the most difficult problems
func (msci *MultipleStronglyCorrelatedInstances) GenerateInstance(r int, m int, k1, k2, d int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profit := 0
		if weights[i] % d == 0 {
			profit = weights[i] + k1
		} else {
			profit = weights[i] + k2
		}
		profits = append(profits, profit)
	}
	return weights, profits
}

type ProfitCeilingInstances struct {}

func (pci *ProfitCeilingInstances) GetName() string {
	return "ProfitCeilingInstances"
}

// we may choose d = 3 to generate the most difficult problems
func (pci *ProfitCeilingInstances) GenerateInstance(r int, m int, d int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profits = append(profits, d * int(math.Ceil(float64(weights[i]) / float64(d))))
	}
	return weights, profits
}

type CircleInstances struct {}

func (ci *CircleInstances) GetName() string  {
	return "CircleInstances"
}

// we may choose d = 2/3 to generate the most difficult problems
func (ci *CircleInstances) GenerateInstance(r int, m int, d float64) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)
	for i := 0; i < m; i++ {
		weights = append(weights, rnd.Intn(r) + 1)
		profit := int(math.Ceil(d * math.Sqrt(float64(4 * r * r - (weights[i] - 2 * r) * (weights[i] - 2 * r)))))
		profits = append(profits, profit)
	}
	return weights, profits
}