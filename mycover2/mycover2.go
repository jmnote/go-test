package mycover2

func Size(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}
	return "enormous"
}

func Color(a int) string {
	switch a {
	case 1:
		if a := 1; a == 1 {
			return "blue"
		}
	case 2:
		return "green"
	}
	return "black"
}
