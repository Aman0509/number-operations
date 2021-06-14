package numbers

func Sum(num ...float64) float64 {
	var temp float64
	for _, i := range num {
		temp += i
	}
	return temp
}

func Product(num ...float64) float64 {
	temp := 1.0
	for _, i := range num {
		temp *= i
	}
	return temp
}
