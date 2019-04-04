package mkp

import (
	"math/rand"
	"time"
	"math"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandInt(min, max int) int {
	if min >= max {
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

type UncorrelatedSpannerInstances struct {}

func (usi *UncorrelatedSpannerInstances) GetName() string {
	return "UncorrelatedSpannerInstances"
}

// we may choose v =2, m = 10 to generate the problems
func (usi *UncorrelatedSpannerInstances) GenerateInstance(r int, m int, v, mt int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)

	tmpWeights := make([]float64, 0)
	tmpProfits := make([]float64, 0)
	for i := 0; i < v; i++ {
		tmpWeights = append(tmpWeights, float64(rnd.Intn(r) + 1) / float64(mt + 1))
		tmpProfits = append(tmpProfits, float64(rnd.Intn(r) + 1) / float64(mt + 1))
	}

	choose := 0
	for i := 0; i < m; i++ {
		a := rnd.Intn(mt) + 1
		weights = append(weights, int(math.Ceil(tmpWeights[choose] * float64(a))))
		profits = append(profits, int(math.Ceil(tmpProfits[choose] * float64(a))))
		choose++
		if choose == v {
			choose = 0
		}
	}
	return weights, profits
}

type WeaklyCorrelatedSpannerInstances struct {}

func (wcsi *WeaklyCorrelatedSpannerInstances) GetName() string {
	return "WeaklyCorrelatedSpannerInstances"
}

// we may choose v =2, m = 10 to generate the problems
func (wcsi *WeaklyCorrelatedSpannerInstances) GenerateInstance(r int, m int, v, mt int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)

	tmpWeights := make([]float64, 0)
	tmpProfits := make([]float64, 0)
	for i := 0; i < v; i++ {
		tmpWeights = append(tmpWeights, float64(rnd.Intn(r) + 1) / float64(mt + 1))
		t := float64(RandInt(int(math.Ceil(tmpWeights[i])) - (r / 10), int(math.Ceil(tmpWeights[i])) + (r / 10))) / float64(mt + 1)
		if t < 0 {
			t = -t
		}
		tmpProfits = append(tmpProfits, t)
	}

	choose := 0
	for i := 0; i < m; i++ {
		a := rnd.Intn(mt) + 1
		weights = append(weights, int(math.Ceil(tmpWeights[choose] * float64(a))))
		profits = append(profits, int(math.Ceil(tmpProfits[choose] * float64(a))))
		choose++
		if choose == v {
			choose = 0
		}
	}
	return weights, profits
}

type StronglyCorrelatedSpannerInstances struct {}

func (wcsi *StronglyCorrelatedSpannerInstances) GetName() string {
	return "StronglyCorrelatedSpannerInstances"
}

// we may choose v =2, m = 10 to generate the problems
func (wcsi *StronglyCorrelatedSpannerInstances) GenerateInstance(r int, m int, v, mt int) ([]int, []int) {
	weights := make([]int, 0)
	profits := make([]int, 0)

	tmpWeights := make([]float64, 0)
	tmpProfits := make([]float64, 0)
	for i := 0; i < v; i++ {
		tmpWeights = append(tmpWeights, float64(rnd.Intn(r) + 1) / float64(mt + 1))
		t := float64(weights[i] + (r / 10)) / float64(mt + 1)
		if t < 0 {
			t = -t
		}
		tmpProfits = append(tmpProfits, t)
	}

	choose := 0
	for i := 0; i < m; i++ {
		a := rnd.Intn(mt) + 1
		weights = append(weights, int(math.Ceil(tmpWeights[choose] * float64(a))))
		profits = append(profits, int(math.Ceil(tmpProfits[choose] * float64(a))))
		choose++
		if choose == v {
			choose = 0
		}
	}
	return weights, profits
}

type MultipleStronglyCorrelatedInstances struct{}

func (msci *MultipleStronglyCorrelatedInstances) GetName() string {
	return "MultipleStronglyCorrelatedInstances"
}

// we may choose k1 = 3R / 10, k2 = 2R / 10, d = 6 to generate the most difficult problems
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

type SimilarBag struct {}

func (sb *SimilarBag) GetName() string {
	return "SimilarBag"
}

func (sb *SimilarBag) GenerateCapacity(weights []int, profits []int, n int) []int {
	capacities := make([]int, 0)
	instanceWeights := 0
	for i := 0; i < len(weights); i++ {
		instanceWeights += weights[i]
	}
	for i := 0; i < n; i++ {
		capacities = append(capacities, RandInt(int(math.Ceil(0.4 * float64(instanceWeights / n))),int(math.Ceil(0.6 * float64(instanceWeights / n)))))
	}
	return capacities
}

type DiSimilarBag struct {}

func (dsb *DiSimilarBag) GetName() string {
	return "DiSimilarBag"
}

func (dsb *DiSimilarBag) GenerateCapacity(weights []int, profits []int, n int) []int {
	capacities := make([]int, 0)
	instanceWeights := 0
	for i := 0; i < len(weights); i++ {
		instanceWeights += weights[i]
	}
	for i := 0; i < n; i++ {
		capacities = append(capacities, RandInt(0, instanceWeights / n * 2))
		instanceWeights -= capacities[i]
	}
	return capacities
}