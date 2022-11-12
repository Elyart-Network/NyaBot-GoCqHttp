package Module

func Handler(context interface{}) {
	println(context.(string))
}
