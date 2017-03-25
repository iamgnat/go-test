package testify

// Something Something
func Something(x int) {
	if x < 10000000 {
		Something(x + 1)
	}
}
