//
//func main() {
//	////leetcode.Permute([]int{1, 2,3})
//	//node1 := leetcode.TreeNode{Val: 1}
//	//node1.Left = &leetcode.TreeNode{Val: 2}
//	//node1.Right = &leetcode.TreeNode{Val: 3}
//	//node1.Left.Left = &leetcode.TreeNode{Val: 4}
//	//node1.Left.Right = &leetcode.TreeNode{Val: 5}
//	//node1.Left.Left.Left = &leetcode.TreeNode{Val: 7}
//	//node1.Left.Left.Left.Right = &leetcode.TreeNode{Val: 8}
//	//
//	//ans := leetcode.PostorderTraversal(&node1)
//	//fmt.Printf("\n%v\n", ans)
//}

package main

import (
	"fmt"
)

func f1(s []int) {
	s[0] = 100
}

type A struct {
	a int
	b int
	c []int
}

//
//func main() {
//	m := make(map[int]int)
//	for i := 1; i <= 10; i++ {
//		m[i] = i
//	}
//	//var s chan int
//
//	//select {
//	//case <- s:
//	//
//	//}
//	//for k, v := range(m){
//	//	fmt.Printf("%d  %d\n", k, v)
//	//	delete(m, k+1)
//	//}
//	//
//	//fmt.Printf("\n%v\n", m)
//	//
//	ans := leetcode.SearchMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9)
//	fmt.Printf("%v", ans)
//	//fmt.Printf("\n%v %v\n\n", a==b, &a==&b)
//}

func getType(i interface{}) int {
	switch i.(type) {
	case int:
		return 1
	case map[string]interface{}:
		return 2
	default:
		return 0
	}

}

func convert(param map[string]interface{}) {
	for key, val := range param {
		t := getType(val)
		if t == 1 {
			param[key] = fmt.Sprintf("%d-", val)
		} else if t == 2 {
			res, _ := val.(map[string]interface{})
			convert(res)
		}
	}
}

//func main(){
//	data := make(map[string]interface{})
//	data = map[string]interface{}{"a": 1}
//	data["b"] = map[string]interface{}{"aa": 2, "bb": 15}
//
//	convert(data)
//	fmt.Println(data)
//}
//
//func main(){
//
//	//var s1 []int
//	s2 := make([]int,0)
//	//s4 := make([]int,0)
//	//
//	//fmt.Printf("\n%v, %v,%v", s1, s2, s4)
//	//fmt.Printf("\n%v, %v,%v", &s1, &s2, &s4)
//	s3 := s2
//	for i:=1; i<=20; i++{
//		fmt.Printf(" s2 pointer:%+v  s3 pointer:%+v \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
//		s2 = append(s2, i)
//	}
//
//	f1(s2)
//	fmt.Printf("s20 %v\n", s2[0])
//	//fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)),*(*reflect.SliceHeader)(unsafe.Pointer(&s2)),*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
//	//fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
//	//fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
//}

//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//	a := []int{1, 1, 1}
//	b := []int{2, 2, 2}
//	c := []int{3, 3, 3}
//
//	ans := calculate(a, b, c)
//	fmt.Printf("\n%v\n", ans)
//}
//
//func calculate(in ...[]int) []int {
//	result := make([]int, 3)
//	var wg sync.WaitGroup
//
//	wg.Add(len(in))
//	for i, s := range in {
//		s1 := s
//		i1 := i
//		go func() {
//			defer wg.Add(-1)
//			sum := 0
//			for _, value := range s1 {
//				sum += value
//			}
//			result[i1] = sum
//		}()
//	}
//
//	wg.Wait()
//
//	return result
//}

func BinarySearch2(nums []int, appoint int) bool {
	if len(nums) == 0 {
		return false
	}

	l, r := 0, len(nums)-1

	for l <= r {
		half := (l + r) / 2

		if nums[half] == appoint {
			return true
		}

		if nums[half] > appoint {
			r = half - 1
		} else if nums[half] < appoint {
			l = half + 1
		}

		fmt.Printf("\n%d %d\n", l, r)
	}

	return false
}

func f11() (value int) {

	defer func() {
		value++
	}()

	value += 10

	return value
}

func main() {
	v := f11()
	fmt.Printf("\n%v\n", v)
	//b := BinarySearch2([]int{1,2,3,4,5,6}, 2)
	//fmt.Printf("\n%v\n", b)
}
