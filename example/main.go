package main

import (
	"fmt"
	"log"
	"percentile"
)

func main() {
	// 从一个 int切片 中获取P90的值
	// percentile.Int类型对percentile.Interface接口的实现在percentile包中
	var sliceInt percentile.Int = []int{1, 3, 4, 6, 2, 9, 4, 7, 10, 11}
	p90, err := percentile.CalculateInt(90, sliceInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p90) // return 10.5

	// 从一个 float64切片 中获取中位数（P50）
	// // percentile.Int类型对percentile.Interface接口的实现在percentile包中
	var sliceFloat percentile.Float64 = []float64{1, 3.2, 2.1, 4.99, 5, 8, 5, 7, 8.97, 4, 234, 0}
	p50, err := percentile.CalculateFloat64(50, sliceFloat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p50) // return 4.995

	// 从一个自定义的结构体列表中取出指定元素的P60值
	// Class 并未实现percentile.Interface接口,需要用户自己实现(实现后即可调用percentile.Calculate方法进行取数)
	var class Class = []Student{
		{Name: "a", Score: 65.1},
		{Name: "b", Score: 35.2},
		{Name: "c", Score: 75.3},
		{Name: "d", Score: 95.6},
		{Name: "e", Score: 82.0},
		{Name: "f", Score: 87.5},
		{Name: "g", Score: 50.4},
		{Name: "h", Score: 30.5},
	}
	p60, err := percentile.Calculate(60, class)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f", p60) // return 75.30
}

type Student struct {
	Name  string
	Score float32
}

// Class 实现了 percentile.Interface 的接口 (同时也实现了sort.Interface接口)
type Class []Student

func (c Class) Len() int { return len(c) }

func (c Class) Less(i, j int) bool { return c[i].Score < c[j].Score }

func (c Class) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func (c Class) Get(i int) float64 { return float64(c[i].Score) }

func (c Class) Avg(i, j int) float64 { return float64(c[i].Score+c[j].Score) / 2 }
