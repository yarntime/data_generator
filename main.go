package main

import (
	"github.com/yarntime/data_generator/mkp"
	"fmt"
	"os"
	"bufio"
)

var normalInstances []mkp.InstanceType
var difficultInstances = []string{"UncorrelatedSpannerInstances", "WeaklyCorrelatedSpannerInstances", "StronglyCorrelatedSpannerInstances",
	"MultipleStronglyCorrelatedInstances", "ProfitCeilingInstances", "CircleInstances"}
var normalBags      []mkp.BagType
var R = []int{1000, 10000, 100000}
var NMPair = [][]int{{50, 5}, {200, 20}}
var BASE_DIR = "./generated_data/"

func init() {
	normalInstances = make([]mkp.InstanceType, 0)
	normalInstances = append(normalInstances, &mkp.UncorrelatedInstances{})
	normalInstances = append(normalInstances, &mkp.WeaklyCorrelatedInstances{})
	normalInstances = append(normalInstances, &mkp.StronglyCorrelatedInstances{})
	normalInstances = append(normalInstances, &mkp.InverseStronglyCorrelatedInstances{})
	normalInstances = append(normalInstances, &mkp.AlmostSctronglyCorrelatedInstances{})
	normalInstances = append(normalInstances, &mkp.SubsetSumInstances{})
	normalInstances = append(normalInstances, &mkp.UncorrelatedInstancesWithSimilarWeights{})

	normalBags = make([]mkp.BagType, 0)
	normalBags = append(normalBags, &mkp.SimilarBag{})
	normalBags = append(normalBags, &mkp.DiSimilarBag{})
}

func main() {
	for i := 0; i < len(R); i++ {
		for j := 0; j < len(NMPair); j++ {
			for k := 0; k < len(normalInstances); k++ {
				weights, profits := normalInstances[k].GenerateInstance(R[i], NMPair[j][0])
				similarCapacities := normalBags[0].GenerateCapacity(weights, profits, NMPair[j][1])
				path, name := getFileName(R[i], NMPair[j][0], NMPair[j][1], normalInstances[k].GetName(), normalBags[0].GetName())
				writeToFile(path, name, weights, profits, similarCapacities)

				dissimilarCapacities := normalBags[1].GenerateCapacity(weights, profits, NMPair[j][1])
				path, name = getFileName(R[i], NMPair[j][0], NMPair[j][1], normalInstances[k].GetName(), normalBags[1].GetName())
				writeToFile(path, name, weights, profits, dissimilarCapacities)
			}

			for k := 0; k < len(difficultInstances); k++ {
				weights, profits := getDifficultInstances(R[i], NMPair[j][0], NMPair[j][1], difficultInstances[k])

				similarCapacities := normalBags[0].GenerateCapacity(weights, profits, NMPair[j][1])
				path, name := getFileName(R[i], NMPair[j][0], NMPair[j][1], difficultInstances[k], normalBags[0].GetName())
				writeToFile(path, name, weights, profits, similarCapacities)

				dissimilarCapacities := normalBags[1].GenerateCapacity(weights, profits, NMPair[j][1])
				path, name = getFileName(R[i], NMPair[j][0], NMPair[j][1], difficultInstances[k], normalBags[1].GetName())
				writeToFile(path, name, weights, profits, dissimilarCapacities)
			}
		}
	}
}

func getDifficultInstances(r int, n int, m int, dateType string) ([]int, []int) {
	var weights []int
	var profits []int
	switch dateType {
	case "UncorrelatedSpannerInstances":
		instanceType := &mkp.UncorrelatedSpannerInstances{}
		weights, profits = instanceType.GenerateInstance(r, n, 2, 10)
	case "WeaklyCorrelatedSpannerInstances":
		instanceType := &mkp.WeaklyCorrelatedSpannerInstances{}
		weights, profits = instanceType.GenerateInstance(r, n, 2, 10)
	case "StronglyCorrelatedSpannerInstances":
		instanceType := &mkp.WeaklyCorrelatedSpannerInstances{}
		weights, profits = instanceType.GenerateInstance(r, n, 2, 10)
	case "MultipleStronglyCorrelatedInstances":
		instanceType := &mkp.MultipleStronglyCorrelatedInstances{}
		weights, profits = instanceType.GenerateInstance(r, n, 3 * r / 10, 2 * r / 10, 6)
	case "ProfitCeilingInstances":
		instanceType := &mkp.ProfitCeilingInstances{}
		weights, profits = instanceType.GenerateInstance(r, n, 3)
	case "CircleInstances":
		instanceType := &mkp.CircleInstances{}
		weights, profits = instanceType.GenerateInstance(r, n, 2.0 / 3.0)
	}
	return weights, profits
}

func getFileName(r int, n int, m int, instanceType string, bagType string) (string, string) {
	return fmt.Sprintf("%s/%s/", BASE_DIR, instanceType), fmt.Sprintf("%s_%d_%d_%d.txt", bagType, r, n, m)
}

func writeToFile(path string, name string, weights []int, profits []int, capacities []int) {
	if b, _ := PathExists(path); !b {
		os.MkdirAll(path, os.ModePerm)
	}
	f, err := os.OpenFile(path + "/" + name, os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i := 0; i < len(weights); i++ {
		w.WriteString(fmt.Sprintf("%d,%d\n", weights[i], profits[i]))
	}

	for i := 0; i < len(capacities); i++ {
		w.WriteString(fmt.Sprintf("%d\n", capacities[i]))
	}
	w.Flush()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}