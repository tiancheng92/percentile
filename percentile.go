package percentile

import (
	"errors"
	"sort"
)

type Interface interface {
	sort.Interface
	Get(int) interface{}
}

// checkPercentile 检验 百分位数 是否合法
func checkPercentile(percentile int) error {
	if percentile > 100 || percentile < 1 {
		return errors.New("error percentile, percentile should between 1 to 99")
	}
	return nil
}

func Calculate(percentile int, slice Interface) (interface{}, error) {
	// 验证输入
	err := checkPercentile(percentile)
	if err != nil {
		return nil, err
	}

	// 因为实现了 Interface 接口的对象同时也实现了sort接口，故可直接进行排序
	sort.Stable(slice)

	// 计算百分位数的索引
	index := slice.Len() * percentile / 100
	if index < 1 {
		index = 1
	}

	// 取值并返回
	return slice.Get(index - 1), nil
}

// int 类型的实现
type Int []int

func (p Int) Len() int { return len(p) }

func (p Int) Less(i, j int) bool { return p[i] < p[j] }

func (p Int) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p Int) Get(i int) interface{} { return p[i] }

func CalculateInt(percentile int, slice Int) (int, error) {
	res, err := Calculate(percentile, slice)
	if err != nil {
		return 0, err
	}
	return res.(int), err
}

// float 类型的实现
type Float64 []float64

func (p Float64) Len() int { return len(p) }

func (p Float64) Less(i, j int) bool { return p[i] < p[j] }

func (p Float64) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p Float64) Get(i int) interface{} { return p[i] }

func CalculateFloat64(percentile int, slice Float64) (float64, error) {
	res, err := Calculate(percentile, slice)
	if err != nil {
		return 0, err
	}
	return res.(float64), err
}
