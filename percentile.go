package percentile

import (
	"errors"
	"math"
	"sort"
)

type Interface interface {
	sort.Interface        // sort接口
	Get(int) float64      // 获取指定索引的数据，并将结果转为float64返回
	Avg(int, int) float64 // 获取两个指定索引的数据，计算其算数平均数，并将结果转为float64返回
}

// checkPercentile 检验 百分位数 与 列表 是否合法
func checkPercentile(percentile int, slice Interface) error {
	if percentile > 100 || percentile < 1 {
		return errors.New("error! percentile should between 1 to 99")
	}
	if slice.Len() <= 0 {
		return errors.New("error! slice is empty")
	}
	return nil
}

func Calculate(percentile int, slice Interface) (float64, error) {
	// 验证输入
	err := checkPercentile(percentile, slice)
	if err != nil {
		return 0, err
	}

	// 因为slice实现了percentile.Interface接口的对象,同时也实现了sort.Interface，故可直接进行排序
	sort.Stable(slice)

	// 计算百分位数的索引
	var index int
	i := float64(slice.Len()*percentile) / 100
	if math.Ceil(i) == math.Floor(i) { // 判断是否为整数 （方法较愚蠢，向上取整==向下取整），如果为整数返回 第a个与第a+1个数的算数平均值
		index = int(i - 1)
		return slice.Avg(index, index+1), nil
	} else {
		index = int(math.Ceil(i)) - 1
		return slice.Get(index), nil
	}
}

// int 类型的实现
type IntSlice []int

func (p IntSlice) Len() int { return len(p) }

func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }

func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p IntSlice) Get(i int) float64 { return float64(p[i]) }

func (p IntSlice) Avg(i, j int) float64 { return float64(p[i]+p[j]) / 2 }

func CalculateInt(percentile int, slice IntSlice) (float64, error) {
	res, err := Calculate(percentile, slice)
	if err != nil {
		return 0, err
	}
	return res, err
}

// float 类型的实现
type Float64Slice []float64

func (p Float64Slice) Len() int { return len(p) }

func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] }

func (p Float64Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p Float64Slice) Get(i int) float64 { return p[i] }

func (p Float64Slice) Avg(i, j int) float64 { return (p[i] + p[j]) / 2 }

func CalculateFloat64(percentile int, slice Float64Slice) (float64, error) {
	res, err := Calculate(percentile, slice)
	if err != nil {
		return 0, err
	}
	return res, err
}
