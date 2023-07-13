package catclient

var count = 0

func HTTPGet(name string) (int, string) {
	count++
	switch count {
	case 1:
		return 200, "Hello " + name
	case 2:
		return 200, "Bye " + name
	}
	return 403, "limit exceeded"
}
