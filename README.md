# 小工具：percentile
### 用途：
* 计算百分位数

### 百分位数概念

* 详细释义编辑

    * 说明一：
用99个数值或99个点，将按大小顺序排列的观测值划分为100个等分，则这99个数值或99个点就称为百分位数，分别以Pl，P2，…，P99代表第1个，第2个，…，第99个百分位数。第j个百分位数j=1,2…100。式中Lj，fj和CFj分别是第j个百分位数所在组的下限值、频数和该组以前的累积频数，Σf是观测值的数目。
百分位通常用第几百分位来表示，如第五百分位，它表示在所有测量数据中，测量值的累计频次达5%。以身高为例，身高分布的第五百分位表示有5%的人的身高小于此测量值，95%的身高大于此测量值。
百分位数则是对应于百分位的实际数值。

    * 说明二：
中位数是第50百分位数。
第25百分位数又称第一个四分位数（First Quartile），用Q1表示；第50百分位数又称第二个四分位数（Second Quartile），用Q2表示；第75百分位数又称第三个四分位数（Third Quartile）,用Q3表示。若求得第p百分位数为小数，可完整为整数。
分位数是用于衡量数据的位置的量度，但它所衡量的，不一定是中心位置。百分位数提供了有关各数据项如何在最小值与最大值之间分布的信息。对于无大量重复的数据，第p百分位数将它分为两个部分。大约有p%的数据项的值比第p百分位数小；而大约有(100-p)%的数据项的值比第p百分位数大。对第p百分位数，严格的定义如下。
第p百分位数是这样一个值，它使得至少有p%的数据项小于或等于这个值，且至少有(100-p)%的数据项大于或等于这个值。
高等院校的入学考试成绩经常以百分位数的形式报告。比如，假设某个考生在入学考试中的语文部分的原始分数为54分。相对于参加同一考试的其他学生来说，他的成绩如何并不容易知道。但是如果原始分数54分恰好对应的是第70百分位数，我们就能知道大约70%的学生的考分比他低，而约30%的学生考分比他高。

* 计算步骤:
    * 第1步：以递增顺序排列原始数据（即从小到大排列）。
    * 第2步：计算指数i=np%
    * 第3步：
        1) 若 i 不是整数，将 i 向上取整。大于i的毗邻整数即为第p百分位数的位置。
        2) 若i是整数，则第p百分位数是第i项与第(i+l)项数据的平均值。

* 摘录自[百度百科](https://baike.baidu.com/item/%E7%99%BE%E5%88%86%E4%BD%8D%E6%95%B0)

### 用法：

```go
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

```