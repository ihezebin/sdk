package nativer

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfInt(condition bool, trueVal, falseVal int) int {
	return If(condition, trueVal, falseVal).(int)
}

func IfFloat64(condition bool, trueVal, falseVal float64) float64 {
	return If(condition, trueVal, falseVal).(float64)
}

func IfString(condition bool, trueVal, falseVal string) string {
	return If(condition, trueVal, falseVal).(string)
}

func IfBool(condition bool, trueVal, falseVal bool) bool {
	return If(condition, trueVal, falseVal).(bool)
}

func IfError(condition bool, trueVal, falseVal error) error {
	return If(condition, trueVal, falseVal).(error)
}
