package main

func theType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "unknow"
	}
}
