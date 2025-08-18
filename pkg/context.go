package pkg

var global map[string]string = make(map[string]string)

func GetGlobal() map[string]string {
	// Whatch out
	// This internal var can be manipulated outside
	return global
}
