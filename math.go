package gulc

// 计算[1, a], [1, b] 两个数组配对下和为x的数目 -> nums[x - 1]
func ComputeSumPairs(a, b int) []int {
	res := make([]int, 0)
	return res
}

// CrossPointTwoRects 计算两个矩形的交点
func CrossPointTwoRects(r1, r2 Rectangle) []Point {
    var crossPoints []Point
    existed := make(map[Point]struct{}, 0)
    var xs = []int{r1.Min.X, r1.Max.X, r2.Min.X, r2.Max.X}
    var ys = []int{r1.Min.Y, r1.Max.Y, r2.Min.Y, r2.Max.Y}
    for _, x := range xs {
        if r1.Min.X <= x && x <= r1.Max.X && r2.Min.X <= x && x <= r2.Max.X {
            for _, y := range ys {
                if r1.Min.Y <= y && y <= r1.Max.Y && r2.Min.Y <= y && y <= r2.Max.Y {
                    p := Point{X: x, Y: y}
                    if _, isPresent := existed[p]; isPresent {
                        if crossPoints == nil {
                            crossPoints = make([]Point, 0, 4)
                        }
                        crossPoints = append(crossPoints, p)
                        existed[p] = struct{}{}
                    }
                }
            }
        }
    }
    return crossPoints
}


func mostFrequentIDs(nums []int, freq []int) []int64 {
	
}