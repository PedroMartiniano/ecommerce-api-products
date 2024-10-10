package utils

func UpdateStrValues(current *string, new string) {
	if new != "" {
		*current = new
	}
}

func UpdateFloatValues(current *float64, new float64) {
	if new != 0 {
		*current = new
	}
}
