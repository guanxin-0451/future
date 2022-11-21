package leetcode



func evalRPN(tokens []string) int {
	operators := map[string]func(i, j int)int{
		"+": add,
		"-": sub,
		"*": multiply,
		"/": divide,
	}

	stack := make([]int, 0)

	for _, token:= range(tokens){
		if _, ok := operators[token]; ok{
			op := operators[token]
			second := stack[len(stack)-1]
			first := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			new := op(first, second)
			stack = append(stack, new)
		}else{
			i, _ := strconv.Atoi(token)
			stack = append(stack, i)
		}
	}
	return stack[0]
}


func add(i, j int)int{
	return i + j
}

func sub(i, j int)int{
	return i - j
}

func divide(i, j int)int{
	return i / j
}
func multiply(i, j int)int{
	return i * j
}
