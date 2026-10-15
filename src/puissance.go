package src

func Pow(value int, pow int) int {
	temp := value
	i := 1
	if pow == 0 {
		return 1
	}
	if pow == 1 {
		return value
	}
	for i < pow {
		value = temp * value
		i++
	}
	return value
}
