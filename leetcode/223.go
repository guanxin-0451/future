package leetcode

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area1 := (ax1 - ax2) * (ay1 - ay2)
	area2 := (bx1 - bx2) * (by1 - by2)
	maxX1 := max(ax1, bx1)
	minX2 := min(ax2, bx2)
	maxY1 := max(ay1, by1)
	minY2 := min(ay2, by2)

	x := minX2 - maxX1
	if x < 0 {
		x = 0
	}
	y := minY2 - maxY1
	if y < 0 {
		y = 0
	}

	return area1 + area2 - x*y
}
