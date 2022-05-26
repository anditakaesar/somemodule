package calc

func Add(num ...int) int {
	sum := 0
	for _, x := range num {
		sum += x
	}

	return sum
}
