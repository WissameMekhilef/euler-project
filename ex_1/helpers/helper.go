package helpers

func GetMultiplesOfUnder(a1 int, a2 int, b int) int {
	result := 0
	for i := 1; i <= b; i++ {
		if i%a1 == 0 || i%a2 == 0 {
			result += 1
		}

	}
	return result
}
