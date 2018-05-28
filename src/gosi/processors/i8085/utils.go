package i8085

func btoi(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func inBetween(i, min, max uint16) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
