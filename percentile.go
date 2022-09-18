package percentile

import (
	"errors"
	"fmt"
	"sort"
)

func calc(nums interface{}, n int) (num interface{}, err error) {
	switch nums.(type) {
	case sort.IntSlice:
		nums := nums.(sort.IntSlice)
		if len(nums)*n/100-1 < 0 {
			return nil, errors.New("too little elements")
		}
		i := len(nums)*n/100 - 1
		return nums[i], nil
	case sort.Float64Slice:
		nums := nums.(sort.Float64Slice)
		i := len(nums)*n/100 - 1
		if i < 0 {
			return nil, errors.New("too little elements")
		}
		return nums[i], nil
	}
	return nil, nil
}

func PercentileN(list interface{}, n int) (nums interface{}, err error) {

	if n > 100 {
		return nil, errors.New("Please specify less than 100")
	}

	var numsInt sort.IntSlice
	switch list.(type) {
	case []int:
		for _, v := range list.([]int) {
			numsInt = append(numsInt, int(v))
		}
	case []int32:
		for _, v := range list.([]int32) {
			numsInt = append(numsInt, int(v))
		}
	case []int64:
		for _, v := range list.([]int64) {
			numsInt = append(numsInt, int(v))
		}
	}
	if len(numsInt) > 0 {
		sort.Sort(numsInt)
		return calc(numsInt, n)
	}

	var numsFloat sort.Float64Slice
	switch list.(type) {
	case []float32:
		for _, v := range list.([]float32) {
			numsFloat = append(numsFloat, float64(v))
		}
	case []float64:
		for _, v := range list.([]float64) {
			numsFloat = append(numsFloat, float64(v))
		}
	default:
		fmt.Printf("Not Support type: %T\n", list)
		return nums, errors.New("Not Support type")
	}

	return calc(numsFloat, n)
}
