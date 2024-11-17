/*
	解析几何相关
*/
package gulc


type Point struct {
	X float64
	Y float64
}


// Line 通过两点定义一条线段
type LineSeg struct {
	P1 Point
	P2 Point
	// ax+by+c=0 方程的参数
	a, b, c float64
}

// CrossPoint 返回两条线段的节点，如果没有交点，则第二个参数返回false，如果重叠则返回最小的交点
func (l *LineSeg) CrossPoint(l2 *LineSeg) (*Point, bool) {
	// 首先考虑两条垂线的情况
	if l.P1.X == l.P2.X {
		if l2.P1.X == l2.P2.X {
			if l.P1.X != l2.P1.X || min(l.P1.X, l.P2.X) > max(l2.P1.X, l2.P2.X) || min(l2.P1.X, l2.P2.X) > max(l.P1.X, l.P2.X) {
				return nil, false
			}
			
		}
	}
	return nil, false
}