package leetcode


func getMin(l []int, x int)[]int{
	if x >= len(l){
		return nil
	}

	for i := 0; i<x;i++{
		if l[0] < l[1]{
			l = append(l[:1], l[2:]... )
		}
	}
	return l
}

